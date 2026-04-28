package f1

func (c *Client) GetStatus(opts ...Option) ([]Status, error) {
	var result apiResponse

	path := buildPath("status", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.StatusTable.Status, nil
}
