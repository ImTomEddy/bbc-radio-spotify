package broadcasts

import (
	"bbc-radio-spotify/internal/constants"
	"bbc-radio-spotify/internal/request"
	"bbc-radio-spotify/internal/structures"
	"encoding/json"
	"fmt"
)

//GetLatestBroadcast gets the latest broadcast by the specified radio station
func GetLatestBroadcast(station string) (*structures.Broadcast, error) {
	body, err := request.MakeGetRequest(fmt.Sprintf(constants.BroadcastRequestURI, station))

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
