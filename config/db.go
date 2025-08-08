package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// db.go-> setting up databse connection
var Db *gorm.DB

func ConnectDb() {
	error := godotenv.Load()
	if error != nil {
		log.Fatal("error loading .env file")
	}
	databaseURL := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("falied to connect: %v", err)
	}
	Db = database
}
