package main

import "fmt"

func alunoEstaAprovado(n1, n2 float32) bool {
	// DEFER = ADIAR
	// Achamada é adiada até o último momento possível. Ex: Antes do retorno.
	// É ideal para fechar conexão com o banco de dados, por exemplo.
	defer fmt.Println("Média calculada. O resultado será calculado.")

	fmt.Println("Entrando na função para verificar se o Aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 7 {
		return true
	}

	return false
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 8))
}
