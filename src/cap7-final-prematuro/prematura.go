package main

import (
	"fmt"
	"time"
)

func dormir() {
	fmt.Println("Goroutine dormindo por 5 segundos...")
	time.Sleep(5 * time.Second)
	fmt.Println("Goroutine finalizada.")
}

func main() {
	go dormir()

	fmt.Println("Main dormindo por 3 segundos...")
	time.Sleep(3 * time.Second)
	fmt.Println("Main finalizada.")
}
