package main

import "fmt"

func main() {
	fmt.Println("Function Expression")

	f := func() {
		fmt.Println("in Func Expression")
	}
	f()

	g := func(s string) {
		fmt.Println("in Func Expression ", s)
	}
	g("Hello")

	//func as return value

	fmt.Println("--------func as return value-------------")

	fmt.Println(foo())
	f1 := bar()
	fmt.Println(f1()) //2023
	//or
	fmt.Println(bar()()) //2023

}

func foo() string {
	return "Hello, World!"
}

func bar() func() int {
	return func() int {
		return 2023
	}
}
