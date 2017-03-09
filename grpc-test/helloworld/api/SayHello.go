package api

import (
	"golang.org/x/net/context"
	pb "github.com/herlegs/GoTest/grpc-test/helloworld/protobuf"
)

func (s *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error){
	return &pb.HelloResponse{
		Message: "hello ya " + req.Firstname + req.Lastname,
	}, nil
}