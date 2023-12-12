package main

import "fmt"

func main() {

	c := make(chan int)

	//Send to channel
	go foo(c)

	//Receoive for channel
	bar(c)

	fmt.Println("About to exit")
}

func foo(c chan<- int) {
	c <- 45
}

func bar(c <-chan int) {
	fmt.Println(<-c) //45
}
