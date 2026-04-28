package f1

import (
	"strconv"
	"time"
)

type DriversPage struct {
	Drivers  []Driver
	PageInfo PageInfo
}

func (c *Client) GetDrivers(opts ...Option) (*DriversPage, error) {
	var result apiResponse

	path := buildPath("drivers", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	total, _ := strconv.Atoi(result.MRData.Total)
	limit, _ := strconv.Atoi(result.MRData.Limit)
	offset, _ := strconv.Atoi(result.MRData.Offset)

	return &DriversPage{
		Drivers: result.MRData.DriverTable.Drivers,
		PageInfo: PageInfo{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		},
	}, nil
}

func (c *Client) GetAllDrivers(opts ...Option) ([]Driver, error) {
	var all []Driver
	offset := 0

	for {
		page, err := c.GetDrivers(append(opts, WithLimit(100), WithOffset(offset))...)
		if err != nil {
			return nil, err
		}

		all = append(all, page.Drivers...)

		if !page.PageInfo.HasNext() {
			break
		}

		offset = page.PageInfo.NextOffset()
		time.Sleep(200 * time.Millisecond) // be a good citizen
	}

	return all, nil
}
