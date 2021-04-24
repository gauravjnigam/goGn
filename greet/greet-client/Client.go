package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	greetpb "goGn/greet/greetpb"
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

	//callUninry(c)

	//callServerStreaming(c)

	callClientStreaming(c)

	//callBiDirectionStreaming(c)

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

func callServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Printf("Starting the serverStreaming \n")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Gaurav",
			LastName:  "Nigam",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetingmanyTimes - %v\n", err)
	}

	for {
		msg, err := resStream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming response - %v", err)
		}
		fmt.Printf("Response from GreetingManyTimes : %v\n", msg.GetResult())
	}
}

func callClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Printf("Starting the clientStreaming \n")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Gaurav",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ravi",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Anu",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Dipe",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while sending LongGreet request : %v\n", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req : %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response %v\n", err)
	}

	fmt.Printf("%s", res.GetResult())
}

func callBiDirectionStreaming(c greetpb.GreetServiceClient) {
	fmt.Printf("Starting the BiDirectionalStreaming \n")
}
