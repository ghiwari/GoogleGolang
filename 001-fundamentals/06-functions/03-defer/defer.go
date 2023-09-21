package main

import "fmt"

func main() {

	defer fmt.Println("I should be last statement")
	defer fmt.Println("I should be last but one")
	foo()
	defer bar() // I want to print before those 2 print statements
	coffee()
	tea()

	/*
		Output :
			foo
			coffee
			tea
			bar
			I should be last but one
			I should be last statement
	*/

}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func coffee() {
	fmt.Println("coffee")
}

func tea() {
	fmt.Println("tea")
}
