package httpclient

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

type cache struct {
	db  *sql.DB
	ttl time.Duration
}

func newCache(path string, ttl time.Duration) (*cache, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	// create the cache table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS cache (
			key TEXT PRIMARY KEY,
			value TEXT NOT NULL,
			cached_at TIMESTAMP NOT NULL)`)
	if err != nil {
		return nil, err
	}

	return &cache{db: db, ttl: ttl}, nil
}

func (c *cache) get(key string) (string, bool) {
	var value string
	var cachedAt time.Time

	err := c.db.QueryRow("SELECT value, cached_at FROM cache WHERE key = ?", key).Scan(&value, &cachedAt)
	if err != nil {
		return "", false // cache miss
	}

	if time.Since(cachedAt) > c.ttl {
		return "", false // cache expired
	}

	return value, true // cache hit
}

func (c *cache) set(key, value string) error {
	_, err := c.db.Exec(
		`INSERT OR REPLACE INTO cache (key, value, cached_at) VALUES (?, ?, ?)`,
		key, value, time.Now())
	return err
}
