package main

import (
	"log"

	server "github.com/Two-Skoopz-Development/biblestudy-api/internal"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server.StartServer()
}
