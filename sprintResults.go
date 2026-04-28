package f1

func (c *Client) GetSprintResults(opts ...Option) ([]Result, error) {
	var result apiResponse

	path := buildPath("sprintresults", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.RaceTable.Races[0].SprintResults, nil
}
