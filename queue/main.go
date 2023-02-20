package main

import (
	"fmt"
)

// Method 1
// func main() {
// 	//queue
// 	q := list.New()
// 	q.PushBack(10)
// 	q.PushBack(20)
// 	q.PushBack(30)

// 	//removing elements
// 	front := q.Front()
// 	fmt.Println(front.Value)
// 	//	q.Remove(front)
// 	fmt.Println(front.Next().Value)
// }

// Method 2
type Queue []int

func (q *Queue) Push(val int) {
	*q = append(*q, val)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
func main() {
	var q Queue

	q.Push(10)
	q.Push(20)
	q.Push(30)
	q.Push(40)

	fmt.Println(q)
	q.Pop()
	fmt.Println(q)

}
