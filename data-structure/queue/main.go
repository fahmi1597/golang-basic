package main

import "fmt"

// Queue represent a simple queue.
type Queue struct {
	items []int32
}

// Enqueue stores item to the slice
func (q *Queue) Enqueue(item int32) {
	q.items = append(q.items, item)
}

// Dequeue removes item to the slice, it follows the FIFO method
// So it will removes the item front the front.
func (q *Queue) Dequeue() int32 {
	out := q.items[0]
	q.items = q.items[1:]
	return out
}

func main() {

	q := &Queue{}

	q.Enqueue(10)
	q.Enqueue(50)
	q.Enqueue(20)

	fmt.Println(q)

	fmt.Println("Out:", q.Dequeue())
	fmt.Println(q)
}
