package main

import (
	"encoding/json"
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/kostplu/jolpica-go/livetiming"
)

type trackProjection struct {
	minX, maxX, minY, maxY float64
	scale                  float64
	offsetX, offsetY       float64
	cx, cy                 float64
	degrees                float64
	pixelH                 int
	cellAspect             float64
}

func fetchYears(client *livetiming.Client) tea.Cmd {
	return func() tea.Msg {
		years, err := client.GetAvailableYears()
		if err != nil {
			return errMsg{err}
		}
		items := make([]list.Item, len(years.Years))
		for i, y := range years.Years {
			items[i] = yearItem{path: y.Path, name: fmt.Sprintf("%d", y.Year), year: y.Year}
		}
		return yearsLoadedMsg{items}
	}
}

func fetchMeetings(client *livetiming.Client, yearPath string) tea.Cmd {
	return func() tea.Msg {
		meetings, err := client.GetAvailableMeetings(yearPath)
		if err != nil {
			return errMsg{err}
		}
		items := make([]list.Item, len(meetings.Meetings))
		for i, m := range meetings.Meetings {
			items[i] = meetingItem{meeting: m}
		}
		return meetingsLoadedMsg{items}
	}
}

func fetchSessions(client *livetiming.Client, meeting livetiming.Meeting) tea.Cmd {
	return func() tea.Msg {
		items := make([]list.Item, len(meeting.Sessions))
		for i, s := range meeting.Sessions {
			items[i] = sessionItem{session: s}
		}
		return sessionsLoadedMsg{items}
	}
}

func startReplay(client *livetiming.Client, sessionPath string, session livetiming.Session, selectedYear int) tea.Cmd {
	return func() tea.Msg {
		raceFeeds, err := client.GetRaceFeeds(sessionPath)
		if err != nil {
			return errMsg{err}
		}

		sessionInfoData, err := client.GetFeedRaw(session.Path + raceFeeds.Feeds.SessionInfo.KeyFramePath)
		if err != nil {
			return errMsg{err}
		}
		var sessionInfoJSON livetiming.SessionInfo
		err = json.Unmarshal(sessionInfoData, &sessionInfoJSON)
		if err != nil {
			return errMsg{err}
		}

		mapData, err := client.GetMap(sessionInfoJSON.Meeting.Circuit.Key, selectedYear)
		if err != nil {
			return errMsg{err}
		}
		var mapDataJSON livetiming.Map
		err = json.Unmarshal(mapData, &mapDataJSON)
		if err != nil {
			return errMsg{err}
		}

		data, err := client.GetFeedRaw(session.Path + raceFeeds.Feeds.PositionZ.StreamPath)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		feed, err := livetiming.ParseFeed[livetiming.PositionZ](data)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		config := livetiming.ReplayConfig{StartTime: 59 * time.Minute, Speed: 1}
		// read feed until the start time (e.g. 15 minutes) is reached, then start the feedStream
		// evaluate all read entries to update the initial car locations
		startIdx := 0
		initialLocations := make([]livetiming.CarLocation, 0)

		for i, entry := range feed {
			if entry.Timestamp >= config.StartTime {
				startIdx = i
				break
			}
			// merge entry.Data positions into initialCars
			for _, position := range entry.Data.Position {
				for key, entry := range position.Entries {
					idx := slices.IndexFunc(initialLocations, func(cl livetiming.CarLocation) bool { return cl.DriverNumber == key })
					if idx == -1 {
						initialLocations = append(initialLocations, livetiming.CarLocation{PosX: entry.X, PosY: entry.Y, DriverNumber: key})
					} else {
						initialLocations[idx].PosX = entry.X
						initialLocations[idx].PosY = entry.Y
					}
				}
			}
		}

		feedStream := livetiming.StreamFeed(feed[startIdx:])
		clockStream := livetiming.ReplayFeed(feedStream, config)

		return mapLoaded{mapDataJSON.X, mapDataJSON.Y, mapDataJSON.Rotation, clockStream, initialLocations}
	}
}

func waitForFrames(frames <-chan livetiming.PositionZ, current []livetiming.CarLocation) tea.Cmd {
	return func() tea.Msg {
		// carLocationData := make([]livetiming.CarLocation, 0)

		frame, ok := <-frames
		if !ok {
			return replayDoneMsg{}
		}
		timestamp := time.Time{}
		for _, position := range frame.Position {
			timestamp = position.Timestamp
			for key, entry := range position.Entries {
				idx := slices.IndexFunc(current, func(cl livetiming.CarLocation) bool { return cl.DriverNumber == key })
				if idx == -1 {
					current = append(current, livetiming.CarLocation{PosX: entry.X, PosY: entry.Y, DriverNumber: key})
				} else {
					current[idx].PosX = entry.X
					current[idx].PosY = entry.Y
				}
			}
		}
		return updateLocations{carLocations: current, timestamp: timestamp}
	}
}

func renderTrack(xs, ys []int, width, height int, degrees float64, carLocations []livetiming.CarLocation, timestamp time.Time) string {
	quadrant := []rune{
		' ', // 0000 - empty
		'▘', // 0001 - top-left
		'▝', // 0010 - top-right
		'▀', // 0011 - top half
		'▖', // 0100 - bot-left
		'▌', // 0101 - left half
		'▞', // 0110 - diagonal
		'▛', // 0111 - all except bot-right
		'▗', // 1000 - bot-right
		'▚', // 1001 - diagonal
		'▐', // 1010 - right half
		'▜', // 1011 - all except bot-left
		'▄', // 1100 - bot half
		'▙', // 1101 - all except top-right
		'▟', // 1110 - all except top-left
		'█', // 1111 - full
		'O', // car location
	}

	rx, ry, cx, cy := rotatePoints(xs, ys, degrees)

	const padding = 4

	minX, maxX := slices.Min(rx), slices.Max(rx)
	minY, maxY := slices.Min(ry), slices.Max(ry)

	rangeX := maxX - minX
	rangeY := maxY - minY

	// available space
	availW := (width - padding*2) * 2
	availH := (height - padding*2) * 2

	// aspect ratio correction: terminal cells are ~2x taller than wide
	const cellAspect = 0.5

	// uniform scale — fit the larger dimension, preserve ratio
	scaleX := float64(availW) / rangeX
	scaleY := float64(availH) / rangeY / cellAspect

	scale := math.Min(scaleX, scaleY)

	// center the track in the available space
	offsetX := (float64(availW)-rangeX*scale)/2 + float64(padding*2)
	offsetY := (float64(availH)-rangeY*scale*cellAspect)/2 + float64(padding*2)

	pixels := make([][]bool, availH+padding*4)
	for i := range pixels {
		pixels[i] = make([]bool, availW+padding*4)
	}

	for i := range xs {
		tx := int((rx[i]-minX)*scale + offsetX)
		ty := int((ry[i]-minY)*scale*cellAspect + offsetY)
		ty = (len(pixels) - 1) - ty // flip Y
		if tx >= 0 && tx < len(pixels[0]) && ty >= 0 && ty < len(pixels) {
			pixels[ty][tx] = true
		}
	}
	trackProjection := trackProjection{
		minX:       minX,
		maxX:       maxX,
		minY:       minY,
		maxY:       maxY,
		scale:      scale,
		offsetX:    offsetX,
		offsetY:    offsetY,
		cx:         cx,
		cy:         cy,
		degrees:    degrees,
		pixelH:     len(pixels),
		cellAspect: cellAspect,
	}
	var sb strings.Builder
	for row := 0; row < len(pixels)-1; row += 2 {
		for col := 0; col < len(pixels[0])-1; col += 2 {
			idx := 0
			if pixels[row][col] {
				idx |= 1
			} // top-left
			if pixels[row][col+1] {
				idx |= 2
			} // top-right
			if pixels[row+1][col] {
				idx |= 4
			} // bot-left
			if pixels[row+1][col+1] {
				idx |= 8
			} // bot-right
			if carAtLocation(carLocations, col, row, trackProjection) {
				idx = 16
			}
			sb.WriteRune(quadrant[idx])
		}
		sb.WriteByte('\n')
	}
	header := fmt.Sprintf("T+%v\n", timestamp)
	return header + sb.String()
}

func rotatePoints(xs, ys []int, degrees float64) ([]float64, []float64, float64, float64) {
	rad := degrees * math.Pi / 180.0

	// find center
	minX, maxX := slices.Min(xs), slices.Max(xs)
	minY, maxY := slices.Min(ys), slices.Max(ys)
	cx := float64(minX+maxX) / 2
	cy := float64(minY+maxY) / 2

	rx := make([]float64, len(xs))
	ry := make([]float64, len(ys))

	for i := range xs {
		x := float64(xs[i]) - cx
		y := float64(ys[i]) - cy

		rx[i] = x*math.Cos(rad) - y*math.Sin(rad) + cx
		ry[i] = x*math.Sin(rad) + y*math.Cos(rad) + cy
	}

	return rx, ry, cx, cy
}

func rotatePoint(x, y int, degrees, cx, cy float64) (float64, float64) {
	rad := degrees * math.Pi / 180.0

	nx := float64(x) - cx
	ny := float64(y) - cy

	rx := nx*math.Cos(rad) - ny*math.Sin(rad) + cx
	ry := nx*math.Sin(rad) + ny*math.Cos(rad) + cy
	return rx, ry
}

func carAtLocation(carLocations []livetiming.CarLocation, col, row int, trackProjection trackProjection) bool {
	for _, loc := range carLocations {
		rx, ry := rotatePoint(loc.PosX, loc.PosY, trackProjection.degrees, trackProjection.cx, trackProjection.cy)

		tx := int((rx-trackProjection.minX)*trackProjection.scale + trackProjection.offsetX)
		ty := int((ry-trackProjection.minY)*trackProjection.scale*trackProjection.cellAspect + trackProjection.offsetY)
		ty = (trackProjection.pixelH - 1) - ty // flip Y

		if tx >= col && tx <= col+1 && ty >= row && ty <= row+1 {
			return true
		}
	}
	return false
}
