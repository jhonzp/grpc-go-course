package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/jhonzp/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm GO Client")
	//get client connection
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//doUnary(c)
	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do a Unary RPC .. ")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirtsName: "Jhon Steven",
			LastName:  "Zuñiga",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("could not invoked to Greet grpc function: %v", err)
	}

	log.Printf("Response: %v", res.Result)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC .. ")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirtsName: "Jhon",
			LastName:  "ZP",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("could not invoked to Greet Server Streaming grpc function: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		fmt.Printf("Response from GreerManyTimes: %+v\n", msg.GetResult())
	}

}
