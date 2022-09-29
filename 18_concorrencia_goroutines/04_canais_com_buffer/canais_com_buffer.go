package main

import (
	"fmt"
)

func main() {
	/*
		Quando usado sem a capacidade ou enviando dados superior a capacidade na mesma função,
		o canal causa deadlock, pois não há "ninguém" aguardando o recebimento dos dados.
	*/
	canal := make(chan string, 2) // Sem a capacidade dará deadlock se usado na mesma função.
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"
	// canal <- "Terceira mensagem" // Causa deadlock, pois abaixo só faz duas leituras.

	mensagem := <-canal
	fmt.Println(mensagem)

	mensagem = <-canal
	fmt.Println(mensagem)
}
