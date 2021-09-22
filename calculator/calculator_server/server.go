package main

import (
	"context"
	"fmt"
	"log"
	"net"

	calculatorpb "github.com/jhonzp/grpc-go-course/calculator/calculator_pb"
	"google.golang.org/grpc"
)

type server struct{}

func (c *server) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	result := req.Adding.Addone + req.Adding.Addtwo
	fmt.Println("Sum of ", req.Adding.Addone, req.Adding.Addtwo, "equals", result)
	res := &calculatorpb.AddResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello go Sum!!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
