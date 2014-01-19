package main

import (
	"fmt"
	"time"
)

type Operacao interface {
	Calcular() int
}

type Soma struct {
	operando1, operando2 int
}

func (s Soma) Calcular() int {
	return s.operando1 + s.operando2
}

func (s Soma) String() string {
	return fmt.Sprintf("%d + %d", s.operando1, s.operando2)
}

type Subtracao struct {
	operando1, operando2 int
}

func (s Subtracao) Calcular() int {
	return s.operando1 - s.operando2
}

func (s Subtracao) String() string {
	return fmt.Sprintf("%d - %d", s.operando1, s.operando2)
}

type Idade struct {
	anoNascimento int
}

func (i Idade) Calcular() int {
	return time.Now().Year() - i.anoNascimento
}

func (i Idade) String() string {
	return fmt.Sprintf("Idade desde %d", i.anoNascimento)
}

func acumular(operacoes []Operacao) int {
	acumulador := 0

	for _, op := range operacoes {
		valor := op.Calcular()
		fmt.Printf("%v = %d\n", op, valor)
		acumulador += valor
	}

	return acumulador
}

func main() {
	operacoes := make([]Operacao, 4)
	operacoes[0] = Soma{10, 20}
	operacoes[1] = Subtracao{30, 15}
	operacoes[2] = Subtracao{10, 50}
	operacoes[3] = Soma{5, 2}

	fmt.Println("Valor acumulado =", acumular(operacoes))

	idades := make([]Operacao, 3)
	idades[0] = Idade{1969}
	idades[1] = Idade{1977}
	idades[2] = Idade{1999}

	fmt.Println("Idades acumuladas =", acumular(idades))
}
