package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	calc_pb "calc/calc_pb"
)

type calcServer struct{}

func (*calcServer) CalcSum(context context.Context, request *calc_pb.SumRequest) (response *calc_pb.SumResponse, err error) {
	num1 := request.GetSumMessage().GetNum1()
	num2 := request.GetSumMessage().GetNum2()
	sum := num1 + num2
	res := calc_pb.SumResponse{
		Result: sum,
	}

	return res

}

func main() {
	fmt.Printf("Calculator server is starting ...")

	lis, err := net.Listen("tcp", "0.0.0.0:50002")

	if err != nil {
		log.Fatal("Server : Faild to listen - %v", err)
	}

	s := grpc.NewServer()

	calc_pb.RegisterCalculatorServiceServer(s, &calcServer{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("Server : failed to serve - %v", err)
	}
}
