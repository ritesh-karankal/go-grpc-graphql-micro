package account 

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"strings"

	"github.com/ritesh-karankal/go-grpc-graphql-micro/account/pb"
	"github.com/ritesh-karankal/go-grpc-graphql-micro/lifecycle"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service Service
}

// ListenGRPC serves until ctx is cancelled, then drains in-flight RPCs.
func ListenGRPC(ctx context.Context, s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	serv := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	pb.RegisterAccountServiceServer(serv, &grpcServer{service: s})
	reflection.Register(serv)
	return lifecycle.ServeGRPC(ctx, serv, lis)
}

// gRPC status codes tell callers whose fault an error is: NotFound and InvalidArgument
// are the client's (bad input), Internal is ours. The gateway's SLO metrics rely on this.

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	if strings.TrimSpace(r.Name) == "" {
		return nil, status.Error(codes.InvalidArgument, "account name is required")
	}

	a, err := s.service.PostAccount(ctx, r.Name)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to post account", "err", err)
		return nil, status.Error(codes.Internal, "could not create account")
	}
	return &pb.PostAccountResponse{Account: &pb.Account{
		Id: a.ID,
		Name: a.Name,
	}}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	a, err := s.service.GetAccount(ctx, r.Id)
	if errors.Is(err, sql.ErrNoRows) {
		slog.WarnContext(ctx, "Account not found", "id", r.Id)
		return nil, status.Error(codes.NotFound, "account not found")
	}
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get account", "err", err)
		return nil, status.Error(codes.Internal, "could not get account")
	}

	return &pb.GetAccountResponse{
		Account: &pb.Account{
			Id: a.ID,
			Name: a.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	res, err := s.service.GetAccounts(ctx, r.Skip, r.Take)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get accounts", "err", err)
		return nil, status.Error(codes.Internal, "could not list accounts")
	}
	accounts := []*pb.Account{}
	for _, p := range res {
		accounts = append(
			accounts,
			&pb.Account{
				Id: p.ID,
				Name: p.Name,
			},
		)
	}
	return &pb.GetAccountsResponse{Accounts: accounts}, nil
}
