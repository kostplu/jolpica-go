package f1

func (c *Client) GetSeasons() ([]Season, error) {
	var result apiResponse

	if err := c.get("seasons.json", &result); err != nil {
		return nil, err
	}

	return result.MRData.SeasonTable.Seasons, nil
}
