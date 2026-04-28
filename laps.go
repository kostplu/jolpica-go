package f1

import (
	"strconv"
	"time"
)

type LapPage struct {
	Laps     []Lap
	PageInfo PageInfo
}

func (c *Client) GetLaps(opts ...Option) (*LapPage, error) {
	var result apiResponse

	path := buildPath("laps", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	total, _ := strconv.Atoi(result.MRData.Total)
	limit, _ := strconv.Atoi(result.MRData.Limit)
	offset, _ := strconv.Atoi(result.MRData.Offset)

	return &LapPage{
		Laps: result.MRData.RaceTable.Races[0].Laps,
		PageInfo: PageInfo{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		},
	}, nil
}

func (c *Client) GetAllLaps(opts ...Option) ([]Lap, error) {
	var all []Lap
	offset := 0

	for {
		page, err := c.GetLaps(append(opts, WithLimit(100), WithOffset(offset))...)
		if err != nil {
			return nil, err
		}

		all = append(all, page.Laps...)

		if !page.PageInfo.HasNext() {
			break
		}

		offset = page.PageInfo.NextOffset()
		time.Sleep(200 * time.Millisecond) // be a good citizen
	}

	return all, nil
}
