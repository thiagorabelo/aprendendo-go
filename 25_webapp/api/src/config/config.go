package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// URLBanco é a URL de conexão com o banco de dados
	URLBanco = ""

	// Porta onde a API irá escutar
	Porta = 0

	// SecretKey é a chave usada para assinar o token
	SecretKey []byte
)

func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Porta = 9000
	}

	// MySQL
	URLBanco = fmt.Sprintf(
		// user:password@tcp(localhost:5555)/dbname?charset=utf8&parseTime=True&loc=Local
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",

		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// // PostgreSQL
	// URLBanco = fmt.Sprintf(
	// 	// postgres://user:password@localhost:port/dbname?sslmode=disable
	// 	"postgres://%s:%s@%s:%s/%s?sslmode=disable",

	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASS"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_NAME"),
	// )

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
