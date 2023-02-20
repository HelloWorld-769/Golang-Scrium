package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mut sync.RWMutex
var wait sync.WaitGroup
var rwlock sync.RWMutex

func main() {
	// basics()
	readWrite()
}
func readWrite() {
	go write()
	go read()
	go read()

	go read()
	time.Sleep(5 * time.Second)
	fmt.Println("Done wiuth reading and writing")
}

func read() {

	rwlock.RLock()
	defer rwlock.RUnlock()

	fmt.Println("Read Locking...")
	time.Sleep(1 * time.Second)

	fmt.Println("Reading Unlocked...")

}

func write() {

	rwlock.Lock()
	defer rwlock.Unlock()

	fmt.Println("Write Locking...")
	time.Sleep(1 * time.Second)

	fmt.Println("Writing Unlocked...")

}
