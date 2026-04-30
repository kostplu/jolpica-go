package f1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	defaultBaseURL = "https://api.jolpi.ca/ergast/f1/"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	cache      *cache
}

type ClientOption func(*Client)

func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func NewClientWithBaseURL(baseURL string, opts ...ClientOption) *Client {
	c := &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithCache(path string, ttl time.Duration) ClientOption {
	return func(c *Client) {
		cache, err := newCache(path, ttl)
		if err != nil {
			return
		}
		c.cache = cache
	}
}

func WitTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

func (c *Client) get(path string, target any) error {
	url := c.baseURL + path

	// check cache first
	if c.cache != nil {
		if cached, found := c.cache.get(url); found {
			return json.Unmarshal([]byte(cached), target)
		}
	}

	// cache miss, make HTTP request
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// store in cache
	if c.cache != nil {
		c.cache.set(url, string(body))
	}

	return json.Unmarshal(body, target)
}
