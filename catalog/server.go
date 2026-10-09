package catalog

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"strings"

	"github.com/ritesh-karankal/go-grpc-graphql-micro/catalog/pb"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	elastic "gopkg.in/olivere/elastic.v5"
)

type grpcServer struct {
	pb.UnimplementedCatalogServiceServer
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	serv := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	pb.RegisterCatalogServiceServer(serv, &grpcServer{
		UnimplementedCatalogServiceServer: pb.UnimplementedCatalogServiceServer{},
		service: s,
	})
	reflection.Register(serv)
	return serv.Serve(lis)
}

// gRPC status codes tell callers whose fault an error is: NotFound and InvalidArgument
// are the client's (bad input), Internal is ours. The gateway's SLO metrics rely on this.

func (s *grpcServer) PostProduct(ctx context.Context, r *pb.PostProductRequest) (*pb.PostProductResponse, error) {
	if strings.TrimSpace(r.Name) == "" || r.Price < 0 {
		return nil, status.Error(codes.InvalidArgument, "product needs a name and a non-negative price")
	}

	p, err := s.service.PostProduct(ctx, r.Name, r.Description, r.Price)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to post product", "err", err)
		return nil, status.Error(codes.Internal, "could not create product")
	}

	return &pb.PostProductResponse{Product: &pb.Product{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}}, nil

}

func (s *grpcServer) GetProduct(ctx context.Context, r *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	p, err := s.service.GetProduct(ctx, r.Id)
	// elastic.v5 reports a missing document as a 404 error, not as Found=false
	if errors.Is(err, ErrNotFound) || elastic.IsNotFound(err) {
		slog.WarnContext(ctx, "Product not found", "id", r.Id)
		return nil, status.Error(codes.NotFound, "product not found")
	}
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get product", "err", err)
		return nil, status.Error(codes.Internal, "could not get product")
	}

	return &pb.GetProductResponse{
		Product: &pb.Product{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		},
	}, nil
}

func (s *grpcServer) GetProducts(ctx context.Context, r *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	var res []Product
	var err error
	if r.Query != "" {
		res, err = s.service.SearchProducts(ctx, r.Query, r.Skip, r.Take)

	} else if len(r.Ids) != 0 {
		res, err = s.service.GetProductsByIds(ctx, r.Ids)
	} else {
		res, err = s.service.GetProducts(ctx, r.Skip, r.Take)
	}

	if err != nil {
		slog.ErrorContext(ctx, "Failed to get products", "err", err)
		return nil, status.Error(codes.Internal, "could not get products")
	}

	products := []*pb.Product{}
	for _, p := range res {
		products = append(
			products,
			&pb.Product{
				Id:          p.ID,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
			},
		)
	}
	return &pb.GetProductsResponse{Products: products}, nil
}

