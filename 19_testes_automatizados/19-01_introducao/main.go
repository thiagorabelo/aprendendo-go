package main

// go test ./...

/* Buscar e testar todas as funções em todos os pacotes */
// go test ./... -v

/* Mostra a cobertura dos testes */
// go test ./... --cover

/* Gera relatório de cobertura dos testes */
// go test ./... --coverprofile cobertura.txt

/* Analisa o arquivo de profile de cobertura dos teste e exibe um relatório legivel para humanos */
// go tool cover --func=cobertura.txt

/* Analisa o arquivo de profile de cobertura dos teste e exibe um relatório legivel para humanos e mais detalhado */
// go tool cover --html=cobertura.txt

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}
