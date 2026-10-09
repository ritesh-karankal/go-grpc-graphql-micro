package account

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/ritesh-karankal/go-grpc-graphql-micro/account/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type fakeService struct{ err error }

func (f fakeService) PostAccount(context.Context, string) (*Account, error) { return nil, f.err }
func (f fakeService) GetAccount(context.Context, string) (*Account, error)  { return nil, f.err }
func (f fakeService) GetAccounts(context.Context, uint64, uint64) ([]Account, error) {
	return nil, f.err
}

func TestGetAccountCodes(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want codes.Code
	}{
		{"unknown ID is NotFound", sql.ErrNoRows, codes.NotFound},
		{"database failure is Internal", errors.New("connection reset"), codes.Internal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &grpcServer{service: fakeService{err: tt.err}}
			_, err := s.GetAccount(context.Background(), &pb.GetAccountRequest{Id: "x"})
			if got := status.Code(err); got != tt.want {
				t.Errorf("code = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostAccountCodes(t *testing.T) {
	s := &grpcServer{service: fakeService{err: errors.New("insert failed")}}

	_, err := s.PostAccount(context.Background(), &pb.PostAccountRequest{Name: "  "})
	if got := status.Code(err); got != codes.InvalidArgument {
		t.Errorf("empty name: code = %v, want InvalidArgument", got)
	}

	_, err = s.PostAccount(context.Background(), &pb.PostAccountRequest{Name: "liz"})
	if got := status.Code(err); got != codes.Internal {
		t.Errorf("insert failure: code = %v, want Internal", got)
	}
}
