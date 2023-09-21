package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("JSON Example")

	//MARSHALLING Example

	fmt.Println("Marshaling Example")
	p1 := person{"Prashant", 32}
	p2 := person{"Parshya", 29}

	people := []person{p1, p2}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs)) //[{"Name":"Prashant","Age":32},{"Name":"Parshya","Age":29}]

	fmt.Println("----------UnMarshal Example---------")

	s := `[{"Name":"Prashant","Age":32},{"Name":"Parshya","Age":29}]`

	var pa []person

	err = json.Unmarshal([]byte(s), &pa) //&pa is required instead of pa - Need to pass pointer

	if err == nil {
		fmt.Println(pa)
	}

}
