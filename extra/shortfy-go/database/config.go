package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDb() (*gorm.DB, error) {
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")
	port, err := strconv.ParseUint(os.Getenv("PORT"), 10, 64)
	sslmode := os.Getenv("SSLMODE")
	timezone := os.Getenv("TIMEZONE")

	if err != nil {
		return nil, err
	}

	urlDb := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		host,
		user,
		password,
		dbName,
		port,
		sslmode,
		timezone,
	)
	db, err := gorm.Open(postgres.Open(urlDb), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDb.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
