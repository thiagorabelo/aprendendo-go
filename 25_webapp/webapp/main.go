package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func init() {
	utils.CarregarTemplates()
}

func main() {
	router := router.Gerar()

	fmt.Println("Rodando WebApp")
	log.Fatal(http.ListenAndServe(":8000", router))
}
