package f1

func (c *Client) GetDriverStandings(opts ...Option) ([]DriverStanding, error) {
	var result apiResponse

	path := buildPath("driverStandings", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.StandingsTable.StandingsLists[0].DriverStandings, nil
}
