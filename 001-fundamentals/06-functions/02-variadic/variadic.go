package main

import "fmt"

func main() {
	fmt.Println("Variadic")
	foo(2, 4, 6, 8, 10)

	//unferling/unpack slice example

	xi := []int{10, 20, 30, 40, 50}
	foo(xi...) //Unpack slice

}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
