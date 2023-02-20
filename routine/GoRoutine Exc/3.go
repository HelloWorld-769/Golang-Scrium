package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	counter := 0
	mut := sync.Mutex{}
	wg.Add(1)

	for i := 0; i < 10; i++ {
		go func() {
			mut.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mut.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
