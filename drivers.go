package f1

import "fmt"

func (c *Client) GetDrivers(season string) ([]Driver, error) {
	var result struct {
		MRData struct {
			mrData
			DriverTable driverTable `json:"DriverTable"`
		} `json:"MRData"`
	}

	path := fmt.Sprintf("%s/drivers.json", season)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	return result.MRData.DriverTable.Drivers, nil
}
