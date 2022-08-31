package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/DiegoAraujoJS/golang-books-api/pkg/utils"
	"github.com/DiegoAraujoJS/golang-books-api/pkg/models"
)

var NewBook models.Book

// type Book struct {
//	gorm.model
//	Name string `gorm:""json:"name"`
//	Author string `json:"author"`
//	Publication string `json:"publication"`
//}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}


func GetBookById(w http.ResponseWriter, r *http.Request){
	var book Book
	allBooks = models.GetAllBooks()
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv(bookId, 0, 0)
	if err != nil {
		fmt.Println("error parsing")
	}
	book, _ = models.GetModelById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	var newBook Book
	vars := mux.Vars(r)
	newBook
	createdBook := models.CreateBook(&NewBook)
}
