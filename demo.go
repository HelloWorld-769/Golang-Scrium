package main

import (
	"fmt"
	"time"
)

func main() {
	tym := time.Now()

	fmt.Println(time.Until(tym))
}
