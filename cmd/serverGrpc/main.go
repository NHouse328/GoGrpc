package main

import (
	"context"
	"hello/grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + in.GetName()}, nil
}

func main() {
	print("Running gRPC server")

	listener, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("fail to server: %v", err)
	}
}
