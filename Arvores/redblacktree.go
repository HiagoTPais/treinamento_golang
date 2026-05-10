package main

import (
	"fmt"
)

const (
	RED   = true
	BLACK = false
)

type Node struct {
	key    int
	color  bool
	left   *Node
	right  *Node
	parent *Node
}

type RBTree struct {
	root *Node
}

func newNode(key int) *Node {
	return &Node{key: key, color: RED}
}

func (t *RBTree) leftRotate(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *RBTree) rightRotate(x *Node) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

func (t *RBTree) Insert(key int) {
	n := newNode(key)
	t.insertNode(n)
	t.fixInsert(n)
}

func (t *RBTree) insertNode(z *Node) {
	y := (*Node)(nil)
	x := t.root

	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.parent = y
	if y == nil {
		t.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
}

func (t *RBTree) fixInsert(z *Node) {
	for z != t.root && z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rightRotate(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

func (t *RBTree) InOrder(n *Node) {
	if n != nil {
		t.InOrder(n.left)
		fmt.Printf("%d(%s) ", n.key, colorString(n.color))
		t.InOrder(n.right)
	}
}

func colorString(c bool) string {
	if c == RED {
		return "R"
	}
	return "B"
}

func main() {
	tree := &RBTree{}
	valores := []int{10, 20, 30, 15, 25, 5, 1}

	for _, v := range valores {
		tree.Insert(v)
	}

	fmt.Println("Árvore em ordem (valor(cor)):")
	tree.InOrder(tree.root)
	fmt.Println()
}
