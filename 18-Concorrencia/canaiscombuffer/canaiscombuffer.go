package main

import "fmt"

func main() {
	canal := make(chan string)
	canal <- "Ola mundo"
	canal <- "Programando em Go"
	canal <- "Programando em Go 2"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
