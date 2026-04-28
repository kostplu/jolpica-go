package f1

import "fmt"

func (c *Client) GetQualifyingResults(season string, round int) ([]QualifyingResult, error) {
	var result apiResponse

	path := fmt.Sprintf("%s/qualifying.json", season)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.RaceTable.Races[round].QualifyingResults, nil
}
