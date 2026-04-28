package f1

func (c *Client) GetResults(opts ...Option) ([]Result, error) {
	var result apiResponse

	path := buildPath("results", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.RaceTable.Races[0].Results, nil
}
