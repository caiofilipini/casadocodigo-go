package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func executar(controle *sync.WaitGroup) {
	defer controle.Done()

	duracao := time.Duration(1+rand.Intn(5)) * time.Second
	fmt.Printf("Dormindo por %s...\n", duracao)
	time.Sleep(duracao)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var controle sync.WaitGroup
	inicio := time.Now()

	for i := 0; i < 5; i++ {
		controle.Add(1)
		go executar(&controle)
	}

	controle.Wait()

	fmt.Printf("Finalizado em %s.\n", time.Since(inicio))
}
