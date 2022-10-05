package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := router.Gerar()

	fmt.Println("API Rodando")
	log.Fatal(http.ListenAndServe(":5000", router))
}
