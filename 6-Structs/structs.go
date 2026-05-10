package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	fmt.Println("Arquivo")

	var u usuario
	u.nome = "Davi"

	fmt.Println(u)
}
