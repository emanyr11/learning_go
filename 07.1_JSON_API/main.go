package main

import (

	// "strconv"

	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Mock data structure for testing that takes in an author struct
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Isbn   string  `json:"isbn"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

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
	// If the book doesnt exist point to empty book address
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {

	// create variable book of type book
	var book Book
	// create a blank identifier that is a new decoder that refernces the variable
	// Take the json data from post body an martial it into the structure
	_ = json.NewDecoder(r.Body).Decode(&book)
	// create a new book for books slice and assign to books
	books = append(books, book)
	// return the new book
	json.NewEncoder(w).Encode(book)

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	// Get the url parameters that are passed into the endpoint(id parameter)
	params := mux.Vars(r)
	// Create an index and iterate over the books slice
	for Index, Item := range books {
		// if the item in the params matches the Item.ID then assign the book to the index then remove it and increment
		// then break
		if Item.ID == params["id"] {
			// delete all from :inedex, too index+1: ...
			books = append(books[:Index], books[Index+1:]...)
			break
		}
	}
	// return the new books slice with deleted item
	json.NewEncoder(w).Encode(books)

}

func main() {
	//Initialize the router
	router := mux.NewRouter()
	// Mock Data
	books = append(books, Book{ID: "1", Title: "Harry Potter", Isbn: "12345", Author: &Author{Firstname: "JK", Lastname: "Rowling"}})
	books = append(books, Book{ID: "2", Title: "Lord of the rings", Isbn: "12322", Author: &Author{Firstname: "JRR", Lastname: "Tolkein"}})

	//Creat route handlers and establish endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	// Run the server
	log.Fatal(http.ListenAndServe(":8000", router))

}
