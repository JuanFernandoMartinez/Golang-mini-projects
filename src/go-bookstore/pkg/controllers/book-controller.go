package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"example.com/pkg/models"
	util "example.com/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Print("error while parsing\n")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	util.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BookId := vars["bookId"]
	ID, err := strconv.ParseInt(BookId, 0, 0)
	if err != nil {
		fmt.Print("Error converting ID")
	}

	DeleteBook := models.DeleteBook(ID)

	res, _ := json.Marshal(DeleteBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	NewBook := models.Book{}

	util.ParseBody(r, NewBook)

	vars := mux.Vars(r)

	id := vars["bookId"]
	BookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Print("Error converting Id")
	}

	BookToUpdate, db := models.GetBookById(BookId)

	if NewBook.Author != "" {
		BookToUpdate.Author = NewBook.Author
	}

	if NewBook.Name != "" {
		BookToUpdate.Name = NewBook.Name
	}

	if NewBook.Publication != "" {
		BookToUpdate.Publication = NewBook.Publication
	}

	db.Save(BookToUpdate)

}
