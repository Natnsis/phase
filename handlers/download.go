package handlers

import (
	"log"

	"github.com/joho/godotenv"
)

func DownloadFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unabel to load")
	}
}
