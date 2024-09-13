package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nhoc20170861/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Println("Server started at http://localhost:4000")
	log.Println("Listening on port 4000...")

	// Start the server
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Printf("Could not start server: %v", err)
	}
}
