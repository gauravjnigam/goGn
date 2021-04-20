package main

import (
	"fmt"
	"sync"
	"time"
)

type Host struct {
	semaphore chan int
}

func (host *Host) grantPerms() {
	host.semaphore <- 1
}

func (host *Host) releasePerms() {
	<-host.semaphore
}

var wg sync.WaitGroup

func main() {
	maxGoroutines := 2
	guard := make(chan int, maxGoroutines)
	host := Host{semaphore: guard}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Printf("called - %d \n", i)
		go func(n int) {
			worker(n, &host)

		}(i)
	}

	wg.Wait()
}

func worker(i int, host *Host) {

	for j := 0; j < 3; j++ {
		//wg.Add(1)
		host.grantPerms()
		fmt.Println("doing work on", i)
		time.Sleep(1000 * time.Millisecond)
		host.releasePerms()
		//wg.Done()
	}

	wg.Done()
}
