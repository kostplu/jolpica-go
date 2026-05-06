package livetiming

func (c *Client) GetAvailableYears() (*YearsResponse, error) {
	var result YearsResponse
	if err := c.get("Index.json", &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAvailableMeetings(path string) (*MeetingsResponse, error) {
	var result MeetingsResponse
	if err := c.get(path+"Index.json", &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetRaceFeeds(path string) (*FeedsResponse, error) {
	var result FeedsResponse
	if err := c.get(path+"Index.json", &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetFeed(path string, dest any) error {
	if err := c.get(path, dest); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetFeedRaw(path string) ([]byte, error) {
	data, err := c.getRaw(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}
