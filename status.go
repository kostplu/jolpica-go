package f1

import (
	"strconv"
	"time"
)

type StatusPage struct {
	Statuses []Status
	PageInfo PageInfo
}

func (c *Client) GetStatus(opts ...Option) (*StatusPage, error) {
	var result apiResponse

	path := buildPath("status", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	total, _ := strconv.Atoi(result.MRData.Total)
	limit, _ := strconv.Atoi(result.MRData.Limit)
	offset, _ := strconv.Atoi(result.MRData.Offset)

	return &StatusPage{
		Statuses: result.MRData.StatusTable.Status,
		PageInfo: PageInfo{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		},
	}, nil
}

func (c *Client) GetAllStatus(opts ...Option) ([]Status, error) {
	var all []Status
	offset := 0

	for {
		page, err := c.GetStatus(append(opts, WithLimit(100), WithOffset(offset))...)
		if err != nil {
			return nil, err
		}

		all = append(all, page.Statuses...)

		if !page.PageInfo.HasNext() {
			break
		}

		offset = page.PageInfo.NextOffset()
		time.Sleep(200 * time.Millisecond) // be a good citizen
	}

	return all, nil
}
