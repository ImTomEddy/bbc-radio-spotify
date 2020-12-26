package spotifyclient

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/zmb3/spotify"
)

var auth spotify.Authenticator

//Client is the Spotify Client used to make spotify changes
var Client *spotify.Client = nil

//RequestAuthentication creates the HTTP listener for the request
func RequestAuthentication() {
	auth = spotify.NewAuthenticator(os.Getenv("SPOTIFY_REDIRECT_URI"), spotify.ScopePlaylistModifyPublic)
	auth.SetAuthInfo(os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET"))
	url := auth.AuthURL("")
	log.Println("Please authenticate with Spotify:")
	log.Println(url)

	http.HandleFunc("/", redirectHandler)
	go http.ListenAndServe(":8081", nil)

	for Client == nil {
		time.Sleep(5000000000)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	token, err := auth.Token("", r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusNotFound)
		return
	}

	c := auth.NewClient(token)
	Client = &c
}
