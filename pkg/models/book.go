package models

import (
	"database/sql"
	"log"

	"github.com/DiegoAraujoJS/golang-books-api/pkg/config"
)

var db *sql.DB

type Book struct {
	ID     int
	Title  string
	Author string
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (b *Book) CreateBook() *Book {
	return b
}

func GetAllBooks() []Book {
	books, err := db.Query("SELECT * FROM book")
	defer books.Close()
	if err != nil {
		panic("error fetching books from DB")
	}
	var structBooks []Book
	for books.Next() {
		var book Book
		if err := books.Scan(&book.ID, &book.Title, &book.Author); err != nil {
			log.Fatal(err.Error())
		}
		structBooks = append(structBooks, book)
	}
	return structBooks
}
