package models

import (
	"github.com/nhoc20170861/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

// Book represents the structure of a book in the bookstore
type Book struct {
	gorm.Model
	Name        string `json:"name"`        // Name of the book
	Author      string `json:"author"`      // Author of the book
	Publication string `json:"publication"` // Publication year or details
}

func init() {
	config.ConnectDatabase()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("id = ?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBookByID(id int64) Book {
	var book Book
	db.Where("id = ?", id).Delete(&book)
	return book
}
