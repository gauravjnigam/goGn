package main

import (
	"fmt"
	"sync"
	"time"
)

/*
A race condition is when two or more routines have access to the same resource,
such as a variable or data structure and attempt to read and write to that resource without any regard to the other routines.

we will use an example of shared resource counter(i.e. integer value)
we will have 2 go rountines which will update the counter (i.e. shared resource)
result would not be expected due to race condition with shared variable


*/
var wg sync.WaitGroup

func main() {
	// lets take a shared variable counter and initialize with 0.
	Counter := 0

	// we will use 2 go routines
	wg.Add(2)

	// will use loop for starting 2 go routines
	for i := 0; i < 2; i++ {
		// all routines would run in parallel independently and update the counter by 1
		go func() {
			// each go routines updating 10 times in loop
			for j := 0; j < 10; j++ {
				value := Counter
				time.Sleep(1 * time.Nanosecond) // added sleep to create race condition
				value++
				Counter = value
			}
			wg.Done()
		}()
	}

	wg.Wait()

	// final value should be 20.. but you might see different value like 10, 11 or any random value due to race condition with shared variable
	fmt.Println("Counter value = ", Counter)

	fmt.Println("Main program exiting ... ")
}
