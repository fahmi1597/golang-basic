package main

import "fmt"

// Stack represent a stack and stores collection of items.
type Stack struct {
	items []int32
}

// Push stores an item to the stack.
func (s *Stack) Push(item int32) {
	s.items = append(s.items, item)
}

// Pop removes an item from the stack, it follows the LIFO method
// and return the removed item.
func (s *Stack) Pop() int32 {
	out := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	fmt.Println(s)
	return out
}

// Slice cheatsheet :
// https://ueokande.github.io/go-slice-tricks/
func main() {
	s := &Stack{}
	s.Push(10)
	s.Push(30)
	s.Push(2)

	fmt.Println(s)
	fmt.Println("Item removed: ", s.Pop())
}
