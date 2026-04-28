package f1

import (
	"strconv"
	"time"
)

type QualifyingResultPage struct {
	QualifyingResults []QualifyingResult
	PageInfo          PageInfo
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
		QualifyingResults: result.MRData.RaceTable.Races[0].QualifyingResults,
		PageInfo: PageInfo{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		},
	}, nil
}

func (c *Client) GetAllQualifyingResults(opts ...Option) ([]QualifyingResult, error) {
	var all []QualifyingResult
	offset := 0

	for {
		page, err := c.GetQualifyingResults(append(opts, WithLimit(100), WithOffset(offset))...)
		if err != nil {
			return nil, err
		}

		all = append(all, page.QualifyingResults...)

		if !page.PageInfo.HasNext() {
			break
		}

		offset = page.PageInfo.NextOffset()
		time.Sleep(200 * time.Millisecond) // be a good citizen
	}

	return all, nil
}
