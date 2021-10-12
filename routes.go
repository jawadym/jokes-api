package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
