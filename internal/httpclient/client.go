package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	cache      *cache
}

type Option func(*Client)

func NewClient(opts ...Option) *Client {
	c := &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithCache(path string, ttl time.Duration) Option {
	return func(c *Client) {
		cache, err := newCache(path, ttl)
		if err != nil {
			return
		}
		c.cache = cache
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

func (c *Client) Get(url string, dest any) error {
	if c.cache != nil {
		if cached, found := c.cache.get(url); found {
			return json.Unmarshal([]byte(cached), dest)
		}
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	if c.cache != nil {
		c.cache.set(url, string(body))
	}

	return json.Unmarshal(body, dest)
}

func (c *Client) GetRaw(url string) ([]byte, error) {
	if c.cache != nil {
		if cached, found := c.cache.get(url); found {
			return []byte(cached), nil
		}
	}
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	if c.cache != nil {
		c.cache.set(url, string(body))
	}

	return body, nil
}
