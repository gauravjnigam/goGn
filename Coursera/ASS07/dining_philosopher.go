package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var count1 int = 0
var count2 int = 0

type Chops struct{ sync.Mutex }
type Philo struct {
	leftcs, rightcs *Chops
}

func (p Philo) eat(a int) {
	if count1 < 3 {
		p.leftcs.Lock()
		defer wg.Done()
		p.rightcs.Lock()
		defer wg.Done()
		fmt.Printf("starting to eat %d \n", a)
		p.rightcs.Unlock()
		p.leftcs.Unlock()
		fmt.Printf("finishing eating %d \n", a)
		count1++
	}
}
func (p Philo) eat1(a int) {
	if count2 < 3 {
		p.leftcs.Lock()
		defer wg.Done()
		p.rightcs.Lock()
		defer wg.Done()
		fmt.Printf("starting to eat %d \n", a)
		p.rightcs.Unlock()
		p.leftcs.Unlock()
		fmt.Printf("finishing eating %d \n", a)
		count2++
	}

}

func main() {

	CStck := make([]*Chops, 5)
	for i := 0; i < 5; i++ {
		CStck[i] = new(Chops)
	}

	philos := make([]*Philo, 5)
	for j := 0; j < 5; j = j + 2 {
		if j == 4 {
			philos[j] = &Philo{&Chops{}, &Chops{}}

			count1 = 0
			for count1 < 3 {
				wg.Add(2)
				go philos[j].eat(j + 1)
				wg.Wait()
			}

		} else {
			philos[j] = &Philo{&Chops{}, &Chops{}}
			philos[j+1] = &Philo{&Chops{}, &Chops{}}

			count1 = 0
			count2 = 0
			for count1 < 3 || count2 < 3 {
				wg.Add(4)
				go philos[j].eat(j + 1)
				go philos[j+1].eat1(j + 2)
				wg.Wait()

			}
		}
	}

}
