package request

import (
	"encoding/json"
	"fmt"

	"github.com/imtomeddy/bbc-radio-spotify/internal/constants"
	"github.com/imtomeddy/bbc-radio-spotify/internal/structures"
	"github.com/pkg/errors"
)

//GetLatestSong gets the latest song played by the specified radio station
func GetLatestSong(station string) (*structures.Title, error) {
	body, err := MakeGetRequest(fmt.Sprintf(constants.SongRequestURI, station))

	if err != nil {
		return nil, errors.Wrap(err, "Error getting song for station "+station)
	}

	var jsonBody struct {
		Data []struct {
			Song structures.Title `json:"titles"`
		} `json:"data"`
	}

	err = json.Unmarshal(body, &jsonBody)

	if err != nil {
		return nil, err
	}

	if len(jsonBody.Data) != 1 {
		return nil, fmt.Errorf("Recieved %d values for the song, expected 1 - for station %s", len(jsonBody.Data), station)
	}

	return &jsonBody.Data[0].Song, nil
}
