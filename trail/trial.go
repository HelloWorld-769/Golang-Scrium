package main

import "fmt"

// func returnFunc(x string) func(y string) {
// 	return func(y string) {
// 		fmt.Println(x)
// 		fmt.Println(y)
// 	}
// }
type student struct {
	name   string
	grades []int
	age    int
}

func (s student) getAge() int {
	return s.age
}
func main() {
	s1 := student{"Tim", []int{70, 80, 90, 85}, 19}

	fmt.Println(s1.getAge())
}
