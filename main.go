package main

import (
	"log"
	"net/http"
	"os"

	"github.com/imtomeddy/bbc-radio-spotify/internal/sounds"
	"github.com/imtomeddy/bbc-radio-spotify/internal/spotify"
)

func main() {
	err := spotify.Authenticate()

	if err != nil {
		panic(err)
	}

	p, err := spotify.GetOrCreatePlaylist("!!== TEST ==!!", "Test playlist please ignore")

	log.Println(p)

	log.Println("Starting BBC Radio to Spotify Service")
	log.Printf("Supported Radio Stations: (%d) %s", len(sounds.Stations), sounds.Stations)

	m := http.NewServeMux()
	m.HandleFunc("/", triggerHandler)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), m))
}

func triggerHandler(w http.ResponseWriter, r *http.Request) {
	// jobs.Update()
}
