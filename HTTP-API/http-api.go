package main

import (
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

func getBooks(w http.ResponseWriter, w *http.Request) {

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

	r.HandleFunc("/api/books", getBooks).methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).methods("GET")
	r.HandleFunc("/api/books", createBook).methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
