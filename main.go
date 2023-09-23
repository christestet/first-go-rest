package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/christestet/first-go-rest/fallbackjokes"
	"github.com/christestet/first-go-rest/jokesapi"
	"github.com/christestet/first-go-rest/renderer"
)

var lastJoke fallbackjokes.Joke

func getFallbackJoke() fallbackjokes.Joke {
	rand.Seed(time.Now().UnixNano())

	for {
		joke := fallbackjokes.JokesList[rand.Intn(len(fallbackjokes.JokesList))]
		if joke.Setup != lastJoke.Setup {
			lastJoke = joke
			log.Println("Last go fallback joke:", joke.Setup, joke.Punchline)
			return joke
		}
	}
}

func main() {
	const port = "9123"

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/fallback", func(w http.ResponseWriter, r *http.Request) {
		joke := getFallbackJoke()
		renderer.RenderJoke(w, joke)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if joke, ok := jokesapi.FetchJoke(); ok {
			renderer.RenderJoke(w, joke)
		} else {
			log.Println("Using fallback joke due to API error or timeout.")
			fallback := getFallbackJoke()
			renderer.RenderJoke(w, map[string]interface{}{
				"Setup":     fallback.Setup,
				"Punchline": fallback.Punchline,
				"Image":     fallback.Image,
			})
		}
	})

	address := ":" + port
	log.Printf("🌐 Server started. Go get your jokes at: http://localhost:%s", port)

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
