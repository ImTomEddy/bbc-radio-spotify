package main

import (
	"bbc-radio-spotify/internal/broadcasts"
	"bbc-radio-spotify/internal/constants"
	"bbc-radio-spotify/internal/songs"
	"log"
)

func main() {
	log.Println("Starting BBC Radio to Spotify Service")
	log.Printf("Supported Radio Stations: (%d) %s", len(constants.Stations), constants.Stations)

	for _, station := range constants.Stations {
		log.Println(station)
		broadcastRef, err := broadcasts.GetLatestBroadcast(station)

		if err != nil {
			log.Fatal(err)
			break
		}

		broadcast := *broadcastRef
		log.Println(broadcast)

		songRef, err := songs.GetLatestSong(station)

		if err != nil {
			log.Println(err)
			continue
		}

		song := *songRef
		log.Println(song)
	}
}
