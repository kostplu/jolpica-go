package f1

func (c *Client) GetCircuits() ([]Circuit, error) {
	var result struct {
		MRData struct {
			mrData
			CircuitTable circuitTable `json:"CircuitTable"`
		} `json:"MRData"`
	}

	if err := c.get("circuits.json", &result); err != nil {
		return nil, err
	}

	return result.MRData.CircuitTable.Circuits, nil
}
