package f1

func (c *Client) GetConstructorStandings(opts ...Option) ([]ConstructorStanding, error) {
	var result apiResponse

	path := buildPath("constructorStandings", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.StandingsTable.StandingsLists[0].ConstructorStandings, nil
}
