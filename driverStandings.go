package f1

import "fmt"

func (c *Client) GetDriverStandings(season string) ([]DriverStanding, error) {
	var result struct {
		MRData struct {
			mrData
			StandingsTable standingsTable `json:"StandingsTable"`
		} `json:"MRData"`
	}

	path := fmt.Sprintf("%s/driverStandings.json", season)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.StandingsTable.StandingsLists[0].DriverStandings, nil
}
