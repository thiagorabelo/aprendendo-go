package main

import (
	"api/src/config"
	"api/src/router"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

func GerarSecretKey() {
	const privateKeySize = 64
	chave := make([]byte, privateKeySize)

	if _, err := rand.Read(chave); err != nil {
		log.Fatal(err)
	}

	secretKey := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(secretKey)
}

func init() {
	// GerarSecretKey()

	config.Carregar()
}

func main() {
	router := router.Gerar()

	fmt.Printf("API Rodando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router))
}
