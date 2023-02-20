package main

import (
	"fmt"
)

// func main() {
// 	a := [3][3]int{
// 		{1, 0, 1},
// 		{0, 1, 0},
// 		{0, 1, 0},
// 	}

// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			//right diagonal
// 			if a[i][j] == a[i+1][j+1] && a[i+1][j+1] == a[i+2][j+2] && a[i][j] == 1 && i == 0 {
// 				fmt.Println("1 wins")
// 				return
// 			}
// 			if a[i][j] == a[i+1][j+1] && a[i+1][j+1] == a[i+2][j+2] && a[i][j] == 0 {
// 				fmt.Println("0 wins")
// 				return
// 			}

// 		}
// 	}

// 	//rows change
// 	for i := 0; i < 3; i++ {
// 		if a[i][0] == a[i][1] && a[i][1] == a[i][2] && a[i][0] == 1 {
// 			if a[i][0] == 1 {
// 				fmt.Println("1 wins")
// 				return
// 			} else {
// 				fmt.Println("0 wins")
// 				return
// 			}
// 		}
// 	}

// 	// columns change
// 	for j := 0; j < 3; j++ {
// 		if a[0][j] == a[1][j] && a[1][j] == a[2][j] && a[0][j] == 1 {
// 			if a[0][j] == 1 {
// 				fmt.Println("1 wins")
// 				return
// 			} else {
// 				fmt.Println("0 wins")
// 				return
// 			}
// 		}
// 	}

// 	fmt.Println("Draw")
// 	return

// }

// func main() {
// 	//arr := []byte{'c', 'd', 'b', 'd', 'a', 'a', 'd', 'b', 'd', 'c', 'e', 'c', 'h', 'e', 'b', 'c', 'f', 'a', 'e', 'd', 'b', 'b', 'b', 'h', 'f', 'f', 'b', 'h', 'g', 'e', 'g', 'g', 'b', 'c', 'b', 'c', 'h', 'z', 'b', 'h', 'a', 'b', 'e', 'a', 'c', 'e', 'd', 'h', 'c', 'b', 'h', 'h', 'c', 'b', 'f', 'h', 'e', 'c', 'c', 'a', 'h', 'g', 'a', 'g', 'a', 'e', 'c', 'f', 'h', 'q', 'c', 'd', 'b', 'f', 'g', 'c', 'e', 'f', 'a', 'e', 'f', 'f', 'a', 'g', 'g', 'g', 'g', 'h', 'o', 'a', 'a', 'e', 'b', 'g', 'f', 'c', 'g', 'a', 'a', 'h', 'a', 'f', 'f'}

// 	var a int32 = 5
// 	v1 := 'v'
// 	fmt.Println(v1 + a)
// 	// mp := make(map[string]int)

// 	// for _, val := range arr1 {
// 	// 	mp[val]++
// 	// }
// 	// //fmt.Print(mp)
// 	// for key, val := range mp {
// 	// 	fmt.Print(key + ":")
// 	// 	fmt.Print(val)
// 	// 	fmt.Println()
// 	// }

// }

func sqrt(x int) int {
	i := 1
	res := 1

	for res <= x {
		i++
		res = i * i
	}
	return i - 1

}
func main() {
	var num int
	fmt.Println("Enter a number whose sqrt is to be found:")
	fmt.Scanln(&num)
	fmt.Println(sqrt(num))

	// var a *[2]int
	// fmt.Println(a)
}
