package main

import (
	"context"
	"fmt"
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

	c := calculatorpb.NewAddServiceClient(cc)
	doUnary(c)

}

func doUnary(c calculatorpb.AddServiceClient) {
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
