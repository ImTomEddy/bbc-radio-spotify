package main

import (
	"fmt"
	"log"
	"time"

	"github.com/imtomeddy/bbc-radio-spotify/internal/constants"
	"github.com/imtomeddy/bbc-radio-spotify/internal/jobs"
	"github.com/imtomeddy/bbc-radio-spotify/internal/spotifyclient"
)

func main() {
	spotifyclient.RequestAuthentication()

	log.Println("Starting BBC Radio to Spotify Service")
	log.Printf("Supported Radio Stations: (%d) %s", len(constants.Stations), constants.Stations)

	for _, station := range constants.Stations {
		playlistName := fmt.Sprintf("%s | All Tracks Daily | %s", constants.StationNames[station], time.Now().Format("02/01/2006"))
		playlistDesc := "Playlist generated using bbc-radio-spotify (https://github.com/imtomeddy/bbc-radio-spotify)."

		spotifyclient.GetPlaylist(station+"_daily", playlistName, playlistDesc)
	}

	jobs.SetupJobs()

	select {}
}
