package main

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
