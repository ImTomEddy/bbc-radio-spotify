package request

import (
	"encoding/json"
	"fmt"

	"github.com/imtomeddy/bbc-radio-spotify/internal/constants"
	"github.com/imtomeddy/bbc-radio-spotify/internal/structures"
)

//GetLatestBroadcast gets the latest broadcast by the specified radio station
func GetLatestBroadcast(station string) (*structures.Broadcast, error) {
	body, err := MakeGetRequest(fmt.Sprintf(constants.BroadcastRequestURI, station))

	if err != nil {
		return nil, err
	}

	var jsonBody struct {
		Data []structures.Broadcast
	}

	err = json.Unmarshal(body, &jsonBody)

	if err != nil {
		return nil, err
	}

	if len(jsonBody.Data) != 1 {
		return nil, fmt.Errorf("Recieved %d values, expected 1", len(jsonBody.Data))
	}

	return &jsonBody.Data[0], nil
}
