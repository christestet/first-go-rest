package jokesapi

import (
	"encoding/json"
	"net/http"
	"time"
)

const JokeAPIURL = "https://official-joke-api.appspot.com/jokes/programming/random"

var Client = &http.Client{Timeout: 2 * time.Second}

type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func FetchJoke() (Joke, bool) {
	resp, err := Client.Get(JokeAPIURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return Joke{}, false
	}
	defer resp.Body.Close()

	var jokes []Joke
	if err = json.NewDecoder(resp.Body).Decode(&jokes); err != nil {
		return Joke{}, false
	}
	if len(jokes) == 0 {
		return Joke{}, false
	}

	return jokes[0], true
}
