package config

import (
	"context"
	"database/sql"
	"time"
)

var (
	ctx context.Context
	db  *sql.DB
)

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "$Bigley2209"
	dbName := "allergy_db"
	dbIp := "tcp(localhost:3306)"

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	db, _ := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbIp+"/"+dbName)
	if err := db.PingContext(ctx); err != nil {
		panic(err.Error())
	}
	return db
}
