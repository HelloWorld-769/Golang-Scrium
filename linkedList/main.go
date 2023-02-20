package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	val  int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) Insert(val int) {
	newNode := &Node{val: val, next: nil}
	if l.head == nil {
		l.head = newNode
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = newNode
	}
}
func (l *List) printList() {
	p := l.head
	for p.next != nil {
		fmt.Print(p.val)
		if p.next.next != nil {
			fmt.Print(" -> ")
		}
		p = p.next
	}

}

// not working
func (l *List) reverse() {
	curr := l.head
	var prev *Node = nil
	var next *Node = nil

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.head = prev
}
func main() {
	l := List{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 5; i++ {
		l.Insert(rand.Intn(100))
	}

	l.printList()

	l.reverse()
	fmt.Println()
	l.printList()

}
