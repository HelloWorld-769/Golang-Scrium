package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//not understood
type Job struct {
	id       int
	randomNO int
}
type Result struct {
	job         Job
	sumOfDigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	num := number
	for num != 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}

	return sum
}
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		op := Result{job, digits(job.randomNO)}
		results <- op
	}
	wg.Done()
}
func createWorkerPool(noOfWorker int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorker; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	close(results)
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		job := Job{i, rand.Intn(999)}
		jobs <- job
	}

	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomNO, result.sumOfDigits)
	}
	done <- true
}
func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
