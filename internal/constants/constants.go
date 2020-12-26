package constants

// SongRequestURI defines the BBC Radio Current Song API
const SongRequestURI = "https://rms.api.bbc.co.uk/v2/services/%s/segments/latest?experience=domestic&offset=0&limit=1"

// BroadcastRequestURI defines the BBC Radio Current Broadcast API
const BroadcastRequestURI = "https://rms.api.bbc.co.uk/v2/broadcasts/poll/%s?limit=1"

// Stations define the API Paths for the BBC Radio Stations
var Stations = [...]string{
	"bbc_radio_one",
	"bbc_radio_one_dance",
	"bbc_1xtra",
	"bbc_radio_two",
	"bbc_radio_three",
	"bbc_6music",
	"bbc_asian_network",
}
