package main

import (
	"fmt"
	"sync"
)

/*
In this example, we create one input channel (in) and one output channel (out). We then launch one goroutine that writes integers to the input channel and closes it when it's done.

We also launch three goroutines that read from the input channel and write their squared value to the output channel.

Finally, we launch a fourth goroutine that reads from the output channel and prints each integer. We use time.Sleep to wait for the goroutines to finish.

*/

func main() {
	in := make(chan int)
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(4)
	//1st go routine
	go func() {
		for i := 0; i < 5; i++ {
			in <- i
		}
		close(in)
		wg.Done()
	}()

	//three go routines
	for i := 0; i < 3; i++ {
		go func() {
			for n := range in {
				out <- n * n
			}
			wg.Done()
		}()
	}

	//4th go routine
	go func() {
		for i := 0; i < 5*3; i++ {
			fmt.Println(<-out)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done!")
}
