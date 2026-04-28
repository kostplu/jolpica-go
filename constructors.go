package f1

import (
	"strconv"
	"time"
)

type ConstructorPage struct {
	Constructors []Constructor
	PageInfo     PageInfo
}

func (c *Client) GetConstructors(opts ...Option) (*ConstructorPage, error) {
	var result apiResponse

	path := buildPath("constructors", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	total, _ := strconv.Atoi(result.MRData.Total)
	limit, _ := strconv.Atoi(result.MRData.Limit)
	offset, _ := strconv.Atoi(result.MRData.Offset)

	return &ConstructorPage{
		Constructors: result.MRData.ConstructorTable.Constructors,
		PageInfo: PageInfo{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		},
	}, nil
}

func (c *Client) GetAllConstructors(opts ...Option) ([]Constructor, error) {
	var all []Constructor
	offset := 0

	for {
		page, err := c.GetConstructors(append(opts, WithLimit(100), WithOffset(offset))...)
		if err != nil {
			return nil, err
		}

		all = append(all, page.Constructors...)

		if !page.PageInfo.HasNext() {
			break
		}

		offset = page.PageInfo.NextOffset()
		time.Sleep(200 * time.Millisecond) // be a good citizen
	}

	return all, nil
}
