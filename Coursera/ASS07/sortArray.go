package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

/*
Write a program to sort an array of integers. The program should partition the array into 4 parts,
each of which is sorted by a different goroutine. Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.

*/

var wg sync.WaitGroup

func main() {
	fmt.Println("input integers (separated by space)...")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	inputSlice := strings.Split(input, " ")
	intSlice := make([]int, len(inputSlice))

	for i, v := range inputSlice {
		intSlice[i], _ = strconv.Atoi(v)
	}
	fmt.Printf("intput array : %v\n", intSlice)

	const partitionSize int = 4

	totalPartition := (len(intSlice) / partitionSize) + 1
	wg.Add(totalPartition)
	var mapOfSlice = make(map[int][]int)

	for i := 0; i < len(intSlice); i++ {
		index := i % totalPartition
		if _, ok := mapOfSlice[index]; ok {
			mapOfSlice[index] = append(mapOfSlice[index], intSlice[i])
		} else {
			mapOfSlice[index] = make([]int, 0)
			mapOfSlice[index] = append(mapOfSlice[index], intSlice[i])
		}

	}

	for j := 0; j < totalPartition; j++ {
		go sortArr(mapOfSlice[j])
	}

	wg.Wait() // holding for all the partitioned slice sorting to complete

	// merging slices now
	for i := 1; i < totalPartition; i++ {
		mapOfSlice[0] = mergeSortedArray(mapOfSlice[0], mapOfSlice[i])
	}

	fmt.Printf("sorted array : %v \n ", mapOfSlice[0])

}

func mergeSortedArray(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

func sortArr(arr []int) {
	sort.Ints(arr)
	wg.Done()
}
