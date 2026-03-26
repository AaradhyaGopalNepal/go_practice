package main

import (
	"context"
	mainpb "gRPC_client/proto/gen"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := mainpb.NewCalculatorClient(conn)
	ctx := context.Background()
	req := &mainpb.FibonacciRequest{
		N: 10,
	}
	stream, err := client.GenerateFibonacci(ctx, req)
	if err != nil {
		log.Fatalln("Error calling GenerateFibonacci func:", err)
		return
	}

	for {

		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of stream")
			return
		}
		if err != nil {
			log.Fatalln("Error receiving GenerateFibonacci stream:", err)
			return
		}
		log.Println("Fibonacci Number:", resp.GetNumber())

	}

}
