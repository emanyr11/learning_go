package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Creating book struct (model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Creating author struct to be called in book struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Initialize book variable as a slice book structure
var books []Book

// all route handler functions takes in response and request similar to app.get(``,(req,res))
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get the url parameters that are passed into the endpoint(id parameter)
	params := mux.Vars(r)
	// _ is a blank identiier used because only the second item in the range is needed
	// Iterate of the range of books
	for _, Item := range books {
		// if the item matches the params id encode the item and return
		if Item.ID == params["id"] {
			json.NewEncoder(w).Encode(Item)
			return
		}
	}
	// If the book doesnt exist return an empty object
	json.NewEncoder(w).Encode(&Book{})

}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// create variable book of type book
	var book Book
	// create a blank identifier that is a new decoder that refernces the variable
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	// create a new book
	books = append(books, book)
	// encode the new book
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get the url parameters that are passed into the endpoint(id parameter)
	params := mux.Vars(r)
	// Create an index and iterate over the books slice
	for index, item := range books {
		// if the item in the params matches the Item.ID then assign the book to the index then remove it and increment
		// then break
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	// encode the new books slice
	json.NewEncoder(w).Encode(books)
}

func main() {
	//Initialize the router
	router := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "298733", Title: "Harry Potter", Author: &Author{Firstname: "J.K", Lastname: "Rowling"}})
	books = append(books, Book{ID: "2", Isbn: "346859", Title: "Lord of the Rings", Author: &Author{Firstname: "J. R. R.", Lastname: "Tolkien"}})

	//Creat route handlers and establish endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Run the server
	log.Fatal(http.ListenAndServe(":8000", router))

}
