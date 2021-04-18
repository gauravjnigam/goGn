package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

func main() {
	initial_size := 3
	var slice = make([]int, initial_size)
	printSlice(slice)
	var input string
	var idx int = 0
	for {
		fmt.Print("Enter the integer or X for exit : ")

		fmt.Scan(&input)
		if strings.Compare(input, "X") == 0 {
			fmt.Println("Exiting....")
			break
		} else {
			intVal, _ := strconv.ParseInt(input, 10, 0)
			if idx < initial_size {
				slice[idx] = int(intVal)
			} else {
				slice = append(slice, int(intVal))
			}
			sorted := make([]int, len(slice))
			copy(sorted, slice)
			sort.Ints(sorted)

			printSlice(sorted)
		}
		idx++
	}
}

func printSlice(s []int) {
	fmt.Printf("length=%d capacity=%d %v\n", len(s), cap(s), s)
}
