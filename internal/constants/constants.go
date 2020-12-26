package constants

import "os"

//SongRequestURI defines the BBC Radio Current Song API
const SongRequestURI = "https://rms.api.bbc.co.uk/v2/services/%s/segments/latest?experience=domestic&offset=0&limit=1"

//BroadcastRequestURI defines the BBC Radio Current Broadcast API
const BroadcastRequestURI = "https://rms.api.bbc.co.uk/v2/broadcasts/poll/%s?limit=1"

//UserID defines the spotify user's ID
var UserID = os.Getenv("SPOTIFY_USER_ID")

//Stations define the API Paths for the BBC Radio Stations
var Stations = [...]string{
	"bbc_radio_one",
	"bbc_radio_one_dance",
	"bbc_1xtra",
	"bbc_radio_two",
	"bbc_radio_three",
	"bbc_6music",
	"bbc_asian_network",
	"bbc_radio_wales_fm",
}

//StationNames are the nice names for stations
var StationNames map[string]string = map[string]string{
	"bbc_radio_one":       "BBC Radio 1",
	"bbc_radio_one_dance": "BBC Radio 1 Dance",
	"bbc_1xtra":           "BBC Radio 1Xtra",
	"bbc_radio_two":       "BBC Radio Two",
	"bbc_radio_three":     "BBC Radio Three",
	"bbc_6music":          "BBC 6 Music",
	"bbc_asian_network":   "BBC Asian Network",
	"bbc_radio_wales_fm":  "BBC Radio Wales",
}
