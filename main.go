package main

import (
	"log"

	"github.com/imtomeddy/bbc-radio-spotify/internal/constants"
	"github.com/imtomeddy/bbc-radio-spotify/internal/jobs"
	"github.com/imtomeddy/bbc-radio-spotify/internal/spotifyclient"
)

func main() {
	spotifyclient.RequestAuthentication()

	log.Println("Starting BBC Radio to Spotify Service")
	log.Printf("Supported Radio Stations: (%d) %s", len(constants.Stations), constants.Stations)

	jobs.SetupJobs()

	select {}
}
