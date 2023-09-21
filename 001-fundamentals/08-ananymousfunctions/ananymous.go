package main

import "fmt"

func main() {

	func() {
		fmt.Println("Ananymous")
	}()

	func(x int) {
		fmt.Println("Ananymous ", x)
	}(22)
}
