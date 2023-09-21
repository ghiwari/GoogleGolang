package main

import "fmt"

type person struct {
	name string
	age  int
}

// This is like toString method in java
func (p person) String() string {
	return fmt.Sprintf("%s ::: %d", p.name, p.age)
}

func main() {
	fmt.Println("DataTypes")

	var a int = 20
	var b float64 = 10.5
	var c bool = true
	var d string = "Hello, World!"

	const i int = 1
	const (
		x int     = 10
		y float64 = 15.5
		z bool    = false
		s string  = "Hey Man"
	)

	const (
		j = iota * 2 //0
		k            //2
		l            //4
		m            //6
	)
	const (
		mk = iota * 2 //0
		n  = iota     // 1
		o             //2
		p             //3
	)

	const (
		_  = iota
		kb = 1 << (iota * 10)
		mb
		gb
	)
	fmt.Println(kb, mb, gb) //1024 1048576 1073741824
	fmt.Println("mk", mk)   //0

	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)
	fmt.Printf("%T, %v\n", c, c)
	fmt.Printf("%T, %v\n", d, d)

	ps := person{"Prashant", 32}
	fmt.Println(ps)
}
