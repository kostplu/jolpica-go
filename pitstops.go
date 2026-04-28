package f1

import "fmt"

func (c *Client) getPitStops(season, round string) ([]PitStop, error) {
	var result apiResponse

	path := fmt.Sprintf("%s/%s/pitstops.json", season, round)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.RaceTable.Races[0].PitStops, nil
}
