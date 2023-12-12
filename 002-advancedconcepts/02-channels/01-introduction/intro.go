package main

import "fmt"

func maian() {

	fmt.Println("Channel Introduction")

	/*
				The below lines of code will give error.
				fatal error: all goroutines are asleep - deadlock!

					ch := make(chan int)
					ch <- 42  //line No 14
						fmt.Println(<-ch)

			goroutine 1 [chan send]:
			main.main()
			        C:/Users/prashant/GolandProjects/googlesGo/002-advancedconcepts/02-channels/01-introduction/intro.go:14 +0x7a



		you can make above program run by making chan as buffered.
		And dont forgot to mention buffer size in make function
		Syntax : ch := make(chan int, buffSize)
		ch := make(chan int, 1) // this is buffered channel with size 1.
		Buffered channel holds some value in it so that some receiver will recieve it later
	*/

	//Correct way to us channel is through different go routine or buffered channel

	//First Method : use goroutine
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch) //42

	//Second Method : use buffered channel with buffer size
	//Syntax - ch := make(chan int, buffSize)
	//But buffered channel will have only 1 space
	//if you try sending 2 values then it will be deadlock
	// c2 <- 23; c2 <= 24  - this gives deadlock to resolve that give buffere size to 2 as make(chan int, 2)

	c2 := make(chan int, 1) //Buffered channel. Note the 1 in the parameter : This is buffered size

	c2 <- 23

	fmt.Println(<-c2) //23

	//Directional channels

	c3 := make(<-chan int) //receive-only type channel - only receive from channel is possible
	c4 := make(chan<- int) //Send-only type channel - sending value to channnel is possible
	_, _ = c3, c4

	//Conversion of channels between bidirectional and unidirectional

	c := make(chan int)    //Bidirectional channel
	cr := make(<-chan int) //receive-only type channel
	cs := make(chan<- int) // send-only type channel
	_, _ = cs, cr

	//Below are noit allowed
	//c = cr //cannot use cr as type chan int in assignment
	//c = cs //cannot use cs as type chan int in assignment
	//cs = cr //cannot use cr as typechan<- int in assignment

	//Below are allowed
	cs = c
	cr = c

	//conversion example
	fmt.Println((chan<- int)(c)) // converting chan c to send only type
	fmt.Println((<-chan int)(c)) // converting chan c to receive only type
}
