package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"goGn/calc/calcpb"
)

type calcServer struct{}

func (*calcServer) CalcSum(context context.Context, request *calcpb.SumRequest) (response *calcpb.SumResponse, err error) {
	log.Printf("Recieved request - %v", request)
	num1 := request.GetSumMessage().GetNum1()
	num2 := request.GetSumMessage().GetNum2()
	sum := num1 + num2
	res := calcpb.SumResponse{
		Result: sum,
	}

	return &res, nil

}

func (*calcServer) PrimeNumberDecomposition(request *calcpb.PrimeNumRequest, streamData calcpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Recived PrimeNumber Decomposition request %v \n", request)
	primeNum := request.GetPrimeNumMessage().GetPrimeNum()

	var k int64 = 2
	n := primeNum
	for n > 1 {
		if n%k == 0 {
			res := &calcpb.PrimeNumResponse{
				Result: k,
			}
			streamData.Send(res)
			n = n / k
			time.Sleep(1000 * time.Millisecond)
		} else {
			k = k + 1
		}
	}
	fmt.Printf("Prime number decomposition is completed by server!!!\n")
	return nil
}

func main() {
	fmt.Printf("Calculator server is starting ...\n")

	lis, err := net.Listen("tcp", "0.0.0.0:50002")

	if err != nil {
		log.Fatalf("Server : Faild to listen - %v", err)
	}

	s := grpc.NewServer()

	calcpb.RegisterCalculatorServiceServer(s, &calcServer{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Server : failed to serve - %v", err)
	}
}
