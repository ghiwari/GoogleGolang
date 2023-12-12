package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("WaitGroup Examle")
	wg.Add(1)
	go foo()
	bar()
	wg.Wait()
}

func foo() {
	for i := 1; i <= 5; i++ {
		fmt.Println("foo ", i)
	}
	wg.Done()
}

func bar() {
	for i := 1; i <= 5; i++ {
		fmt.Println("bar", i)
	}
}
