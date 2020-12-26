package songs

import (
	"bbc-radio-spotify/internal/constants"
	"bbc-radio-spotify/internal/request"
	"bbc-radio-spotify/internal/structures"
	"encoding/json"
	"fmt"
)

//GetLatestSong gets the latest song played by the specified radio station
func GetLatestSong(station string) (*structures.Title, error) {
	body, err := request.MakeGetRequest(fmt.Sprintf(constants.SongRequestURI, station))

	if err != nil {
		return nil, err
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
		return nil, fmt.Errorf("Recieved %d values, expected 1", len(jsonBody.Data))
	}

	return &jsonBody.Data[0].Song, nil
}
