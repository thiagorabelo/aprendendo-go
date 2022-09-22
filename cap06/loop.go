package main

import (
	"fmt"
)

func main() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	fmt.Println("\n---\n ")

	i := 1
	for i <= 3 { // Equivalente ao while
		fmt.Println(i)
		i += 1
	}

	fmt.Println("\n---\n ")

	for {
		fmt.Println("Loop sem fim")
		break // parando o loop infinito
	}

	fmt.Println("\n---\n ")

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue // pulando o resto do bloco e voltando para o início do laço
		}
		fmt.Println(n)
	}
}
