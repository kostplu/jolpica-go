package f1

import "fmt"

func (c *Client) GetConstructors(season string) ([]Constructor, error) {
	var result struct {
		MRData struct {
			mrData
			ConstructorTable constructorTable `json:"ConstructorTable"`
		} `json:"MRData"`
	}

	path := fmt.Sprintf("%s/constructors.json", season)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.ConstructorTable.Constructors, nil
}
