package anonimas

import "fmt"

func anonimas() {
	func(text string) {
		fmt.Println(text)
	}("Parametros")
}
