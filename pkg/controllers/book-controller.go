package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/nhoc20170861/go-bookstore/pkg/models" // Replace with your actual module path
	"github.com/nhoc20170861/go-bookstore/pkg/utils"
)

// CreateBook handles the creation of a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newBook := &models.Book{}

	// Parse the request body
	if err := utils.ParseBody(r, newBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newBook.CreateBook()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

// GetBooks retrieves all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newBooks := models.GetAllBooks()
	json.NewEncoder(w).Encode(newBooks)
}

// GetBookByID retrieves a book by its ID
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bookDetails, _ := models.GetBookByID(ID)

	if bookDetails == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(bookDetails)
}

// UpdateBook updates a book by its ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedBook := &models.Book{}
	// Parse the request body
	if err := utils.ParseBody(r, updatedBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookDetails, db := models.GetBookByID(ID)
	if bookDetails == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}

	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}

	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}
	db.Save(bookDetails)
	json.NewEncoder(w).Encode(bookDetails)
}

// DeleteBook deletes a book by its ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bookDetails := models.DeleteBookByID(ID)

	json.NewEncoder(w).Encode(bookDetails)
}
