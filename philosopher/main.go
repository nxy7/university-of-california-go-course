package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	number      int
	left, right *Chopstick
}

func (p *Philosopher) Eat3Times() {
	for x := 0; x < 3; x++ {
		p.Eat()
	}
}

var hostLimit = make(chan struct{}, 2)

func (p *Philosopher) Eat() {
	hostLimit <- struct{}{}
	p.left.Lock()
	p.right.Lock()
	fmt.Println("starting to eat ", p.number)

	time.Sleep(time.Millisecond)

	p.left.Unlock()
	p.right.Unlock()
	<-hostLimit
	fmt.Println("finished eating ", p.number)
}

func main() {
	_, philosophers := Init()

	wg := sync.WaitGroup{}
	for _, p := range philosophers {
		wg.Add(1)
		go func(p Philosopher) {
			p.Eat3Times()
			wg.Done()
		}(p)
	}
	wg.Wait()

}

func Init() ([]Chopstick, []Philosopher) {
	chopsticks := make([]Chopstick, 0)
	philosophers := make([]Philosopher, 0)

	for x := 0; x < 5; x++ {
		chopsticks = append(chopsticks, Chopstick{sync.Mutex{}})
	}
	for x := 0; x < 5; x++ {
		lInd := x - 1
		rInd := x

		if lInd == -1 {
			lInd = len(chopsticks) - 1
		}
		philosophers = append(philosophers, Philosopher{
			number: x + 1,
			left:   &chopsticks[lInd],
			right:  &chopsticks[rInd],
		})
	}
	return chopsticks, philosophers
}
