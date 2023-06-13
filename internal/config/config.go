package config

import (
	"log"

	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Connect() *gorm.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "$Bigley2209"
	dbName := "allergy_db"
	dbIp := "tcp(allergy-project.cmdsxuexncin.us-east-1.rds.amazonaws.com:3306)"

	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@"+dbIp+"/"+dbName)

	if err != nil {
		log.Printf("Error %s when creating db\n", err)
	} else {
		log.Println("Connection Established")
	}

	defer db.Close()
	return db
}
