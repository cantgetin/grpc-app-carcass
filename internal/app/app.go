package app

import (
	"context"
	"google.golang.org/grpc"
	"grpc-carcass/internal/app/service"
	"grpc-carcass/pkg/api"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(ctx context.Context) error {
	// grpc server
	s := grpc.NewServer()
	ctx, cancel := context.WithCancel(ctx)

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen tcp, %v", err)
	}

	// service
	svc := service.New()
	api.RegisterSampleServiceServer(s, svc)

	go func() {
		log.Println("grpc server started")
		if err := s.Serve(l); err != nil {
			log.Fatalf("failed to service grpc server, %v", err)
		}
	}()

	gracefulShutDown(ctx, s, cancel)
	return nil
}

func gracefulShutDown(ctx context.Context, s *grpc.Server, cancel context.CancelFunc) {
	const waitTime = 5 * time.Second

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	select {
	case sig := <-ch:
		log.Printf("os signal received %s", sig.String())
	case <-ctx.Done():
		log.Printf("ctx done %s", ctx.Err().Error())
	}

	s.GracefulStop()
	cancel()
	time.Sleep(waitTime)
}
