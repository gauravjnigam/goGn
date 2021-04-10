package main

import (
	"fmt"
	"log"
	"net"

	greetpb "greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

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
