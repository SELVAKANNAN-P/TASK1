package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Append(data int) {
	newNode := &Node{Data: data, Next: nil}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (list *LinkedList) Display() {
	current := list.Head

	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	linkedList := LinkedList{}

	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	linkedList.Display()
}
