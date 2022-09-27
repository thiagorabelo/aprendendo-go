package main

import (
	"fmt"
	"time"
)

func while() {
	i := 0
	for i < 10 { // Equivalente ao while
		i++
		fmt.Printf("Incrementando i=%d\n", i)
		time.Sleep(time.Second / 5)
	}
	fmt.Println("O valor final de i é", i)
}

func forLoop() {
	for i := 0; i < 10; i += 2 {
		fmt.Printf("Incrementando i=%d\n", i)
		time.Sleep(time.Second / 5)
	}
}

func forRange() {
	nomes := [3]string{"João", "Davi", "Lucas"}
	for indice, nome := range nomes {
		fmt.Printf("Índice: %d, Valor: %v\n", indice, nome)
	}
	fmt.Println()
}

func forRangeString() {
	palavra := "Órgão"
	for indice, codePoint := range palavra {
		fmt.Printf("Índice=%d, CodePoint=%d, Caracter='%c'\n", indice, codePoint, codePoint)
	}
}

func forRangeMap() {
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Printf("Chave=\"%v\", Valor=\"%v\"\n", chave, valor)
	}
}

func forInfinito() {
	i := 0
	for {
		i++
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second / 5)
		if i > 5 {
			fmt.Println("Parando o loop")
			break // Parando o loop
		}
	}
}

func main() {
	fmt.Println("Loops")

	fmt.Println("# Equivalente ao While")
	while()

	fmt.Println("\n# For Loop")
	forLoop()

	fmt.Println("\n# For Range")
	forRange()

	fmt.Println("\n# For Range String")
	forRangeString()

	fmt.Println("\n# For Range Map")
	forRangeMap()

	fmt.Println("\n# For Infinito")
	forInfinito()
}
