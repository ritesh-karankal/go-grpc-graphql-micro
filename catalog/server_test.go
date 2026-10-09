package catalog

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ritesh-karankal/go-grpc-graphql-micro/catalog/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	elastic "gopkg.in/olivere/elastic.v5"
)

type fakeService struct{ err error }

func (f fakeService) PostProduct(context.Context, string, string, float64) (*Product, error) {
	return nil, f.err
}
func (f fakeService) GetProduct(context.Context, string) (*Product, error) { return nil, f.err }
func (f fakeService) GetProducts(context.Context, uint64, uint64) ([]Product, error) {
	return nil, f.err
}
func (f fakeService) GetProductsByIds(context.Context, []string) ([]Product, error) {
	return nil, f.err
}
func (f fakeService) SearchProducts(context.Context, string, uint64, uint64) ([]Product, error) {
	return nil, f.err
}

func TestGetProductCodes(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want codes.Code
	}{
		{"repository not found", ErrNotFound, codes.NotFound},
		{"elasticsearch 404", &elastic.Error{Status: http.StatusNotFound}, codes.NotFound},
		{"elasticsearch down", errors.New("no available connection"), codes.Internal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &grpcServer{service: fakeService{err: tt.err}}
			_, err := s.GetProduct(context.Background(), &pb.GetProductRequest{Id: "x"})
			if got := status.Code(err); got != tt.want {
				t.Errorf("code = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostProductRejectsBadInput(t *testing.T) {
	s := &grpcServer{service: fakeService{}}
	for _, r := range []*pb.PostProductRequest{
		{Name: "", Price: 10},
		{Name: "Lamp", Price: -1},
	} {
		if _, err := s.PostProduct(context.Background(), r); status.Code(err) != codes.InvalidArgument {
			t.Errorf("%+v: code = %v, want InvalidArgument", r, status.Code(err))
		}
	}
}

// Unknown IDs in a multi-get come back with found=false and no _source. This used to
// dereference a nil pointer and crash the service.
func TestListProductsWithIDsSkipsMissing(t *testing.T) {
	es := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"docs":[
			{"_index":"catalog","_type":"product","_id":"p1","found":true,
			 "_source":{"name":"Lamp","description":"LED","price":29.99}},
			{"_index":"catalog","_type":"product","_id":"missing","found":false}
		]}`))
	}))
	defer es.Close()

	client, err := elastic.NewClient(elastic.SetURL(es.URL), elastic.SetSniff(false), elastic.SetHealthcheck(false))
	if err != nil {
		t.Fatal(err)
	}
	repo := &elasticRepository{client: client}

	products, err := repo.ListProductsWithIDs(context.Background(), []string{"p1", "missing"})
	if err != nil {
		t.Fatal(err)
	}
	if len(products) != 1 || products[0].ID != "p1" || products[0].Name != "Lamp" {
		t.Errorf("products = %+v, want only p1", products)
	}
}
