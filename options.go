package f1

import "fmt"

type requestOptions struct {
	season      int
	round       int
	constructor string
	driver      string
	limit       int
	offest      int
}

// Option is a function that modifies the request options.
type Option func(*requestOptions)

func buildPath(endpoint string, opts []Option) string {
	o := &requestOptions{
		limit:  30, // default limit
		offest: 0,  // default offset
	}

	for _, opt := range opts {
		opt(o)
	}

	path := ""

	if o.season != 0 {
		path += fmt.Sprintf("/%d", o.season)
	}
	if o.round != 0 {
		path += fmt.Sprintf("/%d", o.round)
	}
	if o.constructor != "" {
		path += fmt.Sprintf("/constructors/%s", o.constructor)
	}
	if o.driver != "" {
		path += fmt.Sprintf("/drivers/%s", o.driver)
	}

	path += endpoint + ".json"

	path += fmt.Sprintf("?limit=%d&offset=%d", o.limit, o.offest)

	return path
}

func WithSeason(season int) Option {
	return func(opts *requestOptions) {
		opts.season = season
	}
}

func WithRound(round int) Option {
	return func(opts *requestOptions) {
		opts.round = round
	}
}

func WithConstructor(constructor string) Option {
	return func(opts *requestOptions) {
		opts.constructor = constructor
	}
}

func WithDriver(driver string) Option {
	return func(opts *requestOptions) {
		opts.driver = driver
	}
}

func WithLimit(limit int) Option {
	return func(opts *requestOptions) {
		opts.limit = limit
	}
}

func WithOffset(offset int) Option {
	return func(opts *requestOptions) {
		opts.offest = offset
	}
}
