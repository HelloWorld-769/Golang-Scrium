package main

import (
	"fmt"
	"time"
)

// type myString string

// func (m myString) add1(i int) string {
// 	fmt.Println(i)
// 	return "hello"
// }

// func main() {

// 	//fmt.Println(m + "hello")
// 	var m myString
// 	i := 5
// 	a := myString.add1
// 	fmt.Println(a)                   //gives address
// 	fmt.Println(myString.add1(m, i)) // we need to pass m

// 	fmt.Println(m.add1(i)) //no neeed to pass ma
// }

//Slice capacity code..
// func main() {

// 	var s []int
// 	printSlice(s)

// 	// append works on nil slices.
// 	s = append(s, 0)
// 	printSlice(s)

// 	// The slice grows as needed.
// 	s = append(s, 1)
// 	printSlice(s)

// 	// We can add more than one element at a time.

// 	s = append(s, 2, 3, 4, 41)
// 	printSlice(s)

// 	// s = append(s, 2)
// 	// printSlice(s)

// 	// s = append(s, 3)
// 	// printSlice(s)

// 	// s = append(s, 4)

// 	s = append(s, 2, 10, 9, 5, 7, 4, 7, 100000)
// 	printSlice(s)

// 	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
// 	printSlice(s)

// 	// s := make([]int, 0, 1)
// 	// printSlice(s)
// 	// s = append(s, 1, 2, 3)
// 	// printSlice(s)

// 	// s = append(s, 1, 2, 3, 4)
// 	// printSlice(s)

// }
// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }

// func main() {
// 	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

// 	keys := make([]string, 0, len(m))
// 	for k := range m {
// 		keys = append(keys, k)
// 	}

// 	//fmt.Println(keys)
// 	sort.SliceStable(keys, func(i, j int) bool {
// 		return m[keys[i]] < m[keys[j]]
// 	})

// 	for _, k := range keys {
// 		fmt.Println(k, m[k])
// 	}

// }

/* Three dots/ellipses notation and variadic function example
   (in go command ... means all files in the directory and subdirectories)
*/

// show is a variadic function, it can have none or many arguments
// func show(s ...string) {
// 	println()
// 	for i, val := range s {
// 		fmt.Printf("%d. %s\n", i, val)
// 	}
// }
// func main() {
// 	fmt.Println("Arrays are value types, fixed-length sequences of items of the same type.")
// 	fmt.Printf("Slices are reference types, can be resized by append.\n\n")

// 	seasons := []string{"Spring", "Summer", "Autumn", "Winter"} // slice
// 	choices := [2]string{"Good", "Bad"}
// 	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"} //array

// 	dsun := make([]string, 7, 7) //slice

// 	// Let's populate dsun with Sunday first
// 	for i, d := range days {
// 		if i == 6 {
// 			dsun[0] = d
// 		} else {
// 			dsun[i+1] = d
// 		}
// 	}

// 	fmt.Println("Length of slice seasons is", len(seasons), ", capacity is ", cap(seasons))
// 	fmt.Println("Length of slice dsun is", len(dsun), ", capacity is ", cap(dsun))
// 	fmt.Println("Length of array days is", len(days))
// 	fmt.Println("Length of array choices is", len(choices))

// 	show("a", "b")
// 	show(seasons...)
// 	show(days...) //cannot use days (type [7]string) as type []string in argument to show
// 	// show(choices...) //cannot use choices (type [2]string) as type []string in argument to show
// 	show(dsun...)
// 	show("The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog")
// }

// //Passing fucntion as argument to other function
// func x(f func(x int)) {
// 	f(2)
// }

// func main() {
// 	a := func(x int) {
// 		fmt.Println(x)
// 	}
// 	x(a)
// }

//Returning function from a function

// func main() {
// 	areaF := getAreaFunc()

// 	print(3, 4, areaF)
// }

// //
// func print(x, y int, area func(int, int) int) {
// 	fmt.Printf("Area is: %d\n", area(x, y))
// }

//	func getAreaFunc() func(int, int) int {
//		return func(x, y int) int {
//			fmt.Println("zzzzzzzzzzzzzzzzz")
//			return x * y
//		}
//	}

//	func main() {
//		n := (-1) * 13
//		itob := strconv.FormatInt(int64(n), 2)
//		fmt.Print(itob)
//		//fmt.Println(-13 >> 1)
//	}
//
// function which return "geeks"
func main() {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("HelloWorld")

}
