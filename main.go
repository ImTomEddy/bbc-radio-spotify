package main

import (
	"log"

	"github.com/imtomeddy/bbc-radio-spotify/internal/constants"
	"github.com/imtomeddy/bbc-radio-spotify/internal/jobs"
	"github.com/robfig/cron/v3"
)

func main() {
	log.Println("Starting BBC Radio to Spotify Service")
	log.Printf("Supported Radio Stations: (%d) %s", len(constants.Stations), constants.Stations)

	c := cron.New()
	c.AddFunc("* * * * *", jobs.UpdateInfo)
	c.Start()

	select {}
}
