package sounds

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

const broadcastRequestURI = "https://rms.api.bbc.co.uk/v2/broadcasts/poll/%s?limit=1"

//UserID defines the spotify user's ID
var UserID = os.Getenv("SPOTIFY_USER_ID")

//Stations define the API Paths for the BBC Radio Stations
var Stations = [...]string{
	"bbc_radio_one",
	"bbc_radio_one_dance",
	"bbc_1xtra",
	"bbc_radio_two",
}

//StationNames are the nice names for stations
var StationNames map[string]string = map[string]string{
	"bbc_radio_one":       "BBC Radio 1",
	"bbc_radio_one_dance": "BBC Radio 1 Dance",
	"bbc_1xtra":           "BBC Radio 1Xtra",
	"bbc_radio_two":       "BBC Radio Two",
}

//GetLatestBroadcast gets the latest broadcast by the specified radio station
func GetLatestBroadcast(station string) (*Broadcast, error) {
	body, err := makeGetRequest(fmt.Sprintf(broadcastRequestURI, station))

	if err != nil {
		return nil, errors.Wrap(err, "Error getting broadcast for station "+station)
	}

	var jsonBody struct {
		Data []Broadcast
	}

	err = json.Unmarshal(body, &jsonBody)

	if err != nil {
		return nil, errors.Wrap(err, "Error unmarshaling for station "+station)
	}

	if len(jsonBody.Data) != 1 {
		return nil, fmt.Errorf("Recieved %d values for the broadcast, expected 1 - for station %s", len(jsonBody.Data), station)
	}

	return &jsonBody.Data[0], nil
}
