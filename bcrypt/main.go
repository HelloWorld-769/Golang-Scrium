package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("Playing with bcrypt in golang....")
	s := "password123"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	fmt.Println(string(bs))

	loginPass := "password1234"

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPass))

	if err != nil {
		fmt.Println("Loggedin failed")
		fmt.Println(err)
		return
	}

	fmt.Println("Logged in successfully..")

}
