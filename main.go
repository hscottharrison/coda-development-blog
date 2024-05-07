package main

import "log"

const portNumber = ":8080"

func main() {
	store, err := NewPostgresStore()

	if err != nil {
		log.Fatal("ERROR in creating store", err)
	}
	apiServer := NewAPIServer(portNumber, store)
	apiServer.Run()
}
