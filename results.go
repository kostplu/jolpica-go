package f1

import (
	"strconv"
	"time"
)

type RaceResultPage struct {
	Results  []Result
	PageInfo PageInfo
}

func (c *Client) GetResults(opts ...Option) (*RaceResultPage, error) {
	var result apiResponse

	path := buildPath("results", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	total, _ := strconv.Atoi(result.MRData.Total)
	limit, _ := strconv.Atoi(result.MRData.Limit)
	offset, _ := strconv.Atoi(result.MRData.Offset)

	return &RaceResultPage{
		Results: result.MRData.RaceTable.Races[0].Results,
		PageInfo: PageInfo{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		},
	}, nil
}

func (c *Client) GetAllResults(opts ...Option) ([]Result, error) {
	var all []Result
	offset := 0

	for {
		page, err := c.GetResults(append(opts, WithLimit(100), WithOffset(offset))...)
		if err != nil {
			return nil, err
		}

		all = append(all, page.Results...)

		if !page.PageInfo.HasNext() {
			break
		}

		offset = page.PageInfo.NextOffset()
		time.Sleep(200 * time.Millisecond) // be a good citizen
	}

	return all, nil
}
