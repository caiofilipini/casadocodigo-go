package main

import (
	"fmt"
	"sort"
)

func main() {
	quadrados := make(map[int]int, 15)

	for n := 1; n <= 15; n++ {
		quadrados[n] = n * n
	}

	numeros := make([]int, 0, len(quadrados))

	for n := range quadrados {
		numeros = append(numeros, n)
	}
	sort.Ints(numeros)

	for _, numero := range numeros {
		quadrado := quadrados[numero]
		fmt.Printf("%d^2 = %d\n", numero, quadrado)
	}
}
