package main

import (
	"fmt"
	"sync"
)

func main() {

	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("I am Done")
}

func send(even, odd chan<- int) {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(odd)
	close(even)
}

func receive(even, odd <-chan int, fanin chan<- int) {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for e := range even {
			fanin <- e
		}
		wg.Done()
	}()

	go func() {
		for o := range odd {
			fanin <- o
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
