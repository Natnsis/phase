package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load env file")
	}

	database_url := os.Getenv("DATABASE_URL")
	conn, err := gorm.Open(postgres.Open(database_url), &gorm.Config{})
	if err != nil {
		log.Fatal("unable to connect to the db")
	}

	Db = conn
}
