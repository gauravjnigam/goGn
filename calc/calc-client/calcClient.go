package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	calc_pb "calc/calc_pb"
)

func main() {
	fmt.Println("Hello SumClient")

	conn, err := grpc.Dial("localhost:50002", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("client is not able to connect : %v", err)
	}

	defer conn.Close()

	c := calc_pb.NewCalculatorServiceClient(conn)

	fmt.Printf("created client : %f", c)

}
