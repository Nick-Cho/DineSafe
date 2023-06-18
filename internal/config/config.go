package config

import (
	"database/sql"
	"log"
	"time"
)

var (
	db  *sql.DB
	err error
)

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "admin"
	dbPass := "bigley2209"
	dbName := "allergy_db"
	dbIp := "tcp(allergy-project.cmdsxuexncin.us-east-1.rds.amazonaws.com:3306)"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbIp+"/"+dbName)

	if err != nil {
		log.Printf("Error %s when creating db\n", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db
}
