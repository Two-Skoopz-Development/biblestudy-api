package main

import (
	"log"

	apiServer "github.com/Two-Skoopz-Development/biblestudy-api/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiServer.StartServer()
}
