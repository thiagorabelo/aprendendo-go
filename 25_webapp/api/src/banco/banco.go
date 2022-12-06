package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq"
)

func Conectar() (*sql.DB, error) {
	// cfg := mysql.Config{
	// 	User:   os.Getenv("DB_USER"),
	// 	Passwd: os.Getenv("DB_PASS"),
	// 	Net:    "tcp",
	// 	Addr:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
	// 	DBName: os.Getenv("DB_NAME"),
	// 	AllowNativePasswords: true,
	// 	Params: map[string]string{
	// 		"charset":   "utf8",
	// 		"parseTime": "True",
	// 		"loc":       "Local",
	// 	},
	// }
	// db, err := sql.Open("mysql", cfg.FormatDSN())

	// db, err := sql.Open("postgres", config.URLBanco)
	db, err := sql.Open("mysql", config.URLBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
