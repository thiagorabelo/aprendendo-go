package main

import "fmt"

func main() {
	fmt.Println("Funções Pt2: Funções Anônimas")

	func(texto string) {
		fmt.Println(texto)
	}("Olá, mundo")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido: %s", texto)
	}("Olá, GoLang")
	fmt.Println(retorno)
}
