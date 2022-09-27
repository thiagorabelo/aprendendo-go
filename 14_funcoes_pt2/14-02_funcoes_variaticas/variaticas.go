package main

import "fmt"

func soma(numeros ...int /* slice */) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func main() {
	fmt.Println("Funções Pt2: Funções Variáticas")

	total := soma(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
}
