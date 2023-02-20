package main

import (
	"fmt"
)

/*func main() {
	fmt.Println("{Playing with go channels....}")

	myChan := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	// myChan <- 5
	// fmt.Println(<-myChan)
	go func(ch chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myChan
		fmt.Println(val)
		fmt.Println(isChanelOpen)
		wg.Done()
	}(myChan, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myChan <- 5
		// close(myChan)  -> this is used the close the
		wg.Done()
	}(myChan, wg)
	wg.Wait()
}
*/

//doubt ..........
/*func main() {
	myChan := make(chan int, 1)

	go func() {
		myChan <- 45
		myChan <- 87
	}()
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)

}*/

//Directional channels

// func main() {
// 	myChan := make(chan<- int, 2) //this channel can only be used as sender not receiver..

// 	myChan <- 5
// 	myChan <- 2

// 	//fmt.Println("Val 1:", <-myChan)  //will give error

// 	/*
// 		// Receiver channel only
// 		myChan:=<- chan int
// 	*/
// }

// func main() {
// 	c := make(chan int)

// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			c <- i
// 		}
// 		close(c)
// 	}()

// 	for v := range c {
// 		fmt.Println(v)
// 	}

// 	fmt.Println("About to exit")
// }

//select in channels

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)

	receive(eve, odd, quit)

}
func receive(e, o, q <-chan int) {
	for {

		select {
		case v := <-e:
			fmt.Println("From the eve channel:", v)
		case v := <-o:
			fmt.Println("From the odd channel:", v)
		case v := <-q:
			fmt.Println("From quit channel :", v)
			return
		}

	}

}
func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- -1
}
