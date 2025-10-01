package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Mahima-Prajapati/BookStore/pkg/models"
	"github.com/Mahima-Prajapati/BookStore/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bookId"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		panic(err)
	}
	book, _ := models.GetBookById(bookId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	book = book.CreateBook()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bookId"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		panic(err)
	}
	book := models.DeleteBook(bookId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	id := vars["bookId"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		panic(err)
	}
	bookDetails, db := models.GetBookById(bookId)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
