package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jhonzp/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetReponse, error) {
	fmt.Printf("grpc Greet function was invoked with %v", req)
	firstName := req.GetGreeting().GetFirtsName()
	lastName := req.GetGreeting().GetLastName()
	result := fmt.Sprintln("Hello!!", firstName, lastName)
	res := &greetpb.GreetReponse{
		Result: result,
	}
	return res, nil

}

func main() {
	fmt.Println("Hello Go World!!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
