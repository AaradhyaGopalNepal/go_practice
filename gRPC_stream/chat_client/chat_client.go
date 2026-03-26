package main

import (
	"context"
	mainpb "gRPC_strems/proto/gen"
	"io"
	"log"
	"time"

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
	stream, err := client.Chat(ctx)

	waitChat := make(chan struct{})

	go func() {
		messages := []string{"Hello", "How are you?", "Goodbye"}
		for _, message := range messages {
			err :=
				stream.Send(&mainpb.ChatMessage{
					Message: message,
				})
			if err != nil {
				log.Fatalln(err)
				return
			}
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
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
			log.Println("Chat Received:", resp.GetMessage())

		}
		close(waitChat)
	}()
	<-waitChat
}
