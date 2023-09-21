package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	breadth, height float64
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.breadth * r.height
}

func info(s shape) {
	fmt.Println("Area ", s.area())
}

func main() {

	c := circle{5}
	fmt.Printf("%T\n", c)
	info(&c)
}

/*

	//NOTE : IMPORTANT Here Receiver means eg. area function receiver parameters (c Circle)
	Receiver   Caller
	(t T)     T adn *T
	(t *T)    *T

*/
