package main

import (
	"fmt"
)

/*

type ErrNegativeSqrt float64

func (r ErrNegativeSqrt) add(x, y int) string {
	//z := x + y
	return fmt.Sprintf("asdasdas : %v", r)
}
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", e)
}
func sqrt(x float64) (float64, string) {
	if x < 0 {

		var err2 ErrNegativeSqrt = ErrNegativeSqrt(x)
		return 0, err2.add(1, 2)
	} else {
		return math.Sqrt(x), ""
	}
}
func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
*/
//Error recovery
func recovery() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	panic("Somthing went wrong")
}

func main() {
	recovery()
	fmt.Println("Finished with program execution")
}

//Another recover example
//advanced example..

/*
func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
*/
