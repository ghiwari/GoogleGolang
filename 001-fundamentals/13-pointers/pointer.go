package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	a := 42
	x := &a
	fmt.Printf("%T, %T\n", a, x) //int, *int
	fmt.Println(a, &a)           //42 0xc00001a0e0
	fmt.Println(x, *x)           //0xc00001a0e0 42
	fmt.Println(*&a)             //42

	*x = 43
	fmt.Println(a) //43

	//When to use pointers ? - Ans : When we haev huge data to traverse in program
	//and you want to modify value of variable in another func
	//Lets take an example with varoiable x and try changing value in other func

	p := 20
	changeP(p)
	fmt.Println(p) //20 it doesnt change to 25 because GO is pass by value

	changePWithPointer(&p)
	fmt.Println(p) //25

}

func changeP(x int) {
	x = 25
}

func changePWithPointer(x *int) {
	*x = 25
}
