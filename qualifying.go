package f1

func (c *Client) GetQualifyingResults(opts ...Option) ([]QualifyingResult, error) {
	var result apiResponse

	path := buildPath("qualifying", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.RaceTable.Races[0].QualifyingResults, nil
}
