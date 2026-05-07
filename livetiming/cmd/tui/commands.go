package main

import (
	"encoding/json"
	"fmt"
	"math"
	"slices"
	"strings"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/kostplu/jolpica-go/livetiming"
)

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

		return streaming{mapDataJSON.X, mapDataJSON.Y, mapDataJSON.Rotation}
	}
}

func renderTrack(xs, ys []int, width, height int, degrees float64) string {
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
	}

	rx, ry := rotatePoints(xs, ys, degrees)

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
			sb.WriteRune(quadrant[idx])
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func rotatePoints(xs, ys []int, degrees float64) ([]float64, []float64) {
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

	return rx, ry
}
