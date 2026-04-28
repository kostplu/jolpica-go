package f1

func (c *Client) GetLaps(opts ...Option) ([]Lap, error) {
	var result apiResponse

	path := buildPath("laps", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.RaceTable.Races[0].Laps, nil
}
