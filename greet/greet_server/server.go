package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

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

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimesRequest was invoked: %v", req)
	firstname := req.GetGreeting().GetFirtsName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstname + " number" + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesReponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
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
