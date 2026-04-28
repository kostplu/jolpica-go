package f1

func (c *Client) GetCircuits() ([]Circuit, error) {
	var result apiResponse

	if err := c.get("circuits.json", &result); err != nil {
		return nil, err
	}

	return result.MRData.CircuitTable.Circuits, nil
}
