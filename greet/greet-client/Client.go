package main

import (
	"context"
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

	callUninry(c)

}

func callUninry(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Gaurav ",
			LastName:  "Nigam",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Greet rpc...")
	}

	log.Printf("Greet response : %v", res.Result)
}
