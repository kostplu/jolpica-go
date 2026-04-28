package f1

import (
	"strconv"
	"time"
)

type CircuitsPage struct {
	Circuits []Circuit
	PageInfo PageInfo
}

func (c *Client) GetCircuits(opts ...Option) (*CircuitsPage, error) {
	var result apiResponse

	path := buildPath("circuits", opts)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}

	total, _ := strconv.Atoi(result.MRData.Total)
	limit, _ := strconv.Atoi(result.MRData.Limit)
	offset, _ := strconv.Atoi(result.MRData.Offset)

	return &CircuitsPage{
		Circuits: result.MRData.CircuitTable.Circuits,
		PageInfo: PageInfo{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		},
	}, nil
}

func (c *Client) GetAllCircuits(opts ...Option) ([]Circuit, error) {
	var all []Circuit
	offset := 0

	for {
		page, err := c.GetCircuits(append(opts, WithLimit(100), WithOffset(offset))...)
		if err != nil {
			return nil, err
		}

		all = append(all, page.Circuits...)

		if !page.PageInfo.HasNext() {
			break
		}

		offset = page.PageInfo.NextOffset()
		time.Sleep(200 * time.Millisecond) // be a good citizen
	}

	return all, nil
}
