package jobs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/imtomeddy/bbc-radio-spotify/internal/constants"
	"github.com/imtomeddy/bbc-radio-spotify/internal/request"
	"github.com/imtomeddy/bbc-radio-spotify/internal/spotifyclient"
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

	playlistName := fmt.Sprintf("%s %s | %s | %s", data.Broadcast.Title.Primary, data.Broadcast.Title.Secondary, constants.StationNames[station], time.Now().Format("02/01/2006"))
	playlistDesc := "Playlist generated using bbc-radio-spotify (https://github.com/imtomeddy/bbc-radio-spotify). Boradcast ID: " + data.Broadcast.ID

	playlistRef, err := spotifyclient.GetPlaylist(station, playlistName, playlistDesc)
	playlist := *playlistRef

	if err != nil {
		log.Println(err)
		return
	}

	songName := data.Song.Primary + " " + data.Song.Secondary

	spotifyclient.AddSongToPlaylist(playlist.ID, songName, station)
}

//UpdateInfo updates all information
func UpdateInfo() {
	for _, station := range constants.Stations {
		updateStation(station)

		if os.Getenv("DEV") == "true" {
			break
		}
	}
}
