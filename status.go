package f1

func (c *Client) GetStatus() ([]Status, error) {
	var result apiResponse

	if err := c.get("status.json", &result); err != nil {
		return nil, err
	}

	return result.MRData.StatusTable.Status, nil
}
