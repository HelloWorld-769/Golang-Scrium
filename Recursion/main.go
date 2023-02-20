package main

import (
	"fmt"
	"strconv"
)

func countDown(num int) {
	if num == 0 {
		return
	}
	fmt.Println(num)

	countDown(num - 1)

}
func sum(num int) int {

	if num == 0 {
		return 0
	}

	return num + sum(num-1)
}
func main() {
	// var num rune = 4
	// fmt.Printf("%T", num)

	s1 := "u263a"
	fmt.Println(s1)

	s := strconv.Quote(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s)

	// fmt.Print("Enter a number:")
	// fmt.Scanln(&num)

	//fmt.Println(sum(num))

	//countDown(num)
}
