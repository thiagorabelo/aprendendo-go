package main

import (
	"fmt"
)

func part01() {
	const valor int = 100
	x := 10
	switch /* default é true */ {

	case x < 10: // true == (x < 10)
		fmt.Printf("Xís (%d) é menor que %d", x, valor)

	case x == 10: // true == (x == 10)
		fmt.Printf("Xís (%d) é igual a %d", x, valor)

	case x > 10: // true == (x > 10)
		fmt.Printf("Xís (%d) é maior que %d", x, valor)
	}
}

func part02() {
	const feriado = "quinta"

	switch feriado {
	case "terça":
		fmt.Println("Terça é feriado")

		// Se este bloco for true, avalia o resto
		// como verdadeiro também (comportamento default em C/C++, Java, etc)
		fallthrough
	case "segunda":
		fmt.Println("Segunda é feriado")
		break

	case "quarta":
		fmt.Println("Quarta é feiriado")

	case "quinta":
		fmt.Println("Quinta é feriado")
		fallthrough
	case "sexta":
		fmt.Println("Sexta é feriado")
		fallthrough

	case "sábado", "domingo": // compara com sábado ou domingo
		fmt.Println("Sábado e Domingo é folga")
	default:
		fmt.Println("Não temos feriado essa semana")
	}
}

func part03() {
	var x interface{}

	x = true

	switch x.(type) {
	case int:
		fmt.Println("Xís é um int")
	case bool:
		fmt.Println("Xís é um bool")
	case string:
		fmt.Println("Xís é um string")
	case float64:
		fmt.Println("Xís é um float64")
	default:
		fmt.Println("Sei não")
	}
}

func main() {
	part01()
	fmt.Printf("\n\n")
	part02()
	fmt.Printf("\n\n")
	part03()
}
