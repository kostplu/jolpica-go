package f1

func (c *Client) GetRaces(opts ...Option) ([]Race, error) {
	var result apiResponse

	path := buildPath("races", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.RaceTable.Races, nil
}
