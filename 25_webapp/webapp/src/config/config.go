package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// ApiURL contém o endereço para comunicação com a API
	APIURL string

	// Porta onde esta aplicação irá rodar
	Porta int

	// HashKey é utilizado para autenticar o cookie
	HashKey []byte

	// BlockKey é utilizado para criptografar o cookie
	BlockKey []byte
)

func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}

// API concatena url com o host da API retornando uma URI para comunicação com a API
func API(url string) string {
	return fmt.Sprintf("%s%s", APIURL, url)
}
