package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/Mahima-Prajapati/BookStore/pkg/models"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	// Get all the books
	books := models.GetAllBooks()
	res, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}

	// Set the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	// Get the book id from the request params, key is "bookId"
	vars := mux.Vars(r)
	id := vars["bookId"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		panic(err)
	}

	// Get the book with the specified id
	book, _ := models.GetBookById(bookId)
	res, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}

	// Set the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Get the book data from the request body
	book := &models.Book{}
	if body, err := io.ReadAll(r.Body); err == nil {
		if err = json.Unmarshal(body, &book); err != nil {
			return
		}
	}

	// Create a new book
	newBook := models.CreateBook(*book)
	res, err := json.Marshal(newBook)
	if err != nil {
		panic(err)
	}

	// Set the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Get the book id from the request params, key is "bookId"
	vars := mux.Vars(r)
	id := vars["bookId"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		panic(err)
	}

	// Delete the book with the specified id
	book := models.DeleteBook(bookId)
	res, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}

	// Set the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Get the book data from the request body
	updateBook := &models.Book{}
	if reqBody, err := io.ReadAll(r.Body); err == nil {
		if err = json.Unmarshal(reqBody, &updateBook); err != nil {
			return
		}
	}

	// Get the book id from the request params, key is "bookId"
	vars := mux.Vars(r)
	id := vars["bookId"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		panic(err)
	}

	// Get the older bookdetails from the DB for the specified "bookId"
	bookDetails, db := models.GetBookById(bookId)

	// Update the required fields
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	// Save changes in the DB
	db.Save(&bookDetails)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		panic(err)
	}

	// Set the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
