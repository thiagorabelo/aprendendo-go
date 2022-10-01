package banco

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // Registra o driver do MySQL
)

func Conectar() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	url := fmt.Sprintf("%s:%s@/devbook?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass)

	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
