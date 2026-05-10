package main

import (
	"fmt"
	"strings"
)

func RomanToInt(roman string) int {
	vals := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	r := strings.ToUpper(roman)

	total := 0
	for i := 0; i < len(r); i++ {
		v := vals[r[i]]
		if i+1 < len(r) && vals[r[i+1]] > v {
			total -= v
		} else {
			total += v
		}
	}

	return total
}

func main() {
	examples := []string{
		"MCMXC",
		"MMVIII",
		"MDCLXVI",
		"IV",
		"XLII",
		"",
	}

	for _, s := range examples {
		fmt.Printf("%-8s => %d\n", s, RomanToInt(s))
	}
}
