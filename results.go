package f1

import "fmt"

func (c *Client) GetResults(season string) ([]Result, error) {
	var result struct {
		MRData struct {
			mrData
			RaceTable raceTable `json:"RaceTable"`
		} `json:"MRData"`
	}

	path := fmt.Sprintf("%s/results.json", season)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.RaceTable.Races[0].Results, nil
}
