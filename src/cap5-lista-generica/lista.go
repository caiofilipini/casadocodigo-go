package main

import "fmt"

type ListaGenerica []interface{}

func (lista *ListaGenerica) RemoverIndice(indice int) interface{} {
	l := *lista
	removido := l[indice]
	*lista = append(l[0:indice], l[indice+1:]...)
	return removido
}

func (lista *ListaGenerica) RemoverInicio() interface{} {
	return lista.RemoverIndice(0)
}

func (lista *ListaGenerica) RemoverFim() interface{} {
	return lista.RemoverIndice(len(*lista) - 1)
}

func main() {
	lista := ListaGenerica{
		1, "Café", 42, true, 23, "Bola", 3.14, false,
	}

	fmt.Printf("Lista original:\n%v\n\n", lista)

	fmt.Printf(
		"Removendo do início: %v, após remoção:\n%v\n\n",
		lista.RemoverInicio(), lista)
	fmt.Printf(
		"Removendo do fim: %v, após remoção:\n%v\n\n",
		lista.RemoverFim(), lista)
	fmt.Printf(
		"Removendo do índice 3: %v, após remoção:\n%v\n\n",
		lista.RemoverIndice(3), lista)
	fmt.Printf(
		"Removendo do índice 0: %v, após remoção:\n%v\n\n",
		lista.RemoverIndice(0), lista)
	fmt.Printf(
		"Removendo do último índice: %v, após remoção:\n%v\n\n",
		lista.RemoverIndice(len(lista)-1), lista)
}
