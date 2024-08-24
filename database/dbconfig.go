package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"saurabhkanawade/jwt/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "omsairam"
	dbname   = "jwt"
)

var dns = fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable TimeZone = Asia/Kolkata", host, port, user, password, dbname)

var DB *gorm.DB

func DBConn() {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("The error found while making db conn %s", err)
	}
	DB = db

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("The error found while migrating the table %s", err)
	}

	fmt.Println("Successfully connected to the database.")

}
