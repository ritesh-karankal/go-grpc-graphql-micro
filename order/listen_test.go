package order

import (
	"context"
	"testing"
)

func TestListenGRPCStopsWhenContextIsCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	// gRPC clients connect lazily, so unreachable dependencies don't block startup
	err := ListenGRPC(ctx, fakeOrders{}, "127.0.0.1:1", "127.0.0.1:1", 0)
	if err != nil {
		t.Fatalf("ListenGRPC returned %v, want nil after shutdown", err)
	}
}
