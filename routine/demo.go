package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan string)
	input := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		input <- "1"
		close(input)
		wg.Done()
	}()

	go func() {
		c <- <-input
		close(c)
		wg.Done()
	}()

	for val := range c {
		fmt.Println("value c: ", val)
	}
	for val := range input {
		fmt.Println("value input: ", val)
	}

	defer wg.Wait()
}
