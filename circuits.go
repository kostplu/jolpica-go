package f1

func (c *Client) GetCircuits(opts ...Option) ([]Circuit, error) {
	var result apiResponse

	path := buildPath("circuits", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.CircuitTable.Circuits, nil
}
