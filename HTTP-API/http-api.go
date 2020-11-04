package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, w *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, w *http.Request) {

}

func createBook(w http.ResponseWriter, w *http.Request) {

}

func updateBook(w http.ResponseWriter, w *http.Request) {

}

func deleteBook(w http.ResponseWriter, w *http.Request) {

}

func main() {
	r := mux.NewRouter()

	//Mock Data
	books = append(books, Book{ID: "1", Isbn: "1111", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "1122", Title: "Book Two", Author: &Author{Firstname: "Thomas", Lastname: "Steve"}})
	books = append(books, Book{ID: "3", Isbn: "1133", Title: "Book Three", Author: &Author{Firstname: "Jack", Lastname: "Andrew"}})
	books = append(books, Book{ID: "4", Isbn: "1144", Title: "Book Four", Author: &Author{Firstname: "Tim", Lastname: "C"}})

	r.HandleFunc("/api/books", getBooks).methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).methods("GET")
	r.HandleFunc("/api/books", createBook).methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
