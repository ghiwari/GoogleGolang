package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

type ByAge []person

func (p ByAge) Len() int           { return len(p) }
func (p ByAge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByAge) Less(i, j int) bool { return p[i].age < p[j].age }

type ByName []person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].name < bn[j].name }

func main() {
	fmt.Println("Sorting Slices")

	x := []int{2, 6, 1, 4, 3, 5}
	sort.Ints(x)
	fmt.Println(x)

	s := []string{"Theta", "Delta", "Beeta", "Alpha", "Gamma"}
	sort.Strings(s)
	fmt.Println(s)

	//Custom sort of structs

	people := []person{
		{"Aishwarya", 26},
		{"Shashikala", 53},
		{"Prashant", 32},
		{"Sujata", 29},
	}
	//var peoplecopy []person
	//peoplecopy = make([]person, len(people))
	//copy(peoplecopy, people)
	peoplecopy := []person{}
	peoplecopy = append(peoplecopy, people...)

	fmt.Println("------------ByAge----------------")
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	fmt.Println("------------ByName----------------")
	fmt.Println(peoplecopy)
	sort.Sort(ByName(peoplecopy))
	fmt.Println(peoplecopy)

}
