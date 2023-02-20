package main

import "fmt"

type person struct {
	fName string
	lName string
}

// func main() {
// 	p := person{"alex", "anderson"}

// 	p.updateName("hello")
// 	p.print()
// }
// func (p *person) updateName(updatedName string) {
// 	p.fName = updatedName
// }
// func (p person) print() {
// 	fmt.Printf("%v", p)
// }
func main() {
	s := []string{"hi", "dsbhf", "sdfsdf"}
	update(s)
	fmt.Println(s)
}
func update(s []string) {
	s[0] = "hell"
}
