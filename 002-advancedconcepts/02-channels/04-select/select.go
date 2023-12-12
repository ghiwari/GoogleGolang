package main

import "fmt"

func main() {

	e := make(chan int)
	o := make(chan int)
	q := make(chan int)

	go send(e, o, q)

	receive(e, o, q)

	fmt.Println("About to exit")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Even - ", v)
		case v := <-o:
			fmt.Println("Odd - ", v)
		case v := <-q:
			fmt.Println("Quit-", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 1
}
