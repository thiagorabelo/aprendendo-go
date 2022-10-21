package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"

	"github.com/gorilla/securecookie"
)

func GerarHashAndBlockKey() {
	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))

	fmt.Printf("HASH_KEY=%s\n", hashKey)
	fmt.Printf("BLOCK_KEY=%s\n", blockKey)
	fmt.Println()
}

func init() {
	// GerarHashAndBlockKey()

	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
}

func main() {
	router := router.Gerar()

	fmt.Printf("API Rodando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router))
}
