package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	greetpb "goGn/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

// unary request
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

// server streaming
func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetingManyTimes function is invoked by %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(3000 * time.Millisecond)
	}
	return nil
}

// client streaming
func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("Server received the LongGreet request ..")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// done reading client streamed reqeuest
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatal("Error while streaming LongGreet request \n", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "!! "
	}

	//return nil
}

// Bi-Direction request response
func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Println("Server received the GreetEveryone request ..")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// done reading client streamed reqeuest
			return nil
		}
		if err != nil {
			log.Fatalf("Error while streaming GreetEveryone request - %v\n", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "!! "
		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})

		if sendErr != nil {
			log.Fatalf("Error while sending stream response %v", sendErr)
			return sendErr
		}

	}

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
