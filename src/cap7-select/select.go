package main

import "fmt"

func separar(nums []int, i, p chan<- int, pronto chan<- bool) {
	for _, n := range nums {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	pronto <- true
}

func main() {
	i, p := make(chan int), make(chan int)
	pronto := make(chan bool)
	nums := []int{1, 23, 42, 5, 8, 6, 7, 4, 99, 100}

	go separar(nums, i, p, pronto)

	var impares, pares []int
	fim := false

	for !fim {
		select {
		case n := <-i:
			impares = append(impares, n)
		case n := <-p:
			pares = append(pares, n)
		case fim = <-pronto:
		}
	}

	fmt.Printf("Ímpares: %v | Pares: %v\n", impares, pares)
}
