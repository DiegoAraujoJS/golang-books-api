package main

import (
	"net/http"
    "fmt"
    "errors"
    "os"
    "encoding/json"
	"github.com/DiegoAraujoJS/golang-books-api/pkg/models"
)

func main() {
    http.HandleFunc("/", GetBook)
    err := http.ListenAndServe(":9010", nil)
    if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
    }
}

func GetBook(w http.ResponseWriter, r *http.Request) {
    fmt.Println("get book")
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
