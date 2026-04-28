package f1

import (
	"strconv"
	"time"
)

type ConstructorStandingPage struct {
	ConstructorStandings []ConstructorStanding
	PageInfo             PageInfo
}

func (c *Client) GetConstructorStandings(opts ...Option) (*ConstructorStandingPage, error) {
	var result apiResponse

	path := buildPath("constructorStandings", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	total, _ := strconv.Atoi(result.MRData.Total)
	limit, _ := strconv.Atoi(result.MRData.Limit)
	offset, _ := strconv.Atoi(result.MRData.Offset)

	return &ConstructorStandingPage{
		ConstructorStandings: result.MRData.StandingsTable.StandingsLists[0].ConstructorStandings,
		PageInfo: PageInfo{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		},
	}, nil
}

func (c *Client) GetAllConstructorStandings(opts ...Option) ([]ConstructorStanding, error) {
	var all []ConstructorStanding
	offset := 0

	for {
		page, err := c.GetConstructorStandings(append(opts, WithLimit(100), WithOffset(offset))...)
		if err != nil {
			return nil, err
		}

		all = append(all, page.ConstructorStandings...)

		if !page.PageInfo.HasNext() {
			break
		}

		offset = page.PageInfo.NextOffset()
		time.Sleep(200 * time.Millisecond) // be a good citizen
	}

	return all, nil
}
