package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// u será uma cópia, pois se precisar do objeto original, tem que ser do tipo ponteiro
func (u usuario) maiorDeIdade() bool {
	return u.idade > 18
}

// u será uma cópia, pois se precisar do objeto original, tem que ser do tipo ponteiro
func (u usuario) salvar() {
	fmt.Printf("Salvando %v\n", u)
}

// Altera o próprio objeto e não uma cópia.
func (u *usuario) fazerAniversario() uint8 {
	u.idade++
	return u.idade
}

func main() {
	u1 := usuario{"Usuário 1", 20}
	fmt.Println(u1)
	fmt.Println("É maior de idade:", u1.maiorDeIdade())
	fmt.Println("Aniversário de", u1.fazerAniversario())
	u1.salvar()
}
