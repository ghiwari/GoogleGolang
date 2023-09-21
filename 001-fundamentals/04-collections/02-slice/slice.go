package main

import "fmt"

func main() {
	fmt.Println("Slice")

	sl := []int{3, 5, 7, 9, 11, 15}

	for i, v := range sl {
		fmt.Println(i, v)
	}

	//Slicing the slice
	fmt.Println("----------------------SLICING the slice-------------")
	fmt.Println("sl[:]", sl[:])
	fmt.Println("sl", sl)
	fmt.Println("------------------------")
	fmt.Println("sl[:5]", sl[:5])
	fmt.Println("sl", sl)
	fmt.Println("------------------------")
	fmt.Println("sl[1:4]", sl[1:4])
	fmt.Println("sl", sl)
	fmt.Println("------------------------")
	fmt.Println("sl[1:]", sl[1:])
	fmt.Println("sl", sl)
	fmt.Println("----------Slicing with capacity--------------")
	sl1 := make([]int, 5, 10)
	sl1[0] = 10
	sl1[1] = 20
	sl1[2] = 30
	sl1[3] = 40
	sl1[4] = 50
	fmt.Println(sl1, len(sl1), cap(sl1))

	fmt.Println("sl1[:]", sl1[:], len(sl1), cap(sl1))
	fmt.Println("sl1", sl1, len(sl1), cap(sl1))
	fmt.Println("------------------------")
	fmt.Println("sl1[:5]", sl1[:5], len(sl1), cap(sl1))
	fmt.Println("sl1", sl1, len(sl1), cap(sl1))
	fmt.Println("------------------------")
	fmt.Println("sl1[1:4]", sl1[1:4], len(sl1), cap(sl1))
	fmt.Println("sl1", sl1, len(sl1), cap(sl1))
	fmt.Println("------------------------")
	fmt.Println("sl1[1:]", sl1[1:], len(sl1), cap(sl1))
	fmt.Println("sl1", sl1, len(sl1), cap(sl1))

	fmt.Println("------------------------")
	sl1 = sl1[3:]
	fmt.Println("sl1", sl1, len(sl1), cap(sl1))

	//Append slice to slice

	fmt.Println("---------APPEND-------------")
	fmt.Println(sl)
	sl = append(sl, 45, 55)
	fmt.Println("After appending 2 elements - ", sl)
	x := []int{66, 77, 88, 99}
	sl = append(sl, x...) //Need to add this ... variadic parameter while sending slice to any function
	fmt.Println("After appending whole slice - ", sl)

	fmt.Println("----------DELETing from SLICE")

	fmt.Println("Original - ", sl)       //[3 5 7 9 11 15 45 55 66 77 88 99]
	sl = append(sl[:5], sl[8:]...)       // Deleting 15, 45,55
	fmt.Println("After Deleting - ", sl) //[3 5 7 9 11 66 77 88 99]

	fmt.Println("--------------MAKE function---------------")

	s := make([]int, 5, 10)
	fmt.Println(s, len(s), cap(s))
	z := 20

	//Here capacity grows dynamically from 10 to 20 and then to 40 and so on (doubles)
	for z < 40 {
		s = append(s, z)
		fmt.Println(len(s), cap(s), z, s)
		z++
	}

	//Multidimentional slices

	fmt.Println("-----------MULTIDIMENTIONAL SLICE----------")

	x1 := []string{"alpha", "beeta", "gamma"}
	fmt.Println("x1", x1)
	x2 := []string{"delta", "theta", "geta"}
	fmt.Println("x2", x2)

	xy := [][]string{x1, x2}
	fmt.Println("xy", xy)
	fmt.Println(xy[1][1])
}
