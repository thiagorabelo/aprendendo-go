package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint16
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	fmt.Println(u)

	u.nome = "Thiago"
	u.idade = 38
	fmt.Println(u)

	u2 := usuario{"Thiago", 38, endereco{"Rua sem saída", 11}}
	fmt.Println(u2)

	u3 := usuario{nome: "Thiago"}
	fmt.Println(u3)
}
