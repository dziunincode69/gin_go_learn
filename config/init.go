package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitializeENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
