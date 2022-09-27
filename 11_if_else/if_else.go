package main

import "fmt"

func retornaNumero() int32 {
	return 15
}

func main() {
	fmt.Println("If Else")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outro := retornaNumero(); outro > 0 {
		fmt.Println("Número é maior que zero", outro)
	} else if outro == 0 {
		fmt.Println("Número é igual a zero", outro)
	} else {
		fmt.Println("Número é menor que zero", outro)
	}
}
