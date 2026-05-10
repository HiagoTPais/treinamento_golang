package main

import "fmt"

func funcao1() {
	fmt.Println("funcao 1")
}

func funcao2() {
	fmt.Println("funcao 2")
}

func main() {
	defer funcao1()
	funcao2()
}
