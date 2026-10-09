package order

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/ritesh-karankal/go-grpc-graphql-micro/account"
	accountpb "github.com/ritesh-karankal/go-grpc-graphql-micro/account/pb"
	"github.com/ritesh-karankal/go-grpc-graphql-micro/catalog"
	catalogpb "github.com/ritesh-karankal/go-grpc-graphql-micro/catalog/pb"
	"github.com/ritesh-karankal/go-grpc-graphql-micro/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Fake account and catalog services, served over real gRPC on localhost so the
// order service's clients and the status codes on the wire are the real ones.

type fakeAccounts struct {
	accountpb.UnimplementedAccountServiceServer
	err error
}

func (f *fakeAccounts) GetAccount(_ context.Context, r *accountpb.GetAccountRequest) (*accountpb.GetAccountResponse, error) {
	if f.err != nil {
		return nil, f.err
	}
	return &accountpb.GetAccountResponse{Account: &accountpb.Account{Id: r.Id, Name: "liz"}}, nil
}

type fakeCatalog struct {
	catalogpb.UnimplementedCatalogServiceServer
	products []*catalogpb.Product
	err      error
}

func (f *fakeCatalog) GetProducts(context.Context, *catalogpb.GetProductsRequest) (*catalogpb.GetProductsResponse, error) {
	if f.err != nil {
		return nil, f.err
	}
	return &catalogpb.GetProductsResponse{Products: f.products}, nil
}

type fakeOrders struct {
	err   error
	saved *[]OrderedProduct
}

func (f fakeOrders) PostOrder(_ context.Context, accountID string, products []OrderedProduct) (*Order, error) {
	if f.err != nil {
		return nil, f.err
	}
	*f.saved = products
	return &Order{ID: "o1", AccountID: accountID, Products: products, TotalPrice: 59.98}, nil
}

func (f fakeOrders) GetOrdersForAccount(context.Context, string) ([]Order, error) { return nil, f.err }

func serve(t *testing.T, register func(*grpc.Server)) string {
	t.Helper()
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	s := grpc.NewServer()
	register(s)
	go s.Serve(lis)
	t.Cleanup(s.Stop)
	return lis.Addr().String()
}

func newTestServer(t *testing.T, accounts *fakeAccounts, products *fakeCatalog, orders fakeOrders) *grpcServer {
	t.Helper()
	accountAddr := serve(t, func(s *grpc.Server) { accountpb.RegisterAccountServiceServer(s, accounts) })
	catalogAddr := serve(t, func(s *grpc.Server) { catalogpb.RegisterCatalogServiceServer(s, products) })

	accountClient, err := account.NewClient(accountAddr)
	if err != nil {
		t.Fatal(err)
	}
	catalogClient, err := catalog.NewClient(catalogAddr)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { accountClient.Close(); catalogClient.Close() })

	return &grpcServer{service: orders, accountClient: accountClient, catalogClient: catalogClient}
}

func TestPostOrder(t *testing.T) {
	lamp := &catalogpb.Product{Id: "p1", Name: "Lamp", Price: 29.99}
	request := &pb.PostOrderRequest{
		AccountId: "a1",
		Products:  []*pb.PostOrderRequest_OrderProduct{{ProductId: "p1", Quantity: 2}},
	}

	tests := []struct {
		name     string
		req      *pb.PostOrderRequest
		accounts *fakeAccounts
		catalog  *fakeCatalog
		orderErr error
		want     codes.Code
	}{
		{"order placed", request,
			&fakeAccounts{}, &fakeCatalog{products: []*catalogpb.Product{lamp}}, nil, codes.OK},
		{"empty order", &pb.PostOrderRequest{AccountId: "a1"},
			&fakeAccounts{}, &fakeCatalog{}, nil, codes.InvalidArgument},
		{"unknown account", request,
			&fakeAccounts{err: status.Error(codes.NotFound, "account not found")}, &fakeCatalog{}, nil, codes.NotFound},
		{"account service failing", request,
			&fakeAccounts{err: status.Error(codes.Internal, "db down")}, &fakeCatalog{}, nil, codes.Internal},
		{"catalog unreachable", request,
			&fakeAccounts{}, &fakeCatalog{err: status.Error(codes.Unavailable, "es down")}, nil, codes.Unavailable},
		{"none of the products exist", request,
			&fakeAccounts{}, &fakeCatalog{}, nil, codes.InvalidArgument},
		{"saving the order fails", request,
			&fakeAccounts{}, &fakeCatalog{products: []*catalogpb.Product{lamp}}, errors.New("insert failed"), codes.Internal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var saved []OrderedProduct
			s := newTestServer(t, tt.accounts, tt.catalog, fakeOrders{err: tt.orderErr, saved: &saved})

			res, err := s.PostOrder(context.Background(), tt.req)
			if got := status.Code(err); got != tt.want {
				t.Fatalf("code = %v (%v), want %v", got, err, tt.want)
			}
			if tt.want != codes.OK {
				return
			}
			if len(saved) != 1 || saved[0].ID != "p1" || saved[0].Quantity != 2 {
				t.Errorf("saved products = %+v, want p1 x2", saved)
			}
			if res.Order.Id != "o1" || len(res.Order.Products) != 1 {
				t.Errorf("response = %+v", res.Order)
			}
		})
	}
}
