package jobs

import (
	"log"

	"github.com/imtomeddy/bbc-radio-spotify/internal/constants"
	"github.com/imtomeddy/bbc-radio-spotify/internal/request"
	"github.com/imtomeddy/bbc-radio-spotify/internal/structures"
)

//getLatestInfo gets the current broadcast and the current playign song
func getLatestInfo(station string) (*structures.Title, *structures.Broadcast, error) {
	broadcastRef, err := request.GetLatestBroadcast(station)

	if err != nil {
		return nil, nil, err
	}

	songRef, err := request.GetLatestSong(station)

	if err != nil {
		return nil, nil, err
	}

	return songRef, broadcastRef, nil
}

//updateStation updates information on one station
func updateStation(station string) {
	song, broadcast, err := getLatestInfo(station)

	if err != nil {
		log.Println(err)
		return
	}

	data := structures.DataPacket{
		Song:      *song,
		Broadcast: *broadcast,
	}

	log.Println(data)
}

//UpdateInfo updates all information
func UpdateInfo() {
	log.Printf("Updating information for %d stations", len(constants.Stations))

	for _, station := range constants.Stations {
		go updateStation(station)
	}
}
