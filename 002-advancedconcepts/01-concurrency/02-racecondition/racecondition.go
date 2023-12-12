package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Race Condition")

	var wg sync.WaitGroup
	counter := 0
	const gs int = 100

	wg.Add(gs)

	for i := 1; i <= gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter -", counter)
}
