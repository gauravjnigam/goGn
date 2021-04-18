package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
order.
*/

func main() {
	//var input string
	fmt.Println("input up to 10 integer (separated by space)...")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	inputSlice := strings.Split(input, " ")
	intSlice := make([]int, len(inputSlice))

	for i, v := range inputSlice {
		intSlice[i], _ = strconv.Atoi(v)
	}
	fmt.Printf("input : %v \n", intSlice)
	BubbleSort(intSlice)
	fmt.Printf("Sorted : %v ", intSlice)

}

func BubbleSort(intSlice []int) {
	for i := 0; i < len(intSlice)-1; i++ {
		for j := 0; j < len(intSlice)-i-1; j++ {
			if intSlice[j] > intSlice[j+1] {
				swap(intSlice, j)
			}
		}
	}
}

func swap(s []int, i int) {
	var temp = s[i]
	s[i] = s[i+1]
	s[i+1] = temp
}
