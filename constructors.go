package f1

func (c *Client) GetConstructors(opts ...Option) ([]Constructor, error) {
	var result apiResponse

	path := buildPath("constructors", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.ConstructorTable.Constructors, nil
}
