package main

import "net/http"

// "strconv"
// json Encoding
// http package
// mux package
// log package

// Mock data structure for testing that takes in an author struct

// creat a global variable books slice of type Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	// print out the books using encoder

}

func getBook(w http.ResponseWriter, r *http.Request) {

	// Get the url parameters that are passed into the endpoint(id parameter)

	// _ is a blank identiier used because only the second item in the range is needed
	// Iterate of the range of books

	// if the item matches the params id encode the item and return

	// If the book doesnt exist point to empty book address

}

func createBook(w http.ResponseWriter, r *http.Request) {

	// create variable book of type book

	// create a blank identifier that is a new decoder that refernces the variable
	// Take the json data from post body an martial it into the structure

	// create a new book for books slice and assign to books

	// return the new book

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	// Get the url parameters that are passed into the endpoint(id parameter)

	// Create an index and iterate over the books slice

	// if the item in the params matches the Item.ID then assign the book to the index then remove it and increment then break

	// delete all from :inedex, too index+1: ...

	// return the new books slice with deleted item

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	// Delete the book then create the book
}

func main() {
	//Initialize the router

	// Mock Data

	//Creat route handlers and establish endpoints

	// Run the server

}
