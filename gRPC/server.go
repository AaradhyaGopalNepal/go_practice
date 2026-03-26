package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "gRPC/proto/gen"
	fb "gRPC/proto/gen/farewell"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculateServer
	pb.UnimplementedGreeterServer
	fb.UnimplementedAufWiedersehenServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{
		Sum: req.A + req.B,
	}, nil
}

func (s *server) Greet(ctx context.Context, req *pb.HelloRequets) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Name: fmt.Sprint("Hello", req.Name),
	}, nil
}

func (s *server) AufWiedersehen(ctx context.Context, req *fb.GoodByeRequest) (*fb.GoodByeResponse, error) {
	return &fb.GoodByeResponse{
		Message: fmt.Sprint("Bye Bye", req.Name),
	}, nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalculateServer(grpcServer, &server{})
	pb.RegisterGreeterServer(grpcServer, &server{})
	fb.RegisterAufWiedersehenServer(grpcServer, &server{})
	log.Println("Server is running on port", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
