package main

import "fmt"

// Node represent the binary search tree components
type Node struct {
	Key   int32
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
func (n *Node) Insert(k int32) {
	// If k is larger than the current Node Key
	if k > n.Key {
		// Move to the right if it's empty
		if n.Right == nil {
			n.Right = &Node{Key: k}
		}
		// Otherwise, create new node in the right.
		n.Right.Insert(k)
	}
	// If k is smaller than the current Node Key
	if k < n.Key {
		// Move to the left if it's empty
		if n.Left == nil {
			n.Left = &Node{Key: k}
		}
		// Otherwise, create new node in the left.
		n.Left.Insert(k)
	}
}

// Search will take a key value
// and return TRUE if a node with that value is exist.
func (n *Node) Search(k int32) bool {

	searchCount++

	if n == nil {
		return false
	}

	if k > n.Key {
		// search in right node
		return n.Right.Search(k)
	}

	if k < n.Key {
		// search in left node
		return n.Left.Search(k)
	}

	return true
}

var searchCount int32

func main() {

	tree := &Node{Key: 100}
	tree.Insert(10)
	tree.Insert(75)
	tree.Insert(200)
	tree.Insert(300)

	fmt.Println(tree.Search(10))
	fmt.Println(searchCount)
}
