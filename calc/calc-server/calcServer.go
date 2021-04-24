package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"calc/calcpb"
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

func main() {
	fmt.Printf("Calculator server is starting ...\n")

	lis, err := net.Listen("tcp", "0.0.0.0:50002")

	if err != nil {
		log.Fatal("Server : Faild to listen - %v", err)
	}

	s := grpc.NewServer()

	calcpb.RegisterCalculatorServiceServer(s, &calcServer{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("Server : failed to serve - %v", err)
	}
}
