package f1

import "fmt"

func (c *Client) GetLaps(season, round string) ([]Lap, error) {
	var result apiResponse

	path := fmt.Sprintf("%s/%s/laps.json", season, round)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.RaceTable.Races[0].Laps, nil
}
