package spotifyclient

import (
	"fmt"
	"log"

	"github.com/zmb3/spotify"
)

var mostRecentSongs map[string]string = make(map[string]string)

//AddSongToPlaylist tries to add the song to a playlist, as long as it's not a duplicate
func AddSongToPlaylist(playlist spotify.ID, songName string, station string) error {

	result, err := Client.Search(songName, spotify.SearchTypeTrack)

	if err != nil {
		return err
	}

	if len((*result.Tracks).Tracks) == 0 {
		return fmt.Errorf("Unable to find song '%s'", songName)
	}

	song := (*result.Tracks).Tracks[0]

	if mostRecentSongs[station] == song.ID.String() {
		return nil
	}

	mostRecentSongs[station] = song.ID.String()
	Client.AddTracksToPlaylist(playlist, song.ID)

	log.Printf("Added song %s to playlist %s for station %s", song.ID, playlist, station)
	return nil
}
