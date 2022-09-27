package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numeroPtr *int) {
	*numeroPtr *= -1
}

func main() {
	numero := 20
	numeroOposto := inverterSinal(numero)
	fmt.Println(numero)
	fmt.Println(numeroOposto)

	outroNumero := 40
	fmt.Println(outroNumero)
	inverterSinalComPonteiro(&outroNumero)
	fmt.Println(outroNumero)
}
