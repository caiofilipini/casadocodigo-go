package main

import "fmt"

type ListaDeCompras []string

func (lista ListaDeCompras) Categorizar() ([]string, []string, []string) {
	vegetais, carnes, outros := []string{}, []string{}, []string{}

	for _, e := range lista {
		switch e {
		case "Alface", "Pepino":
			vegetais = append(vegetais, e)
		case "Atum", "Frango":
			carnes = append(carnes, e)
		default:
			outros = append(outros, e)
		}
	}

	return vegetais, carnes, outros
}

func main() {
	lista := make(ListaDeCompras, 6)
	lista[0] = "Alface"
	lista[1] = "Pepino"
	lista[2] = "Azeite"
	lista[3] = "Atum"
	lista[4] = "Frango"
	lista[5] = "Chocolate"

	vegetais, carnes, outros := lista.Categorizar()
	fmt.Println("Vegetais:", vegetais)
	fmt.Println("Carnes:", carnes)
	fmt.Println("Outros:", outros)
}
