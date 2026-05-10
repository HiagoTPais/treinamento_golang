package main

import "fmt"

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func insert(node *TreeNode, data int) *TreeNode {
	if node == nil {
		return &TreeNode{data: data}
	}

	if data <= node.data {
		node.left = insert(node.left, data)
	} else {
		node.right = insert(node.right, data)
	}

	return node
}

func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.data)
	preOrder(node.left)
	preOrder(node.right)
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}

	inOrder(node.left)
	fmt.Printf("%d ", node.data)
	inOrder(node.right)
}

func postOrder(node *TreeNode) {
	if node == nil {
		return
	}

	postOrder(node.left)
	postOrder(node.right)
	fmt.Printf("%d ", node.data)
}

func main() {
	root := insert(nil, 12)

	insert(root, 15)
	insert(root, 11)
	insert(root, 9)
	insert(root, 1)

	preOrder(root)
	inOrder(root)
	postOrder(root)
}
