package request

import (
	"encoding/json"
	"fmt"

	"github.com/imtomeddy/bbc-radio-spotify/internal/constants"
	"github.com/imtomeddy/bbc-radio-spotify/internal/structures"
	"github.com/pkg/errors"
)

//GetLatestBroadcast gets the latest broadcast by the specified radio station
func GetLatestBroadcast(station string) (*structures.Broadcast, error) {
	body, err := MakeGetRequest(fmt.Sprintf(constants.BroadcastRequestURI, station))

	if err != nil {
		return nil, errors.Wrap(err, "Error getting broadcast for station "+station)
	}

	var jsonBody struct {
		Data []structures.Broadcast
	}

	err = json.Unmarshal(body, &jsonBody)

	if err != nil {
		return nil, errors.Wrap(err, "Error unmarshaling for station "+station)
	}

	if len(jsonBody.Data) != 1 {
		return nil, fmt.Errorf("Recieved %d values, expected 1 - for station %s", len(jsonBody.Data), station)
	}

	return &jsonBody.Data[0], nil
}
