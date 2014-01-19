package main

import (
	"fmt"
	"time"
)

func imprimir(n int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", n)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	go imprimir(2)
	imprimir(3)
}
