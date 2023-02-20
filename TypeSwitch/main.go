// Type Switchess
package main

import "fmt"

func do(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Int Type..")
	case string:
		fmt.Println("String Type..")
	default:
		fmt.Println("Unknown Type..")

	}
}
func main() {
	do(21)
	do("Hello")
	do(false)
}
