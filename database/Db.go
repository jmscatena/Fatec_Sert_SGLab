package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CONNECTION struct {
	Conn gorm.DB
}

func Init() (*gorm.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {

		log.Fatalf("Error Loading Configuration File")
	}

	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbase := os.Getenv("DB")
	dbServer := os.Getenv("DBSERVER")
	dbPort := os.Getenv("DBPORT")
	dbURL := "postgres://" + dbUser + ":" + dbPass + "@" + dbServer + ":" + dbPort + "/" + dbase

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	return db, err
}
