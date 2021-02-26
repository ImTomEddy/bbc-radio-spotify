package sounds

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

const songRequestURI = "https://rms.api.bbc.co.uk/v2/services/%s/segments/latest?experience=domestic&offset=0&limit=1"

//GetLatestSong gets the latest song played by the specified radio station
func GetLatestSong(station string) (*Title, error) {
	body, err := makeGetRequest(fmt.Sprintf(songRequestURI, station))

	if err != nil {
		return nil, errors.Wrap(err, "Error getting song for station "+station)
	}

	var jsonBody struct {
		Data []struct {
			Song Title `json:"titles"`
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
