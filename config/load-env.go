package config

import (
	"log"

	"github.com/joho/godotenv"
)

var EnvFile map[string]string

func init() {
	var err error
	EnvFile, err = godotenv.Read("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
