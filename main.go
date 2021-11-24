package main

import (
	"firebird-golang/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := router.Router()

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
