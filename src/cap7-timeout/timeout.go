package main

import (
	"fmt"
	"time"
)

func executar(c chan<- bool) {
	time.Sleep(5 * time.Second)
	c <- true
}

func main() {
	c := make(chan bool, 1)
	go executar(c)

	fmt.Println("Esperando...")

	fim := false
	for !fim {
		select {
		case fim = <-c:
			fmt.Println("Fim!")
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
			fim = true
		}
	}
}
