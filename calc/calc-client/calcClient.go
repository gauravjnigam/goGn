package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"

	"calc/calcpb"
)

func main() {
	fmt.Println("Hello SumClient")

	conn, err := grpc.Dial("localhost:50002", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("client is not able to connect : %v", err)
	}

	defer conn.Close()

	c := calcpb.NewCalculatorServiceClient(conn)

	//fmt.Printf("created client : %f", c)

	result := callCalcSumService(c, 3, 10)
	fmt.Printf("Sum from server : %d \n", result)

	callPrimeNumDecoposition(c)

}

func callPrimeNumDecoposition(client calcpb.CalculatorServiceClient) {
	fmt.Printf("Calling prime number decomposition \n")
	req := &calcpb.PrimeNumRequest{
		PrimeNumMessage: &calcpb.PrimeNumMessage{
			PrimeNum: 120,
		},
	}

	resStream, err := client.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatal("Error while calling PrimeNumberDecomposition request - %v", err)
	}

	for {
		res, err := resStream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error while streaming response - %v", err)
		}
		fmt.Printf("%d,", res.GetResult())

	}

	fmt.Printf("\nClient has completed the processing!!!\n")

}

func callCalcSumService(client calcpb.CalculatorServiceClient, n1 int64, n2 int64) int64 {
	req := &calcpb.SumRequest{
		SumMessage: &calcpb.SumMessage{
			Num1: n1,
			Num2: n2,
		},
	}

	res, err := client.CalcSum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculator service - %v", err)
	}

	return res.GetResult()

	//return 10

}
