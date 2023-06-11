package config

import (
	"database/sql"
	"log"
	"time"
)

var (
	db *sql.DB
)

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "$Bigley2209"
	dbName := "allergy_db"
	dbIp := "tcp(127.0.0.1:3306)"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbIp+"/"+dbName)
	defer db.Close()

	if err != nil {
		log.Printf("Error %s when creating db\n", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db
}
