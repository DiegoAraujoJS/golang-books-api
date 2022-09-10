package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DiegoAraujoJS/golang-books-api/pkg/models"
	"github.com/DiegoAraujoJS/golang-books-api/pkg/utils"
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
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//func GetBookById(w http.ResponseWriter, r *http.Request) {
//	var book Book
//	allBooks = models.GetAllBooks()
//	vars := mux.Vars(r)
//	bookId := vars["bookId"]
//	ID, err := strconv.ParseInt(bookId, 0, 0)
//	if err != nil {
//		fmt.Println("error parsing")
//	}
//	book, _ = models.GetModelById(ID)
//	res, _ := json.Marshal(bookDetails)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write(res)
//}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
