package main

import (
	"fmt"
	//"container/list"
	//"github.com/golang-collections/collections/stack"
)

// func main() {
// 	fmt.Println("Stack Implemented using predefined library")

// 	s := stack.New()

// 	fmt.Println(s.Len())
// 	for i := 0; i < 5; i++ {
// 		s.Push(i)
// 	}
// 	fmt.Println(s.Peek())

// }

// Another Method. This method can be used to implement queue as well
/*func main() {
	s2 := list.New()
	for i := 0; i < 5; i++ {
		fmt.Println("Pushing ", s2.PushBack(i).Value)
	}
	for s2.Len() > 0 {
		back := s2.Back()
		s2.Remove(back)
		fmt.Println("Popping ", back.Value)
	}
}*/

// Another Method
type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	sz := len(*s)
	val := (*s)[sz-1]
	//remove element from stack
	*s = (*s)[:sz-1]
	return val
}
func main() {
	var newStack Stack

	newStack.Push(10)

	newStack.Push(1)
	newStack.Push(100)
	newStack.Push(120)
	newStack.Push(101)

	fmt.Print(newStack)

	newStack.Pop()
	newStack.Pop()

	fmt.Print(newStack)

}
