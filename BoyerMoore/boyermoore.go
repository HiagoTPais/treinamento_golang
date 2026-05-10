package main

import (
	"fmt"
)

func badCharTable(pattern string) map[byte]int {
	table := make(map[byte]int)
	for i := 0; i < len(pattern); i++ {
		table[pattern[i]] = i
	}
	return table
}

func boyerMoore(text, pattern string) {
	n := len(text)
	m := len(pattern)

	if m == 0 {
		fmt.Println("Padrão vazio.")
		return
	}

	badChar := badCharTable(pattern)

	s := 0

	for s <= n-m {
		j := m - 1

		for j >= 0 && pattern[j] == text[s+j] {
			j--
		}

		if j < 0 {
			fmt.Printf("Padrão encontrado na posição %d\n", s)
			if s+m < n {
				s += m - badChar[text[s+m-1]]
			} else {
				s += 1
			}
		} else {
			charInText := text[s+j]
			shift, exists := badChar[charInText]
			if exists {
				s += max(1, j-shift)
			} else {
				s += j + 1
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	text := "ABAAABCD"
	pattern := "ABC"

	boyerMoore(text, pattern)
}
