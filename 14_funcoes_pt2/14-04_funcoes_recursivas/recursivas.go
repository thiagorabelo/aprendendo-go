package main

import "fmt"

func fibonacci(indice uint) uint {
	if indice <= 1 {
		return 1
	}
	return fibonacci(indice-1) + fibonacci(indice-1)
}

func main() {
	fmt.Println("Funções Pt2: Funções Recursivas")

	for indice := uint(0); indice <= 35; indice++ {
		fmt.Printf("%d) %v\n", indice, fibonacci(indice))
	}
}
