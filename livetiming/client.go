package livetiming

import (
	"fmt"
	"sync"
	"time"

	"github.com/kostplu/jolpica-go/internal/httpclient"
)

const baseURL = "https://livetiming.formula1.com/static/"

type Client struct {
	baseURL   string
	httpsOpts []httpclient.Option
	http      *httpclient.Client
	httpOnce  sync.Once
}

type ClientOption func(*Client)

func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		baseURL: baseURL,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithCache(path string, ttl time.Duration) ClientOption {
	return func(c *Client) {
		c.httpsOpts = append(c.httpsOpts, httpclient.WithCache(path, ttl))
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpsOpts = append(c.httpsOpts, httpclient.WithTimeout(timeout))
	}
}

func (c *Client) getHTTP() *httpclient.Client {
	c.httpOnce.Do(func() {
		c.http = httpclient.NewClient(c.httpsOpts...)
	})
	return c.http
}

func (c *Client) get(path string, dest any) error {
	fmt.Println(path)
	return c.getHTTP().Get(c.baseURL+path, dest)
}

func (c *Client) getRaw(path string) ([]byte, error) {
	fmt.Println(path)
	return c.getHTTP().GetRaw(c.baseURL + path)
}
