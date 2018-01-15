package main

import (
	"encoding/json"
	"log"
	"net/http"
	// "math/rand"
	// "strconv"
	// "fmt"
	"github.com/gorilla/mux"
)

// Init books
var books []Book

// Book Strut
type Book struct {
	ID     string  `json:"id`
	Isbn   string  `json:"isbn`
	Title  string  `json:"title"`
	Author *Author `json:"author`
}
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Retrieved params

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBooks(w http.ResponseWriter, r *http.Request) {

}

func deleteBooks(w http.ResponseWriter, r *http.Request) {

}
func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "449123", Title: "Book one",
		Author: &Author{Firstname: "Jean", Lastname: "Dupont"}})
	books = append(books, Book{ID: "1", Isbn: "449123", Title: "Book two",
		Author: &Author{Firstname: "Jean", Lastname: "Dupont"}})
	// Route Endpoints

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
