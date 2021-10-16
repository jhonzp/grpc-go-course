package main

import (
	"context"
	"fmt"
	"io"
	"log"

	calculatorpb "github.com/jhonzp/grpc-go-course/calculator/calculator_pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'm the client calculator")
	//get client connection
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	//doUnary(c)
	//doStreamServer(c)
	doClientStreaming(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.AddRequest{
		Adding: &calculatorpb.Adding{
			Addone: 10,
			Addtwo: 4,
		},
	}
	res, err := c.Add(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not invoke to calculator grpc services: %v", err)
	}
	fmt.Printf("\nResponse %+v\n", res.Result)
}

func doStreamServer(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.PrimeNumberDescompositionRequest{
		PrimeNumber: &calculatorpb.PrimeNumber{
			PrimeNumber: 120,
		},
	}

	resStream, err := c.PrimeNumberDescomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("could not invoked to PrimeNumberDescomposition Server Streaming grpc function: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		fmt.Printf("Response from PrimeNumberDescomposition: %+v\n", msg.GetResult())
	}
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	request := []*calculatorpb.AVGNumberRequest{
		{
			Number: 1,
		},
		{
			Number: 2,
		},
		{
			Number: 3,
		},
		{
			Number: 4,
		},
	}

	stream, err := c.ComputeAvarage(context.Background())

	if err != nil {
		log.Fatalf("Error while calling ComputeAvarage %v", err)
	}

	for _, req := range request {
		fmt.Printf("Sending: %+v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while reciving response : %v", err)
	}

	fmt.Printf("AVGNumberRequest response: %v\n", res)

}
