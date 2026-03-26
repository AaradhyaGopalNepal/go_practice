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
			break
		}
		if err != nil {
			log.Fatalln("Error receiving GenerateFibonacci stream:", err)
			return
		}
		log.Println("Fibonacci Number:", resp.GetNumber())

	}

	stream1, err := client.SendNumbers(ctx)
	if err != nil {
		log.Fatalln("Error sending to SendNumbers func:", err)
		return
	}
	for num := range 9 {
		err := stream1.Send(&mainpb.NumberRequest{Number: int32(num)})
		if err != nil {
			log.Fatalln("Error sending number:", err)
		}

	}
	resp, err := stream1.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error receiving final response:", err)
		return
	}
	log.Println("Sum:", resp.Sum)

}
