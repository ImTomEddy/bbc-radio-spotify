package spotifyclient

import (
	"log"

	"github.com/imtomeddy/bbc-radio-spotify/internal/constants"
	"github.com/zmb3/spotify"
)

var playlists map[string]string = make(map[string]string)
var lastSong map[string]string = make(map[string]string)

//GetPlaylist will get or create a Spotify playlist
func GetPlaylist(station string, playlistName string, playlistDesc string) (*spotify.SimplePlaylist, error) {
	if playlists[station] == "" {
		newPlaylist, err := Client.CreatePlaylistForUser(constants.UserID, playlistName, playlistDesc, true)

		if err != nil {
			return nil, err
		}

		playlists[station] = newPlaylist.ID.String()
		log.Printf("Created playlist %s for station %s", newPlaylist.ID, station)
		return &newPlaylist.SimplePlaylist, nil
	}

	result, err := Client.GetPlaylist(spotify.ID(playlists[station]))

	if err != nil {
		return nil, err
	}

	if result == nil {
		newPlaylist, err := Client.CreatePlaylistForUser(constants.UserID, playlistName, playlistDesc, true)

		if err != nil {
			return nil, err
		}

		playlists[station] = newPlaylist.ID.String()
		log.Printf("Created playlist %s for station %s", newPlaylist.ID, station)
		return &newPlaylist.SimplePlaylist, nil
	}

	if result.Name != playlistName {
		newPlaylist, err := Client.CreatePlaylistForUser(constants.UserID, playlistName, playlistDesc, true)

		if err != nil {
			return nil, err
		}

		playlists[station] = newPlaylist.ID.String()
		log.Printf("Created playlist %s for station %s", newPlaylist.ID, station)
		return &newPlaylist.SimplePlaylist, nil
	}

	return &result.SimplePlaylist, nil
}
