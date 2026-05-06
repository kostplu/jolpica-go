package f1

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// NewTestClient() creates a new Client with a test HTTP server that uses the provided handler.
func newTestClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(func() { server.Close() })

	return &Client{
		baseURL: server.URL + "/",
	}
}

func fixtureHandler(t *testing.T, path string) http.HandlerFunc {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read fixture file: %v", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
