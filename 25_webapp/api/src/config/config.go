package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	URLBanco = ""
	Porta    = 0
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

	URLBanco = fmt.Sprintf(
		// user:password@tcp(localhost:5555)/dbname?charset=utf8&parseTime=True&loc=Local
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",

		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
