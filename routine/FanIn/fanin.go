package main

import (
	"fmt"
)

/*
func main() {
	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(eve, odd)

	go receive(eve, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("About to exit....")
}

func send(eve chan<- int, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(odd)
	close(eve)
}

func receive(eve <-chan int, odd <-chan int, fanin chan<- int) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for v := range eve {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}*/

//Another example

/* In this example, we create two input channels (c1 and c2) and one output channel (out). We then launch two goroutines that write integers to the input channels and close them when they're done.

We also launch another goroutines that read from the input channels and write to the output channel.

*/
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	out := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c1 <- i
		}
		close(c1)
	}()

	go func() {
		for i := 5; i < 10; i++ {
			c2 <- i
		}
		close(c2)
	}()

	//merge the data on channels 1 and c2 into out channel
	go func() {
		for n := range c1 {
			out <- n
		}
	}()

	go func() {
		for n := range c2 {
			out <- n
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-out)
	}

	fmt.Println("Done!!")
}
