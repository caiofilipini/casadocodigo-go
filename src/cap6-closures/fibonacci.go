package main

import "fmt"

func main() {
	a, b := 0, 1

	fib := func() int {
		a, b = b, a+b

		return a
	}

	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", fib())
	}

	fmt.Println()
}
