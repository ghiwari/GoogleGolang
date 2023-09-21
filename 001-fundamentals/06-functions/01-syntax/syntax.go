package main

import "fmt"

func main() {
	fmt.Println("Functions")
	foo()
	bar("Prashant")
	s := next("Next")
	fmt.Println(s)
}

func foo() {
	fmt.Println("I am in foo")
}

func bar(st string) {
	fmt.Println("Hey, ", st)
}

func next(st string) string {
	return "New " + st
}
