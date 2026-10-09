package order

import (
	"errors"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestFromUpstream(t *testing.T) {
	tests := []struct {
		name     string
		upstream error
		wantCode codes.Code
		wantMsg  string
	}{
		{"caller's unknown ID stays NotFound", status.Error(codes.NotFound, "account not found"), codes.NotFound, "client"},
		{"caller's bad input stays InvalidArgument", status.Error(codes.InvalidArgument, "bad"), codes.InvalidArgument, "client"},
		{"dependency down stays Unavailable", status.Error(codes.Unavailable, "connection refused"), codes.Unavailable, "server"},
		{"timeout counts as unavailable", status.Error(codes.DeadlineExceeded, "deadline"), codes.Unavailable, "server"},
		{"dependency failure is Internal", status.Error(codes.Internal, "db down"), codes.Internal, "server"},
		{"non-gRPC error is Internal", errors.New("boom"), codes.Internal, "server"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := status.Convert(fromUpstream(tt.upstream, "client", "server"))
			if got.Code() != tt.wantCode || got.Message() != tt.wantMsg {
				t.Errorf("got %v %q, want %v %q", got.Code(), got.Message(), tt.wantCode, tt.wantMsg)
			}
		})
	}
}
