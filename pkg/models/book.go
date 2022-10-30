package models

import (
	"database/sql"
	"fmt"
	"github.com/DiegoAraujoJS/golang-books-api/pkg/config"
	"log"
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
	fmt.Println("get all books")
	books, err := db.Query("SELECT * FROM book")
	if err != nil {
		return []Book{}
	}
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
