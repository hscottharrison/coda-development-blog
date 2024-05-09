package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	store, err := NewPostgresStore()

	if err != nil {
		log.Fatal("ERROR in creating store", err)
	}

	godotenv.Load()
	port := os.Getenv("PORT")
	apiServer := NewAPIServer(port, store)
	apiServer.Run()
}
