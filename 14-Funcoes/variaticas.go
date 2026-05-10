package soma

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func soma() {
	fmt.Println(soma(1, 2, 3, 4))
	escrever("Numeros: ", 12, 12, 34)
}
