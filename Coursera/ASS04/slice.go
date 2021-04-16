package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	initial_slice_size := 3
	integers := make([]int64, initial_slice_size)
	fmt.Println("Please enter an integer:")
	for current_index := 0; true; current_index++ {
		var input string
		fmt.Scan(&input)

		integer_to_add, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			if strings.ToUpper(input) == "X" {
				fmt.Println("Closing program")
				return
			} else {
				fmt.Println("Invalid entry, please enter a valid integer or X to close:")
				current_index--
				continue
			}
		} else {
			if current_index < initial_slice_size {
				integers[current_index] = integer_to_add
			} else {
				integers = append(integers, integer_to_add)
			}
			sorted := make([]int64, len(integers))
			copy(sorted, integers)
			sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
			fmt.Printf("Sorted integers: %v \n", sorted)
		}
		fmt.Println("Enter another integers or X to close:")
	}
}
