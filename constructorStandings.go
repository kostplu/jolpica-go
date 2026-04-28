package f1

import "fmt"

func (c *Client) GetConstructorStandings(season string) ([]ConstructorStanding, error) {
	var result struct {
		MRData struct {
			mrData
			StandingsTable standingsTable `json:"StandingsTable"`
		} `json:"MRData"`
	}

	path := fmt.Sprintf("%s/constructorStandings.json", season)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.StandingsTable.StandingsLists[0].ConstructorStandings, nil
}
