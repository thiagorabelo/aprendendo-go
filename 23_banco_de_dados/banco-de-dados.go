package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Permite importar o pacote, mas não usá-lo. Dentro do pacote tem uma função
	// `init()` que irá registrar o Driver do MySQL.
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("Carregando variáveis de ambiente")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Não foi possível carregar variáveis de ambiente: %s", err)
	}
}

func main() {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	urlConexao := fmt.Sprintf("%s:%s@/devbook?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass)
	db, err := sql.Open("mysql", urlConexao)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// fmt.Println(db)

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão aberta")

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
