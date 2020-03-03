package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/endpoints"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/pb"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/service"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/service/implementation"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/transport"
	"google.golang.org/grpc"
)

func main() {
	var (
		grpcAddr = flag.String("grpcAddr", ":8081", "gRPC listen port")
	)
	flag.Parse()

	ctx := context.Background()
	errChan := make(chan error)

	var svc service.Service
	svc = implementation.LoremService{}

	endpoints := endpoints.Endpoints{
		WordEndpoint:      endpoints.MakeWordEndpoint(svc),
		SentenceEndpoint:  endpoints.MakeSentenceEndpoint(svc),
		ParagraphEndpoint: endpoints.MakeParagraphEndpoint(svc),
	}

	go func() {
		listener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			errChan <- err
			return
		}

		handler := transport.NewGRPCService(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterLoremServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}
