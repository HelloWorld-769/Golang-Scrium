package main

import "fmt"

// The difference between new and make  is that this gives you a pointer, as is the case anytime you use new.

//The length of a nil slice, map or channel is 0. The capacity of a nil slice or channel is 0
//Every time you append to a slice, it gives you a new slice with a different address (though it points to same underlying array)
func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for i := range image {
		image[i] = make([]uint8, dx)
		for j := range image[i] {
			image[i][j] = uint8((i + j) / 2)
		}
	}
	return image
}

func main() {
	image := Pic(10, 10)
	fmt.Println(image)

	// sort a slice
	// https://yourbasic.org/golang/how-to-sort-in-go/#:~:text=Use%20the%20function%20sort.,of%20equal%20elements%2C%20use%20sort.
}

// func main() {
// 	a := make([]int32, 0)
// 	a = append(a, 10, 20, 30)
// 	fmt.Println(len(a), cap(a), reflect.ValueOf(a).Kind())
// }

// func main() {
//     a := new([]int32)
//     *a = append(*a, 10, 20, 30)
//     fmt.Println(len(*a), cap(*a), reflect.ValueOf(*a).Kind())
//}
