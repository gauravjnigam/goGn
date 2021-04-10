package main

import (
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client")

	conn, err := grpc.Dial("localhost:50001", grpc.withInsecure())
}
