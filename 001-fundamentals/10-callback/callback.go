package main

import "fmt"

func main() {
	fmt.Println("Callback Example")

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(sum(a...)) //45

	fmt.Println(evenSum(sum, a...)) //20
}

func sum(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}

func evenSum(f func(n ...int) int, n ...int) int {
	var ea []int

	for _, v := range n {
		if v%2 == 0 {
			ea = append(ea, v)
		}
	}
	return f(ea...)
}
