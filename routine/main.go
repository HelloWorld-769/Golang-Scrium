package main

import (
	"fmt"
	"runtime"
	"sync"
)

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(3)

// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("First goroutine")
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("Second goroutine")
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("Third goroutine")
// 	}()

// 	wg.Wait()
// 	fmt.Println("All goroutines completed")
// }

// wait Group and mutex
// var signals = []string{}
// var waitGroup sync.WaitGroup // pointer
// //var mut sync.Mutex

// func main() {
// 	// go greeter("hello")
// 	// greeter("hello go")
// 	// // we never waited for the hello to be printed
// 	websiteList := []string{
// 		"http://google.com",
// 		"http://facebook.com",
// 		"http://youtube.com",
// 		"http://stackoverflow.com",
// 		"http://golang.org",
// 		"http://linkedin.com",
// 	}

// 	for _, web := range websiteList {

// 		go getStatusCode(web)

// 		waitGroup.Add(10) // only 1 go routine fired up!
// 	}
// 	waitGroup.Wait() // this will not allow main method to finish
// 	// until it recieves a response signal from the method called
// 	fmt.Println(signals)
// }

// func getStatusCode(endpoint string) {
// 	// this will send signal to wait in the main
// 	defer waitGroup.Done()

// 	res, err := http.Get(endpoint)
// 	if err != nil {
// 		fmt.Println("oops endpoint", err)

// 	} else {
// 		//mut.Lock()
// 		signals = append(signals, endpoint)
// 		//mut.Unlock()
// 		fmt.Printf("%d status code for website %s", res.StatusCode, endpoint)
// 		fmt.Println()
// 	}
// }

//race condition
//run it using go run --race .
// deadlock will occur and data race will occur in this code

/*func main() {
	fmt.Println("Race condition...")
	wg := &sync.WaitGroup{} //wait group

	wg.Add(3)
	score := []int{0}
	go func(wg *sync.WaitGroup) {
		fmt.Println("One R")

		score = append(score, 1)

		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Two R")

		score = append(score, 2)

		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Three R")
		score = append(score, 3)

		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println(score)
}*/

//solution using mutex for synchronization

/*func main() {
	fmt.Println("Race condition...")
	wg := &sync.WaitGroup{} //wait group
	mut := &sync.Mutex{}    //mutex

	wg.Add(3)
	score := []int{0}
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
*/

//Another race condtion example

func main() {
	fmt.Println("CPU's: ", runtime.NumCPU())

	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	wg := sync.WaitGroup{}
	mut := sync.Mutex{}
	counter := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mut.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mut.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())

	}

	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)

}
