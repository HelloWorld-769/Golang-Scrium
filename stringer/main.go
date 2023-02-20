package main

import "fmt"

//stringer is provided by fmt package

//ismein println default type se nhi chelga jo humne String () func bnaya hai uske according chlega
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

//Excercise

/*
package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	var output string = ""
	for i, v := range ip {
		if i != len(ip)-1 {
			output += strconv.Itoa(int(v)) + "."
		} else {
			output += strconv.Itoa(int(v))
		}
	}
	return output
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}


*/
