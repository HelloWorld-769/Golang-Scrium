package main

import (
	"fmt"
	"time"

	cache "github.com/patrickmn/go-cache"
)

func main() {
	c := cache.New(5*time.Second, 10*time.Second)
	c.Set("mykey", "myvalue", cache.DefaultExpiration)
	c.Add("key", "value", cache.DefaultExpiration)
	// //c.Flush() //delete all values present in the cache
	// val, found := c.Get("mykey")
	// if found {
	// 	fmt.Println("Value found in cache:", val)
	// } else {
	// 	fmt.Println("Value not found in cache....")
	// }
	// //c.Delete("key")
	// val, found = c.Get("key")
	// if found {
	// 	fmt.Println("Value found from the cache is:", val)
	// } else {
	// 	fmt.Println("Value not added to cache.")
	// }
	start := time.Now()
	c.Add("intVal", 10, cache.DefaultExpiration)
	c.DecrementInt("intVal", 2)
	fmt.Println("Values in caches are", c.ItemCount())
	val, found := c.Get("intVal")
	if found {
		fmt.Println("Value after decrement is:", val)
	} else {
		fmt.Println("Value not found in cache..")
	}
	fmt.Println("Time taken to execute:", time.Since(start))

	//err := c.Add("intVal", S12, cache.DefaultExpiration)
	c.Set("intVal", 12, cache.DefaultExpiration)
	val, found = c.Get("intVal")
	if found {
		fmt.Println("Value after decrement is:", val)
	} else {
		fmt.Println("Value not found in cache..")
	}

	c.Set("wfb", []string{"SAdsa", "SAdsa", "SAdjjsa"}, cache.DefaultExpiration)

	val, found = c.Get("wfb")
	for i, v := range val {
		fmt.Println(v)
	}
}
