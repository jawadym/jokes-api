package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Joke struct {
	ID          string  `json:"id"`
	JokeType    string  `json:"jokeType"`
	JokeMessage string  `json:"jokeMessage"`
	Author      *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var jokes []Joke

func getJokes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}
func getJoke(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// get correct id
	for _, item := range jokes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Joke{})
}
func createJoke(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var joke Joke
	_ = json.NewDecoder(r.Body).Decode(&joke)
	joke.ID = strconv.Itoa(rand.Intn(10000000))
	jokes = append(jokes, joke)
	json.NewEncoder(w).Encode(joke)
}
func updateJoke(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range jokes {
		if item.ID == params["id"] {
			jokes = append(jokes[:i], jokes[i+1:]...)
			var joke Joke
			_ = json.NewDecoder(r.Body).Decode(&joke)
			joke.ID = params["id"]
			jokes = append(jokes, joke)
			json.NewEncoder(w).Encode(joke)
			return
		}
	}
}
func deleteJokes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range jokes {
		if item.ID == params["id"] {
			jokes = append(jokes[:i], jokes[i+1:]...)
			json.NewEncoder(w).Encode(item)
			break
		}
	}
	json.NewEncoder(w).Encode(&Joke{})

}

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
