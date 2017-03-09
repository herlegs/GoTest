package server

import (
	"net"
	"flag"
	"fmt"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc"
	pb "github.com/herlegs/GoTest/grpc-test/helloworld/protobuf"
	"github.com/herlegs/GoTest/grpc-test/helloworld/api"
	"golang.org/x/net/context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"net/http"
)

var (
	GRPCPort = flag.Int("grpc", 10000, "grpc server port")
	HTTPPort = flag.Int("http", 8000, "http server port")
)

func StartRPC(){
	flag.Parse()
	fmt.Println("running grpc at port", *GRPCPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *GRPCPort))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHWServiceServer(s, &api.Server{})
	if err := s.Serve(lis); err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
	}
}

func StartHTTP(){
	flag.Parse()
	fmt.Println("running http at port", *HTTPPort)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterHWServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", *GRPCPort), opts)
	if err != nil {
		grpclog.Fatalf("failed to register service: %v", err)
	}
	http.ListenAndServe(fmt.Sprintf(":%d", *HTTPPort), mux)
}




