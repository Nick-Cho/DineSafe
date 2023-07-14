package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	db  *sql.DB
	err error
)

func Connect() *sql.DB {

	err := godotenv.Load(".env")

	dbDriver := "mysql"
	dbUser := "admin"
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := "allergy_db"
	dbIp := os.Getenv("DB_IP")

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbIp+"/"+dbName)

	if err != nil {
		log.Printf("Error %s when creating db\n", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db
}
