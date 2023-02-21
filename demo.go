package main

import "fmt"

const (
	January int = 3 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func main() {
	fmt.Println(January, February, March, April, May, June, July, August, September, October, November, December)

}
