package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/christestet/first-go-rest/fallbackjokes"
)

const jokeAPI = "https://official-joke-api.appspot.com/jokes/programming/random"

var client = &http.Client{Timeout: 2 * time.Second}

type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

var lastJoke fallbackjokes.Joke

func getFallbackJoke() fallbackjokes.Joke {
	rand.Seed(time.Now().UnixNano())
	var joke fallbackjokes.Joke

	for {
		joke = fallbackjokes.JokesList[rand.Intn(len(fallbackjokes.JokesList))]
		if joke.Setup != lastJoke.Setup {
			break
		}
	}

	lastJoke = joke
	log.Println("Last go fallback joke:", joke.Setup, joke.Punchline)
	return joke
}

func fetchAPIJoke() (Joke, bool) {
	resp, err := client.Get(jokeAPI)
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

func renderJoke(w http.ResponseWriter, joke interface{}) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		log.Println("Failed to load template:", err)
		return
	}

	tmpl.Execute(w, joke)
}

func main() {
	const port = "9123"

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/fallback", func(w http.ResponseWriter, r *http.Request) {
		joke := getFallbackJoke()
		renderJoke(w, joke)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if joke, ok := fetchAPIJoke(); ok {
			renderJoke(w, joke)
		} else {
			log.Println("Using fallback joke due to API error or timeout.")
			joke := getFallbackJoke()
			renderJoke(w, map[string]interface{}{
				"Setup":     joke.Setup,
				"Punchline": joke.Punchline,
				"Image":     joke.Image,
			})
		}
	})

	address := ":" + port
	log.Printf("🌐 Server started. Get your jokes at: http://localhost:%s", port)

	// Handle SIGINT (Ctrl+C)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	go func() {
		<-sigCh
		log.Println("Don't you want to hear any more jokes? 🤔 Ok, press 'Ctrl+C' again to really exit.")
		<-sigCh
		os.Exit(0)
	}()

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
