package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Book : book type
type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

//ToJson Convert to json
func (b Book) ToJson() []byte {

	byteArr, err := json.Marshal(b)

	if err != nil {
		log.Fatal(err)
	}

	return byteArr
}

func FromJson(data []byte) Book {

	var book Book
	err := json.Unmarshal(data, &book)

	if err != nil {
		log.Fatal(err)
	}

	return book
}

var books = map[string]Book{
	"123":  Book{Title: "hitchikers guide to the galaxy", Author: "Douglas Adams", ISBN: "123"},
	"0000": Book{Title: "cloud native go", Author: "reimer", ISBN: "0000"},
}

func BookHandleFunc(w http.ResponseWriter, req *http.Request) {

	isbn := req.URL.Path[len("/api/books/"):]

	fmt.Println(isbn)

	switch method := req.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			writeJSONSingleBook(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	case http.MethodPut:

		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		book := FromJson(body)
		exists := UpdateBook(isbn, book)

		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		executed := DeleteBook(isbn)
		if executed {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))

	}

}

func BooksHandleFunc(w http.ResponseWriter, req *http.Request) {

	switch method := req.Method; method {

	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		marshalledBooks, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJson(marshalledBooks)
		isbn, created := CreateBook(book)

		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
	}
}

func writeJSON(w http.ResponseWriter, books map[string]Book) {

	marshalledBooks, err := json.Marshal(books)

	if err != nil {
		panic(err)
	}

	w.Write(marshalledBooks)
}

func writeJSONSingleBook(w http.ResponseWriter, book Book) {
	marshaledBook, err := json.Marshal(book)

	if err != nil {
		panic(err)
	}

	w.Write(marshaledBook)
}

//UpdateBook : Update a book
func UpdateBook(isbn string, newBook Book) bool {

	book := books[isbn]

	if book.ISBN == "" {
		return false
	}

	books[isbn] = newBook
	return true

}

//DeleteBook : delete a book
func DeleteBook(isbn string) bool {

	book := books[isbn]

	if book.ISBN == "" {
		return false
	}

	delete(books, isbn)
	return true

}

//CreateBook : adds book to map
func CreateBook(book Book) (string, bool) {
	initLen := len(books)
	books[book.ISBN] = book

	if len(books) == initLen+1 {
		return book.ISBN, true
	} else {
		return "", false
	}

}

func GetBook(isbn string) (Book, bool) {

	book := books[isbn]
	if book.ISBN == "" {
		return book, false
	} else {
		return book, true
	}
}

func AllBooks() map[string]Book {
	return books
}
