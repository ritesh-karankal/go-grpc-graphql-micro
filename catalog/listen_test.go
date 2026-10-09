package catalog

import (
	"context"
	"testing"
)

func TestListenGRPCStopsWhenContextIsCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := ListenGRPC(ctx, fakeService{}, 0); err != nil {
		t.Fatalf("ListenGRPC returned %v, want nil after shutdown", err)
	}
}
