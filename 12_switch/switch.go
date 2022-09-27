package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}
}

func feriado(numero int) []string {
	var feriados []string

	switch numero {
	case 3:
		feriados = append(feriados, "Terça-Feira")
		fallthrough // Se cair aqui, também executa o próximo bloco (padrão em c/c++, java, etc).
	case 2:
		feriados = append([]string{"Segunda-Feira"}, feriados...) // Prepend: append(novo_slice_com_elemento_desejado, slice_antigo_decomposto...).
	case 4:
		feriados = append(feriados, "Quarta-Feira")
	case 5:
		feriados = append(feriados, "Quinta-Feira")
		fallthrough // Se cair aqui, também executa o próximo bloco (padrão em c/c++, java, etc).
	case 6:
		feriados = append(feriados, "Sexta-Feira")
	case 7, 1: // Testa os dois.
		feriados = append(feriados, "Sábado e Domingo é folga")
	default:
		feriados = append(feriados, "Número Inválido")
	}

	return feriados
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(3)
	fmt.Println(dia)

	dia = diaDaSemana2(2)
	fmt.Println(dia)

	feriados := feriado(3)
	fmt.Println(feriados)
}
