package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var jokes []Joke

func main() {
	// constants
	const PORT int = 8000
	r := mux.NewRouter()
	jokes = append(jokes, Joke{ID: "1", JokeType: "Misc", JokeMessage: "Knock Knock", Author: &Author{FirstName: "Muhammad Habib", LastName: "Jawady"}})
	jokes = append(jokes, Joke{ID: "2", JokeType: "Misc", JokeMessage: "VB", Author: &Author{FirstName: "Muhammad Habib", LastName: "Jawady"}})
	r.HandleFunc("/api/jokes", getJokes).Methods("GET")
	r.HandleFunc("/api/jokes/{id}", getJoke).Methods("GET")
	r.HandleFunc("/api/jokes", createJoke).Methods("POST")
	r.HandleFunc("/api/jokes/{id}", updateJoke).Methods("PUT")
	r.HandleFunc("/api/jokes/{id}", deleteJokes).Methods("DELETE")

	// Listen, serve and log
	fmt.Printf("Listening on port %d ...\n", PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), r))
}
