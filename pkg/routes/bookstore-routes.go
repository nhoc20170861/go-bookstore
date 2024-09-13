package routes

import (
	"github.com/gorilla/mux"
	"github.com/nhoc20170861/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	// Create a new book
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")

	// Get all books
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")

	// Get a book by its ID
	router.HandleFunc("/books/{id}", controllers.GetBookByID).Methods("GET")

	// Update a book by its ID
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")

	// Delete a book by its ID
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
