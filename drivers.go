package f1

func (c *Client) GetDrivers(opts ...Option) ([]Driver, error) {
	var result apiResponse

	path := buildPath("drivers", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.DriverTable.Drivers, nil
}
