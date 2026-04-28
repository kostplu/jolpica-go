package f1

import "fmt"

func (c *Client) GetRaces(season string) ([]Race, error) {
	var result struct {
		MRData struct {
			mrData
			RaceTable raceTable `json:"RaceTable"`
		} `json:"MRData"`
	}

	path := fmt.Sprintf("%s/races.json", season)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.RaceTable.Races, nil
}
