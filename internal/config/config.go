package config

import (
	"context"
	"database/sql"
	"fmt"
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

	fmt.Println(db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := db.ExecContext(ctx, "create database")
	fmt.Println(res)
	if err != nil {
		log.Printf("Error %s when creating db\n", err)
	}

	db, _ := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbIp+"/"+dbName)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	if err := db.PingContext(ctx); err != nil {
		panic(err.Error())
	}
	return db
}
