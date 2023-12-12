package main

import (
	"fmt"
	"sync"
	"time"
)

const goroutines int = 10

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanout(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
}

func populate(c1 chan int) {
	for i := 1; i <= 100; i++ {
		c1 <- i
	}
	close(c1)
}

func fanout(c1, c2 chan int) {
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for i := 1; i <= goroutines; i++ {
		go func() {
			for v := range c1 {
				func(v1 int) {
					c2 <- timeconsumingtask(v1)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeconsumingtask(v int) int {
	time.Sleep(time.Millisecond * 50)
	fmt.Println(v)
	return v * 1000
}
