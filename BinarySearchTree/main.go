package binarysearchtree

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) insert(key int) {
	bst.root = insertNode(bst.root, key)
}

func insertNode(node *Node, key int) *Node {
	if node == nil {
		return &Node{key: key}
	}

	if key < node.key {
		node.left = insertNode(node.left, key)
	} else if key > node.key {
		node.right = insertNode(node.right, key)
	}

	return node
}

func main() {
	bst := BinarySearchTree{}
	bst.insert(10)
	bst.insert(5)
	bst.insert(20)

	fmt.Println(bst)
}
