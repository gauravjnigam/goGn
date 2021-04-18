package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// var acceleration float64
	// var initialVelocity float64
	// var displacement float64

	fmt.Print("Enter values for acceleration, initial velocity, and initial displacement : \n")
	input_scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Acceleration ")
	input_scanner.Scan()
	a, _ := strconv.ParseFloat(input_scanner.Text(), 64)

	fmt.Print("initialVelocity ")
	input_scanner.Scan()
	v0, _ := strconv.ParseFloat(input_scanner.Text(), 64)

	fmt.Print("displacement ")
	input_scanner.Scan()
	s0, _ := strconv.ParseFloat(input_scanner.Text(), 64)

	for {
		fmt.Print("Enter value of time: ")
		input_scanner.Scan()
		time, _ := strconv.ParseFloat(input_scanner.Text(), 64)

		fn := GenDisplaceFn(a, v0, s0)

		fmt.Printf("computed displacment - %f\n", fn(time))
	}
}

func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}
