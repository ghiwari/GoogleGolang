package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("BCRYPT")

	password := "Welcome@123"
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(password)
	fmt.Println(bs)
	fmt.Println(string(bs))

	//Compare Hashed password

	err = bcrypt.CompareHashAndPassword(bs, []byte(password))
	if err != nil {
		fmt.Println(err)
	}

}
