package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kostplu/jolpica-go/livetiming"
)

func main() {
	scratch()
}

func scratch() {
	client := livetiming.NewClient(livetiming.WithCache("/tmp/livetiming.db", 24*time.Hour))
	years, err := client.GetAvailableYears()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	for _, year := range years.Years {
		meetings, err := client.GetAvailableMeetings(year.Path)
		if err != err {
			fmt.Printf("error: %v\n", err)
		}
		firstRace := meetings.Meetings[2]
		for _, session := range firstRace.Sessions {
			if session.Type == "Qualifying" {
				qualifyingFeeds, err := client.GetRaceFeeds(session.Path)
				if err != nil {
					fmt.Printf("error: %v\n", err)
				}

				_, err = client.GetFeedRaw(session.Path + qualifyingFeeds.Feeds.PositionZ.StreamPath)
				if err != nil {
					fmt.Printf("error: %v\n", err)
				}
			}
			if session.Type == "Race" {
				raceFeeds, err := client.GetRaceFeeds(session.Path)
				if err != nil {
					fmt.Printf("error: %v\n", err)
				}

				data, err := client.GetFeedRaw(session.Path + raceFeeds.Feeds.PositionZ.StreamPath)
				if err != nil {
					fmt.Printf("error: %v\n", err)
				}
				feed, err := livetiming.ParseFeed[livetiming.PositionZ](data)
				if err != nil {
					fmt.Printf("error: %v\n", err)
				}
				fmt.Printf("parsed %d entries\n", len(feed))

				feedStream := livetiming.StreamFeed(feed)
				clockStream := livetiming.ReplayFeed(feedStream, livetiming.ReplayConfig{StartTime: 30 * time.Minute, Speed: 5.0})
				for frame := range clockStream {
					for _, position := range frame.Position {
						for key, entry := range position.Entries {
							fmt.Printf("Car: %s: X: %s, Y: %s", key, strconv.Itoa(entry.X), strconv.Itoa(entry.Y))
						}
						// livetiming.PrettyLog(position.Entries)
					}
				}
			}
		}
	}
}
