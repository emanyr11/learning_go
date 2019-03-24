package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Create struct of mock data

type Person struct {
	Socialsecurity string  `json:"socialsecurity"`
	Firstname      string  `json:"firstname"`
	Lastname       string  `json:"lastname"`
	Email          string  `json:"email"`
	// pointer to the address struct using type address
	Adress         *Adress `json:"adress"`
}

// Adress struct called by person struct to initialse the adress
type Adress struct {
	City  string `json:"city"`
	State string `json:"state"`
}

// A slice of type Person called people 
var people []Person

// function call to retrieve all people
func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

// Fucntion call te retrieve a person specified by the socialsecurity number
func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range people {
		if item.Socialsecurity == params["socialsecurity"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(people)
}

// function call to create a person by martialing data and assigning to a varaible person 
func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	people = append(people, person)
	json.NewEncoder(w).Encode(person)

}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range people {
		if item.Socialsecurity == params["socialsecurity"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	params := mux.Vars(r)
	for index, item := range people {
		if item.Socialsecurity == params["socialsecurity"] {
			people = append(people[:index], people[index+1:]...)
			var person Person
			person.Socialsecurity = params["socialsecurity"]
			_ = json.NewDecoder(r.Body).Decode(&person)
			people = append(people, person)
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	json.NewEncoder(w).Encode(people)

}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{Socialsecurity: "1", Firstname: "John", Lastname: "Doe", Email: "john@gmail.com", Adress: &Adress{City: "Perth", State: "WA"}})
	people = append(people, Person{Socialsecurity: "2", Firstname: "Jane", Lastname: "Doe", Email: "Jane@gmail.com", Adress: &Adress{City: "Melboune", State: "Victoria"}})
	people = append(people, Person{Socialsecurity: "3", Firstname: "John", Lastname: "Doe", Email: "askdh@gmail.com", Adress: &Adress{City: "Sydney", State: "NSW"}})
	router.HandleFunc("/api/people", getPeople).Methods("GET")
	router.HandleFunc("/api/person/{socialsecurity}", getPerson).Methods("GET")
	router.HandleFunc("/api/person", createPerson).Methods("POST")
	router.HandleFunc("/api/person/{socialsecurity}", deletePerson).Methods("DELETE")
	router.HandleFunc("/api/person/{socialsecurity}", updatePerson).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}
