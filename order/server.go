package order

import (
	"context"
	"log/slog"
	"fmt"
	"net"

	account "github.com/ritesh-karankal/go-grpc-graphql-micro/account"
	catalog "github.com/ritesh-karankal/go-grpc-graphql-micro/catalog"
	"github.com/ritesh-karankal/go-grpc-graphql-micro/order/pb"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	pb.UnimplementedOrderServiceServer
	service       Service
	accountClient *account.Client
	catalogClient *catalog.Client
}

func ListenGRPC(s Service, accountURL, catalogURL string, port int) error {
	accountClient, err := account.NewClient(accountURL)
	if err != nil {
		return err
	}

	catalogClient, err := catalog.NewClient(catalogURL)
	if err != nil {
		accountClient.Close()
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return err
	}

	serv := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	pb.RegisterOrderServiceServer(serv, &grpcServer{
		UnimplementedOrderServiceServer: pb.UnimplementedOrderServiceServer{},
		service:                         s,
		accountClient:                   accountClient,
		catalogClient:                   catalogClient,
	})

	reflection.Register(serv)

	return serv.Serve(lis)
}

// gRPC status codes tell callers whose fault an error is: NotFound and InvalidArgument
// are the client's (bad input), Unavailable and Internal are ours. The gateway's SLO
// metrics rely on this.

// fromUpstream turns an error from the account or catalog service into one for our
// caller. The caller's mistakes keep their code, an unreachable dependency stays
// Unavailable, and anything else becomes Internal.
func fromUpstream(err error, clientMsg, serverMsg string) error {
	switch code := status.Code(err); code {
	case codes.NotFound, codes.InvalidArgument:
		return status.Error(code, clientMsg)
	case codes.Unavailable, codes.DeadlineExceeded:
		return status.Error(codes.Unavailable, serverMsg)
	default:
		return status.Error(codes.Internal, serverMsg)
	}
}

// logUpstream logs the caller's mistakes as warnings and real failures as errors.
func logUpstream(ctx context.Context, msg string, err error) {
	switch status.Code(err) {
	case codes.NotFound, codes.InvalidArgument:
		slog.WarnContext(ctx, msg, "err", err)
	default:
		slog.ErrorContext(ctx, msg, "err", err)
	}
}

func (s *grpcServer) PostOrder(ctx context.Context, r *pb.PostOrderRequest) (*pb.PostOrderResponse, error) {
	if len(r.Products) == 0 {
		return nil, status.Error(codes.InvalidArgument, "order has no products")
	}

	_, err := s.accountClient.GetAccount(ctx, r.AccountId)
	if err != nil {
		logUpstream(ctx, "Failed to get account", err)
		return nil, fromUpstream(err, "account not found", "could not check account")
	}


	productIDs := []string{}
	for _, p := range r.Products {
		productIDs = append(productIDs, p.ProductId)
	}

	orderedProducts, err := s.catalogClient.GetProducts(ctx, 0, 0, productIDs, "")
	if err != nil {
		logUpstream(ctx, "Failed to get products", err)
		return nil, fromUpstream(err, "products not found", "could not load products")
	}

	products := []OrderedProduct{}
	for _, p := range orderedProducts {
		product := OrderedProduct{
			ID:          p.ID,
			Quantity:    0,
			Price:       p.Price,
			Name:        p.Name,
			Description: p.Description,
		}

		for _, rp := range r.Products {
			if rp.ProductId == p.ID {
				product.Quantity = rp.Quantity
				break
			}
		}

		if product.Quantity != 0 {
			products = append(products, product)
		}
	}

	// Unknown product IDs are dropped above; an order with none left is the caller's mistake
	if len(products) == 0 {
		slog.WarnContext(ctx, "Order has no existing products", "requested", len(r.Products))
		return nil, status.Error(codes.InvalidArgument, "none of the ordered products exist")
	}

	order, err := s.service.PostOrder(ctx, r.AccountId, products)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to post order", "err", err)
		return nil, status.Error(codes.Internal, "could not post order")
	}

	orderProto := &pb.Order{
		Id:         order.ID,
		AccountId:  order.AccountID,
		TotalPrice: order.TotalPrice,
		Products:   []*pb.Order_OrderProduct{},
	}

	orderProto.CreatedAt, _ = order.CreatedAt.MarshalBinary()
	for _, p := range order.Products {
		orderProto.Products = append(orderProto.Products, &pb.Order_OrderProduct{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Quantity:    p.Quantity,

		})
	}
	return &pb.PostOrderResponse{Order: orderProto}, nil
}



func (s *grpcServer) GetOrdersForAccount(ctx context.Context, r *pb.GetOrdersForAccountRequest) (*pb.GetOrdersForAccountResponse, error) {
	accountOrders, err := s.service.GetOrdersForAccount(ctx, r.AccountId)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get orders for account", "err", err)
		return nil, status.Error(codes.Internal, "could not get orders")
	}

	productIDMap := map[string]bool{}
	for _, o := range accountOrders {
		for _, p := range o.Products {
			productIDMap[p.ID] = true
		}
	}

	productIDs := []string{}
	for id := range productIDMap {
		productIDs = append(productIDs, id)
	}

	products, err := s.catalogClient.GetProducts(ctx, 0, 0, productIDs, "")
	if err != nil {
		logUpstream(ctx, "Failed to get products for account orders", err)
		return nil, fromUpstream(err, "products not found", "could not load products for orders")
	}

	orders := []*pb.Order{}
	for _, o := range accountOrders {
		op := &pb.Order{
			AccountId:  o.AccountID,
			Id:         o.ID,
			TotalPrice: o.TotalPrice,
			Products:   []*pb.Order_OrderProduct{},
		}

		op.CreatedAt, _ = o.CreatedAt.MarshalBinary()

		for _, product := range o.Products {
			for _, p := range products {
				if product.ID == p.ID {
					product.Name = p.Name
					product.Description = p.Description
					product.Price = p.Price
					break
				}
			}

			op.Products = append(op.Products, &pb.Order_OrderProduct{
				Id:          product.ID,
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Quantity:    product.Quantity,
			})
		}

		orders = append(orders, op)
	}
	return &pb.GetOrdersForAccountResponse{Orders: orders}, nil
}