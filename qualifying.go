package f1

import (
	"strconv"
	"time"
)

type QualifyingResultPage struct {
	Races    []Race
	PageInfo PageInfo
}

func (c *Client) GetQualifyingResults(opts ...Option) (*QualifyingResultPage, error) {
	var result apiResponse

	path := buildPath("qualifying", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	total, _ := strconv.Atoi(result.MRData.Total)
	limit, _ := strconv.Atoi(result.MRData.Limit)
	offset, _ := strconv.Atoi(result.MRData.Offset)

	return &QualifyingResultPage{
		Races: result.MRData.RaceTable.Races,
		PageInfo: PageInfo{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		},
	}, nil
}

func (c *Client) GetAllQualifyingResults(opts ...Option) ([]Race, error) {
	var all []Race
	offset := 0

	for {
		page, err := c.GetQualifyingResults(append(opts, WithLimit(100), WithOffset(offset))...)
		if err != nil {
			return nil, err
		}

		all = append(all, page.Races...)

		if !page.PageInfo.HasNext() {
			break
		}

		offset = page.PageInfo.NextOffset()
		time.Sleep(200 * time.Millisecond) // be a good citizen
	}

	return all, nil
}
