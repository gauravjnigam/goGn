package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

const (
	PHILOS_EATING_COUNT = 3
)

type Host struct {
}

func (host *Host) grantPerms() {

}

type Philosopher struct {
	philosId        int
	leftCS, rigthCS *Chopstick
}

func (philos *Philosopher) Eat() {
	for i := 0; i < PHILOS_EATING_COUNT; i++ {
		philos.leftCS.Lock()
		philos.rigthCS.Lock()

		fmt.Printf("%d Eating ...\n", philos.philosId)
		time.Sleep(100 * time.Millisecond)
		philos.leftCS.Unlock()
		philos.rigthCS.Unlock()

	}
	wg.Done()
}

type Chopstick struct {
	sync.Mutex
}

var wg sync.WaitGroup

func main() {
	fmt.Println("Dining philosophers .. Starting... ")

	// create chops stics
	chops := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chops[i] = new(Chopstick)
	}
	// create 5 philosophers

	philos := make([]Philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = Philosopher{i, chops[i], chops[(i+1)%5]}
	}

	// create Host

	// start the Dining
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].Eat()
	}

	wg.Wait()
	fmt.Println("Dining philosophers .. Ending here... ")

}
