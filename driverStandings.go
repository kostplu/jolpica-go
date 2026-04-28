package f1

import (
	"strconv"
	"time"
)

type DriverStandingPage struct {
	DriverStandings []DriverStanding
	PageInfo        PageInfo
}

func (c *Client) GetDriverStandings(opts ...Option) (*DriverStandingPage, error) {
	var result apiResponse

	path := buildPath("driverStandings", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	total, _ := strconv.Atoi(result.MRData.Total)
	limit, _ := strconv.Atoi(result.MRData.Limit)
	offset, _ := strconv.Atoi(result.MRData.Offset)

	return &DriverStandingPage{
		DriverStandings: result.MRData.StandingsTable.StandingsLists[0].DriverStandings,
		PageInfo: PageInfo{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		},
	}, nil
}

func (c *Client) GetAllDriverStandings(opts ...Option) ([]DriverStanding, error) {
	var all []DriverStanding
	offset := 0

	for {
		page, err := c.GetDriverStandings(append(opts, WithLimit(100), WithOffset(offset))...)
		if err != nil {
			return nil, err
		}

		all = append(all, page.DriverStandings...)

		if !page.PageInfo.HasNext() {
			break
		}

		offset = page.PageInfo.NextOffset()
		time.Sleep(200 * time.Millisecond) // be a good citizen
	}

	return all, nil
}
