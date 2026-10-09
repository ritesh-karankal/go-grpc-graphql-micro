package telemetry

import (
	"context"
	"errors"
	"testing"
)

func TestFlushWaitsForShutdown(t *testing.T) {
	for _, shutdownErr := range []error{nil, errors.New("collector unreachable")} {
		called := false
		Flush(func(ctx context.Context) error {
			called = true
			if _, ok := ctx.Deadline(); !ok {
				t.Error("Flush must bound the shutdown with a deadline")
			}
			return shutdownErr // an error is logged, never fatal
		})
		if !called {
			t.Error("Flush did not call shutdown")
		}
	}
}
