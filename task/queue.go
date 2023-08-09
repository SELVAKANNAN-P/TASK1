package main

import (
	"fmt"
)

type Queue []int

func (q *Queue) Enq(item int) {
	*q = append(*q, item)
}

func (q *Queue) Deq() int {
	if len(*q) == 0 {
		return 0
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func main() {
	var q Queue

	q.Enq(1)
	q.Enq(2)
	q.Enq(3)

	fmt.Println("Dequeued:", q.Deq())
	fmt.Println("Dequeued:", q.Deq())
	fmt.Println("Dequeued:", q.Deq())
	fmt.Println("Dequeued:", q.Deq())
}
