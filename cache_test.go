package f1

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestCache_SecondRequestIsCached(t *testing.T) {
	requestCount := 0

	data, err := os.ReadFile("testdata/drivers_2024.json")
	if err != err {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	cacheFile := t.TempDir() + "/cache.db"

	client := NewClientWithBaseURL(server.URL, WithCache(cacheFile, 1*time.Hour))

	_, err = client.GetDrivers(WithSeason(2024))
	if err != nil {
		t.Fatalf("first request failed: %v", err)
	}
	if requestCount != 1 {
		t.Errorf("expected 1 request, got %d", requestCount)
	}

	_, err = client.GetDrivers(WithSeason(2024))
	if err != nil {
		t.Fatalf("second request failed: %v", err)
	}
	if requestCount != 1 {
		t.Errorf("expected 1 request after second call, got %d", requestCount)
	}
}

func TestCache_ExpiredCacheIsRefreshed(t *testing.T) {
	requestCount := 0

	data, err := os.ReadFile("testdata/drivers_2024.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	cacheFile := t.TempDir() + "/cache.db"

	client := NewClientWithBaseURL(server.URL+"/", WithCache(cacheFile, 1*time.Millisecond))

	_, err = client.GetDrivers(WithSeason(2024))
	if err != nil {
		t.Fatalf("first request failed: %v", err)
	}
	if requestCount != 1 {
		t.Errorf("expected 1 request, got %d", requestCount)
	}

	time.Sleep(10 * time.Millisecond)

	_, err = client.GetDrivers(WithSeason(2024))
	if err != nil {
		t.Fatalf("second request failed: %v", err)
	}
	if requestCount != 2 {
		t.Errorf("expected 2 requests after cache expiration, got %d", requestCount)
	}
}

func TestCache_MissHitsNetwork(t *testing.T) {
	requestCount := 0

	data, err := os.ReadFile("testdata/drivers_2024.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	cacheFile := t.TempDir() + "/cache.db"

	client := NewClientWithBaseURL(server.URL+"/", WithCache(cacheFile, 1*time.Hour))

	_, err = client.GetDrivers(WithSeason(2024))
	if err != nil {
		t.Fatalf("first request failed: %v", err)
	}
	if requestCount != 1 {
		t.Errorf("expected 1 request, got %d", requestCount)
	}
}
