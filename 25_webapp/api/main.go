package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func init() {
	config.Carregar()
}

func main() {
	router := router.Gerar()

	fmt.Printf("API Rodando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router))
}
