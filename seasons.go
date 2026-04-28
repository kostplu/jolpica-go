package f1

func (c *Client) GetSeasons(opts ...Option) ([]Season, error) {
	var result apiResponse

	path := buildPath("seasons", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.SeasonTable.Seasons, nil
}
