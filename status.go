package f1

func (c *Client) GetStatus() ([]Status, error) {
	var result struct {
		MRData struct {
			mrData
			StatusTable statusTable `json:"StatusTable"`
		} `json:"MRData"`
	}

	if err := c.get("status.json", &result); err != nil {
		return nil, err
	}

	return result.MRData.StatusTable.Status, nil
}
