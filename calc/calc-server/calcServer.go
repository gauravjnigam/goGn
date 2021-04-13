package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	calc_pb "calc/calc_pb"
)

type calcServer struct{}

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
