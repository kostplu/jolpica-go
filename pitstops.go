package f1

func (c *Client) GetPitStops(opts ...Option) ([]PitStop, error) {
	var result apiResponse

	path := buildPath("pitstops", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.RaceTable.Races[0].PitStops, nil
}
