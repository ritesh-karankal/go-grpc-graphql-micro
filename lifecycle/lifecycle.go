// Package lifecycle runs the services' gRPC and HTTP servers until shutdown, then
// drains them, so a rolling deploy doesn't cut off requests in flight.
//
// On Kubernetes the sequence is: the pod's preStop hook sleeps (the load balancer and
// Service endpoints stop routing to it meanwhile), then SIGTERM arrives, the server stops
// accepting new requests, in-flight ones finish within DrainTimeout, and main flushes
// telemetry. preStop + DrainTimeout + telemetry flush must fit within the pod's
// terminationGracePeriodSeconds.
package lifecycle

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

// DrainTimeout bounds how long in-flight requests may take to finish after shutdown starts.
const DrainTimeout = 10 * time.Second

// SignalContext is cancelled on SIGTERM (Kubernetes stopping the pod) or Ctrl+C.
func SignalContext() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
}

// ServeGRPC serves until ctx is cancelled, then stops accepting new RPCs and waits up to
// DrainTimeout for running ones before forcing the server closed.
func ServeGRPC(ctx context.Context, s *grpc.Server, lis net.Listener) error {
	errc := make(chan error, 1)
	go func() { errc <- s.Serve(lis) }()

	select {
	case err := <-errc:
		return err
	case <-ctx.Done():
	}

	drained := make(chan struct{})
	go func() {
		s.GracefulStop()
		close(drained)
	}()
	select {
	case <-drained:
	case <-time.After(DrainTimeout):
		s.Stop()
	}
	return nil
}

// ServeHTTP serves until ctx is cancelled, then stops accepting new connections and waits
// up to DrainTimeout for running requests.
func ServeHTTP(ctx context.Context, s *http.Server) error {
	errc := make(chan error, 1)
	go func() { errc <- s.ListenAndServe() }()

	select {
	case err := <-errc:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), DrainTimeout)
	defer cancel()
	return s.Shutdown(shutdownCtx)
}
