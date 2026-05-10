package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Inserir(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Inserir(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Inserir(value)
		}
	}
}

func (n *Node) EmOrdem() {
	if n == nil {
		return
	}
	n.Left.EmOrdem()
	fmt.Print(n.Value, " ")
	n.Right.EmOrdem()
}

func main() {
	raiz := &Node{Value: 10}
	raiz.Inserir(5)
	raiz.Inserir(15)
	raiz.Inserir(3)
	raiz.Inserir(7)
	raiz.Inserir(12)
	raiz.Inserir(17)

	fmt.Print("Árvore em ordem: ")
	raiz.EmOrdem()
}
