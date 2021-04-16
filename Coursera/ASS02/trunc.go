package main

import (
	"fmt"
)

func main() {
	var fpNum float32
	fmt.Print("Enter floating point number - ")
	fmt.Scan(&fpNum)

	fmt.Printf("User entered num %f", fpNum)
	fmt.Println()

	f := fmt.Sprintf("%.2f", fpNum)

	fmt.Printf("truncated num (with precision 2 ) : %s", f)

}
