package main

import (
	"context"
	"fmt"
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
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {

	//fmt.Printf("Created Go Client: %f", c)
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
