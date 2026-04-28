package f1

import (
	"strconv"
	"time"
)

type SeasonPage struct {
	Seasons  []Season
	PageInfo PageInfo
}

func (c *Client) GetSeasons(opts ...Option) (*SeasonPage, error) {
	var result apiResponse

	path := buildPath("seasons", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	total, _ := strconv.Atoi(result.MRData.Total)
	limit, _ := strconv.Atoi(result.MRData.Limit)
	offset, _ := strconv.Atoi(result.MRData.Offset)

	return &SeasonPage{
		Seasons: result.MRData.SeasonTable.Seasons,
		PageInfo: PageInfo{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		},
	}, nil
}

func (c *Client) GetAllSeasons(opts ...Option) ([]Season, error) {
	var all []Season
	offset := 0

	for {
		page, err := c.GetSeasons(append(opts, WithLimit(100), WithOffset(offset))...)
		if err != nil {
			return nil, err
		}

		all = append(all, page.Seasons...)

		if !page.PageInfo.HasNext() {
			break
		}

		offset = page.PageInfo.NextOffset()
		time.Sleep(200 * time.Millisecond) // be a good citizen
	}

	return all, nil
}
