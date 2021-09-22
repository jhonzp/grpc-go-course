package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"
	"time"

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

func (c *server) PrimeNumberDescomposition(req *calculatorpb.PrimeNumberDescompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDescompositionServer) error {
	fmt.Printf("PrimeNumberDescomposition was invoked: %v", req)
	var n float64 = float64(req.GetPrimeNumber().PrimeNumber)
	var k float64 = float64(2)
	for n > 1 {
		if math.Mod(n, k) == 0 {
			res := &calculatorpb.PrimeNumberDescompositionResponse{
				Result: int32(k),
			}
			n = (n / k)
			stream.Send(res)
			time.Sleep(300 * time.Millisecond)
		} else {
			k = k + 1
		}
	}
	return nil
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
