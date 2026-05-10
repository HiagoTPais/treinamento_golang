package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Heranca")

	p1 := pessoa{"Joao", "Pedro", 20, 180}
	fmt.Println(p1)

	el := estudante{p1, "Engenharia", "USP"}
	fmt.Println(el.altura)
}
