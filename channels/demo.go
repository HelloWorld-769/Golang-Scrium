package main

import (
	"fmt"
	//"internal/cpu"
)

func fibbon(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

// func main() {
// 	c := make(chan int)
// 	quit := make(chan int)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(<-c)
// 		}
// 		quit <- 0
// 	}()
// 	fibbon(c, quit)
// }

func calSq(num int, sqAns chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num = num / 10
	}
	sqAns <- sum
}
func calCube(num int, cubeAns chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num = num / 10
	}
	cubeAns <- sum
}

// func main() {
// 	number := 589
// 	sqrch := make(chan int)
// 	cubech := make(chan int)

// 	go calSq(number, sqrch)
// 	go calCube(number, cubech)

// 	//squares, cubes := <-sqrch, <-cubech
// 	fmt.Println("Final ouput is:", <-sqrch+<-cubech)

// }

func main() {
	ch := make(chan int, 1)
	ch1 := make(chan int, 1)

	ch <- 1
	ch1 <- 2

	if <-ch == <-ch1 {
		fmt.Println("same")
	} else {
		fmt.Println("sadas")
	}
}
