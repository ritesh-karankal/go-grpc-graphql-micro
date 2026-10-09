package lifecycle

import (
	"context"
	"io"
	"net"
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/ritesh-karankal/go-grpc-graphql-micro/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// A request that is running when shutdown starts must still complete.

func TestServeHTTPDrainsInFlightRequests(t *testing.T) {
	started := make(chan struct{})
	srv := &http.Server{Addr: "127.0.0.1:0", Handler: http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		close(started)
		time.Sleep(300 * time.Millisecond)
		io.WriteString(w, "done")
	})}

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	srv.Addr = lis.Addr().String()
	lis.Close()

	ctx, cancel := context.WithCancel(context.Background())
	served := make(chan error, 1)
	go func() { served <- ServeHTTP(ctx, srv) }()
	waitForPort(t, srv.Addr)

	got := make(chan string, 1)
	go func() {
		resp, err := http.Get("http://" + srv.Addr)
		if err != nil {
			got <- "error: " + err.Error()
			return
		}
		defer resp.Body.Close()
		b, _ := io.ReadAll(resp.Body)
		got <- string(b)
	}()

	<-started
	cancel() // shutdown begins while the request is still running

	if body := <-got; body != "done" {
		t.Fatalf("in-flight request got %q, want %q", body, "done")
	}
	if err := <-served; err != nil {
		t.Fatalf("ServeHTTP returned %v", err)
	}
}

type slowAccounts struct {
	pb.UnimplementedAccountServiceServer
	started chan struct{}
}

func (s *slowAccounts) GetAccount(_ context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	close(s.started)
	time.Sleep(300 * time.Millisecond)
	return &pb.GetAccountResponse{Account: &pb.Account{Id: r.Id}}, nil
}

func TestServeGRPCDrainsInFlightRPCs(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	svc := &slowAccounts{started: make(chan struct{})}
	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, svc)

	ctx, cancel := context.WithCancel(context.Background())
	served := make(chan error, 1)
	go func() { served <- ServeGRPC(ctx, s, lis) }()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	rpcErr := make(chan error, 1)
	go func() {
		_, err := pb.NewAccountServiceClient(conn).GetAccount(context.Background(), &pb.GetAccountRequest{Id: "a1"})
		rpcErr <- err
	}()

	<-svc.started
	cancel() // shutdown begins while the RPC is still running

	if err := <-rpcErr; err != nil {
		t.Fatalf("in-flight RPC failed: %v", err)
	}
	if err := <-served; err != nil {
		t.Fatalf("ServeGRPC returned %v", err)
	}
}

func waitForPort(t *testing.T, addr string) {
	t.Helper()
	for i := 0; i < 50; i++ {
		if c, err := net.Dial("tcp", addr); err == nil {
			c.Close()
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatalf("server on %s did not start", addr)
}

func TestServeErrorsAreReturned(t *testing.T) {
	// A listener that is already closed makes Serve fail straight away
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	lis.Close()
	if err := ServeGRPC(context.Background(), grpc.NewServer(), lis); err == nil {
		t.Error("ServeGRPC on a closed listener: want error, got nil")
	}

	// An address that is already taken makes ListenAndServe fail
	busy, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer busy.Close()
	if err := ServeHTTP(context.Background(), &http.Server{Addr: busy.Addr().String()}); err == nil {
		t.Error("ServeHTTP on a busy address: want error, got nil")
	}
}

func TestSignalContextCancelledBySIGTERM(t *testing.T) {
	ctx, stop := SignalContext()
	defer stop()

	if err := syscall.Kill(os.Getpid(), syscall.SIGTERM); err != nil {
		t.Fatal(err)
	}
	select {
	case <-ctx.Done():
	case <-time.After(2 * time.Second):
		t.Fatal("context not cancelled by SIGTERM")
	}
}
