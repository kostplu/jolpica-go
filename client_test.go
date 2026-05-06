package f1

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDrivers_404ReturnsError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := NewClientWithBaseURL(server.URL + "/")
	_, err := client.GetDrivers(WithSeason(2024))
	if err == nil {
		t.Error("expected error for 404 response, got nil")
	}
}

func TestGetDrivers_InvalidJSONReturnsError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	client := NewClientWithBaseURL(server.URL + "/")
	_, err := client.GetDrivers(WithSeason(2024))
	if err == nil {
		t.Error("expected error for invalid JSON response, got nil")
	}
}

func TestGetDrivers_TimeoutReturnsError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {}
	}))
	defer server.Close()

	client := NewClientWithBaseURL(server.URL+"/", WithTimeout(100))
	_, err := client.GetDrivers(WithSeason(2024))
	if err == nil {
		t.Error("expected error for request timeout, got nil")
	}
}

func TestGetDrivers_500ReturnsError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewClientWithBaseURL(server.URL + "/")
	_, err := client.GetDrivers(WithSeason(2024))
	if err == nil {
		t.Error("expected error for 500 response, got nil")
	}
}
