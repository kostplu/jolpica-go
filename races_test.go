package f1

import (
	"testing"
)


func TestGetRaces_ReturnsParsedPitStops(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/races_2024.json"))

	page, err := client.GetRaces(WithSeason(2024))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) == 0 {
		t.Fatal("expected at least 1 race, got 0")
	}

	firstRace := page.Races[0]
	if firstRace.RaceName != "Bahrain Grand Prix" {
		t.Errorf("expected first race name to be 'Bahrain Grand Prix', got '%s'", firstRace.RaceName)
	}
	if firstRace.Circuit.CircuitID != "bahrain" {
		t.Errorf("expected first race circuit ID to be 'bahrain', got '%s'", firstRace.Circuit.CircuitID)
	}
	if firstRace.Date.Format("2006-01-02") != "2024-03-02" {
		t.Errorf("expected first race date to be '2024-03-02', got '%s'", firstRace.Date.Format("2006-01-02"))
	}
	if firstRace.FirstPractice.Date.Format("2006-01-02") != "2024-02-29" {
		t.Errorf("expected first practice date to be '2024-02-29', got '%s'", firstRace.FirstPractice.Date.Format("2006-01-02"))
	}
	if page.PageInfo.Limit != 30 {
		t.Errorf("expected page limit to be 30, got %d", page.PageInfo.Limit)
	}
	if page.PageInfo.Offset != 0 {
		t.Errorf("expected page offset to be 0, got %d", page.PageInfo.Offset)
	}
	if page.PageInfo.Total != 24 {
		t.Errorf("expected total races to be 24, got %d", page.PageInfo.Total)
	}
}

func TestGetRaces_EmptySeasonReturnsEmptyList(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/races_2024_empty.json"))

	page, err := client.GetRaces(WithSeason(2024))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 0 {
		t.Errorf("expected 0 races, got %d", len(page.Races))
	}
}

func TestGetRaces_FaultyResponseReturnsNoError(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/races_2024_with_faults.json"))

	page, err := client.GetRaces(WithSeason(2024))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) == 0 {
		t.Fatal("expected at least 1 race, got 0")
	}

	firstRace := page.Races[0]
	if firstRace.RaceName != "" {
		t.Errorf("expected first race name to be empty, got '%s'", firstRace.RaceName)
	}
	if firstRace.Circuit.Location.Country != "" {
		t.Errorf("expected first race circuit country to be empty, got '%s'", firstRace.Circuit.Location.Country)
	}
	if firstRace.FirstPractice.Date.Format("2006-01-02") != "0001-01-01" {
		t.Errorf("expected first practice date to be zero value, got '%s'", firstRace.FirstPractice.Date.Format("2006-01-02"))
	}
}

func TestIntegration_GetRaces(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetRaces(WithSeason(2024))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) == 0 {
		t.Fatal("expected at least 1 race, got 0")
	}
}

func TestIntegration_GetRaces_InvalidSeasonReturnsNoError(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetRaces(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 0 {
		t.Errorf("expected 0 races, got %d", len(page.Races))
	}
}

func TestIntegration_GetAllRaces(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	races, err := client.GetAllRaces()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(races) == 0 {
		t.Fatal("expected at least 1 race, got 0")
	}
}

func TestIntegration_GetAllRaces_InvalidSeasonReturnsNoError(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	races, err := client.GetAllRaces(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(races) != 0 {
		t.Errorf("expected 0 races for invalid season, got %d", len(races))
	}
}