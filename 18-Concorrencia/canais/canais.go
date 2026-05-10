package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Ola, mundo", canal)
	fmt.Println("Despois da funcao escrever")

	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)

	}
}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)

	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
