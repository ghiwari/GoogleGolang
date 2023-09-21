package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Println("Person ", p.name, " is speaking")
}

type agent struct {
	name string
}

func (a agent) speak() {
	fmt.Println("Agent ", a.name, " is speaking")
}

func bar(h human) {
	h.speak()

	//Assertion
	switch h.(type) {
	case person:
		fmt.Println("Person Called")
	case agent:
		fmt.Println("Agent Called")
	}
}

func main() {

	fmt.Println("Interface Example")
	var h human
	h = person{"Prashant", 32}
	h.speak()

	fmt.Println("----------------Polymorphism-----------")

	p := person{"Prashant", 32}
	a := agent{"Bond_007"}

	bar(p) //Person  Prashant  is speaking
	bar(a) //Agent  Bond_007  is speaking

}
