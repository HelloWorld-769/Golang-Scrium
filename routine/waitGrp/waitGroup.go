package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// func main() {

// 	wg.Add(10)
// 	go attack("shdvfdjshvfjh dfk sdjnfjks")
// 	wg.Wait()
// }

// func attack(txt string) {
// 	fmt.Println(txt)
// 	defer wg.Done()
// }

/*
var signals = []string{}
var waitGroup sync.WaitGroup // pointer hona chahiye wrna deadlock aa jayegay

func main() {
	// go greeter("hello")
	// greRace condition...
	websiteList := []string{"http://youtube.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://linkedin.com",
	}

	for _, web := range websiteList {

		go getStatusCode(web)

		waitGroup.Add(1) // only 1 go routine fired up!
	}
	waitGroup.Wait() // this will not allow main method to finish
	// until it recieves a response signal from the method called
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	// this will send signal to wait in the main
	defer waitGroup.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("oops endpoint", err)

	} else {
		signals = append(signals, endpoint)
		fmt.Printf("%d status code for website %s", res.StatusCode, endpoint)
		fmt.Println()
	}
}*/

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

}
