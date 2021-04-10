package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	greetpb "greet/greetpb"
)

func main() {
	fmt.Println("Hello client")

	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("client is not able to connect : %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	fmt.Printf("created client : %f", c)

}
