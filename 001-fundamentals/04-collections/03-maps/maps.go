package main

import "fmt"

func main() {
	fmt.Println("MAPS Example")

	m := map[string]int{}
	m["hello"] = 21
	fmt.Println(m["hello"])

	var m1 map[string]int
	m1 = m
	fmt.Println(m1)
	//m1["hello"] = 21
	fmt.Println(m1["hello"])

	var m2 map[string]int
	m2 = make(map[string]int)
	m2["hi"] = 22
	m2["hey"] = 33
	m2["hello"] = 44
	fmt.Println(m2["hi"])

	v, ok := m2["hello"]

	fmt.Println(v, ok) //0, false

	v, ok = m2["hi"]

	fmt.Println(v, ok) //22 true

	fmt.Println("--------------for loop---------------")

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	fmt.Println("-------------------Deleting from Map-----------------")
	fmt.Println(m2) //map[hello:44 hey:33 hi:22]
	delete(m2, "hi")
	fmt.Println(m2) //map[hello:44 hey:33]

	delete(m2, "Hiiii")
	fmt.Println(m2) //map[hello:44 hey:33]

}
