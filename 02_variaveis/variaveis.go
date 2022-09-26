package main

import "fmt"

func main() {
	var variavel1 string = "variável 1" // tipo explícito
	fmt.Println(variavel1)

	variavel2 := "variável 2" // inferência de tipo
	fmt.Println(variavel2)

	var (
		variavel3 string = "var 3"
		variavel4 string = "var 4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Var 5", "Var 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)
}
