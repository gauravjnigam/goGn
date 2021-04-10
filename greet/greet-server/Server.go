package main

import (
	"context"
	"fmt"
	"log"
	"net"

	greetpb "greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet func is invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	result := "Hello " + firstName + lastName

	res := &greetpb.GreetResponse{
		Result: result,
	}
	log.Printf("Greet func is responding...")
	return res, nil
}

func main() {
	fmt.Println("Hello! Server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50001")

	if err != nil {
		log.Fatal("Server : Faild to listen - %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("Server : failed to serve - %v", err)
	}

}
