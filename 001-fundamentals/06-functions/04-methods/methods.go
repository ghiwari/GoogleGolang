package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Println(p.name, " Speaking")
}

func main() {

	p := person{"Prashant Ghiwari", 32}
	p.speak()

}
