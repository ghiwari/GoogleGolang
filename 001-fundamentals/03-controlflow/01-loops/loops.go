package main

import "fmt"

func main() {
	fmt.Println("Loops Example")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	a, b := 5, 20
	for a < b {
		a++
		fmt.Print(a, ",")
	}
	fmt.Println("Done")

	arr := []int{1, 2, 30, 4, 5}
	fmt.Println("---------------------------")
	for _, k := range arr {
		fmt.Println(k)
	}
}
