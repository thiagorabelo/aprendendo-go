package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint8
	altura float32
}

type estudante struct {
	pessoa    // desempacotando o pessoa detro de estudante
	matricula string
	curso     string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", 21, 1.73}
	fmt.Println(p1)

	e1 := estudante{p1, "222CICOM1", "Bacharelado Em Ciência da Computação"}
	fmt.Println(e1)
	fmt.Println(e1.pessoa.nome)
	fmt.Println(e1.nome)
}
