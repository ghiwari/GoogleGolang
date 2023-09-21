package main

import "fmt"

var x int

func main() {
	fmt.Println("------Scopes--------")

	{
		y := 2
		fmt.Println(y)
		x++
	}
	//fmt.Println(y) // THis will give error. scope of y is inside that block only
	fmt.Println(x)
	x++
	fmt.Println(x)
	bar()
	fmt.Println(x)
	x++
	fmt.Println(x)

	//Closure

	a := increment(100)
	b := increment(200)
	fmt.Println(a()) //101
	fmt.Println(a()) //102
	fmt.Println(a()) //103
	fmt.Println(a()) //104
	fmt.Println(b()) //201
	fmt.Println(b()) //202
}

func bar() {
	x++
}

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}

}
