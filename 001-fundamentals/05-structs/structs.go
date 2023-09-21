package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	fmt.Println("-----STRUCTS--------")
	p := person{"Prashant", "Ghiwari"}
	fmt.Println(p)       //{Prashant Ghiwari}
	fmt.Println(p.fname) //Prashant

	//Embedded structs

	sa := secretAgent{
		person: person{"Prashant", "Ghiwari"},
		ltk:    true,
	}

	fmt.Println(sa) //{{Prashant Ghiwari} true}
	fmt.Println(sa.fname)
	fmt.Println(sa.lname)
	fmt.Println(sa.ltk)

	//Ananymous Struct

	fmt.Println("Ananymous struct")

	pa := struct {
		name string
		age  int
	}{
		"Prashant Ghiwari",
		32,
	}
	fmt.Println(pa) //{Prashant Ghiwari 32}

}
