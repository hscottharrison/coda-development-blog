package main

import (
	"log"
	"os"
)

func main() {
	store, err := NewPostgresStore()

	if err != nil {
		log.Fatal("ERROR in creating store", err)
	}

	port := os.Getenv("PORT")
	apiServer := NewAPIServer(port, store)
	apiServer.Run()
}
