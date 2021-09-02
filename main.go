package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	root *Node
}

type Node struct {
	// Has to be comparable
	key   int
	left  *Node
	right *Node
}

// Tree
func (t *Tree) insert(value int) {
	if t.root == nil {
		t.root = &Node{key: value}
	} else {
		t.root.insert(value)
	}
}

func (t *Tree) remove(value int) {
	if t.root == nil {
		return
	}

	t.root.remove(t.root, value)
}

// Node
func (n *Node) insert(value int) {
	if value <= n.key {
		if n.left == nil {
			n.left = &Node{key: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: value}
		} else {
			n.right.insert(value)
		}
	}
}

func (n *Node) remove(node *Node, value int) *Node {
	switch {
	case node == nil:
		return nil
	case value < node.key:
		node.left = n.remove(node.left, value)
		return node
	case value > node.key:
		node.right = n.remove(node.right, value)
		return node
	case node.left == nil && node.right == nil:
		node = nil
		return nil
	case node.left == nil:
		node = node.right
		return node
	case node.right == nil:
		node = node.left
		return node
	}

	// Start on the left most right side
	lmrs := node.right

	for {
		//find smallest value on the right side
		if lmrs != nil && lmrs.left != nil {
			lmrs = lmrs.left
		} else {
			break
		}
	}

	node.key = lmrs.key
	node.right = n.remove(node.right, node.key)
	return node
}

// Good for copying the tree
func preOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.key)
		preOrder(n.left)
		preOrder(n.right)
	}
}

// Good for deleting the tree
func postOrder(n *Node) {
	if n == nil {
		return
	} else {
		postOrder(n.left)
		postOrder(n.right)
		fmt.Printf("%d ", n.key)
	}
}

// Sorted
func inOrder(n *Node) {
	if n == nil {
		return
	} else {
		inOrder(n.left)
		fmt.Printf("%d ", n.key)
		inOrder(n.right)
	}
}

func reverseOrder(n *Node) {
	if n == nil {
		return
	} else {
		reverseOrder(n.right)
		fmt.Printf("%d ", n.key)
		reverseOrder(n.left)
	}
}

func main() {
	var t Tree

	for i := 0; i < 30; i++ {
		t.insert(rand.Intn(256))
	}

	fmt.Println("Pre: ")
	preOrder(t.root)
	fmt.Println()

	fmt.Println("Post: ")
	postOrder(t.root)
	fmt.Println()

	fmt.Println("Sorted: ")
	inOrder(t.root)
	fmt.Println()

	fmt.Println("Reverse Sorted: ")
	reverseOrder(t.root)
	fmt.Println()

	fmt.Println("Removed: ")
	// These only exist because we didn't seed the random number generator
	// above.
	t.remove(15)
	t.remove(248)

	// Doesn't exist
	t.remove(1111)

	inOrder(t.root)
	fmt.Println()
}
