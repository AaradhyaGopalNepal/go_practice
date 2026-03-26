package main

import (
	"bufio"
	"fmt"
	mainpb "gRPC_strems/proto/gen"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
)

type server struct {
	mainpb.UnimplementedCalculatorServer
}

func (s *server) GenerateFibonacci(req *mainpb.FibonacciRequest, stream mainpb.Calculator_GenerateFibonacciServer) error {
	n := req.N
	a, b := 0, 1
	for i := 0; i < int(n); i++ {
		err := stream.Send(&mainpb.FibonacciResponse{
			Number: int32(a),
		})
		if err != nil {
			return err
		}
		a, b = b, a+b

	}
	return nil
}

func (s *server) SendNumbers(stream mainpb.Calculator_SendNumbersServer) error {
	var sum int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&mainpb.NumberResponse{
				Number: 0,
				Sum:    sum,
			})
		}
		if err != nil {
			return err
		}
		log.Println(req.GetNumber())
		sum += req.GetNumber()
	}

}

func (s *server) Chat(stream mainpb.Calculator_ChatServer) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
			return err
		}
		log.Println("Received Message:", req.GetMessage())
		fmt.Println("Enter response:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
			return err
		}
		input = strings.TrimSpace(input)
		err = stream.Send(&mainpb.ChatMessage{
			Message: input,
		})
		if err != nil {
			return err
		}
	}
	fmt.Println("Returning Control")
	return nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	mainpb.RegisterCalculatorServer(grpcServer, &server{})

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalln(err)
	}
}
