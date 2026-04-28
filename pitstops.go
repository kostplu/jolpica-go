package f1

import "fmt"

func (c *Client) getPitStops(season, round string) ([]PitStop, error) {
	var result struct {
		MRData struct {
			mrData
			RaceTable raceTable `json:"RaceTable"`
		} `json:"MRData"`
	}

	path := fmt.Sprintf("%s/%s/pitstops.json", season, round)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.RaceTable.Races[0].PitStops, nil
}
