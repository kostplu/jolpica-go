package f1

import "testing"

func TestGetPitStops_ReturnsParsedPitStops(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/pitstops_2024_1_firstpage.json"))

	page, err := client.GetPitStops(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].PitStops) != 30 {
		t.Errorf("expected 30 pit stops, got %d", len(page.Races[0].PitStops))
	}

	firstRace := page.Races[0]
	firstPitStop := firstRace.PitStops[0]
	if firstRace.Round != 1 {
		t.Errorf("expected first race to have round 1, got %d", firstRace.Round)
	}
	if firstRace.Circuit.CircuitName != "Bahrain International Circuit" {
		t.Errorf("expected first race to be at Bahrain International Circuit, got '%s'", firstRace.Circuit.CircuitName)
	}
	if firstRace.Circuit.Location.Country != "Bahrain" {
		t.Errorf("expected first race to be in Bahrain, got '%s'", firstRace.Circuit.Location.Country)
	}
	if firstPitStop.DriverID != "hulkenberg" {
		t.Errorf("expected first pit stop to be for driver 'hulkenberg', got '%s'", firstPitStop.DriverID)
	}
	if firstPitStop.Time != "18:05:33" {
		t.Errorf("expected first pit stop time to be '18:05:33', got '%s'", firstPitStop.Time)
	}
	if firstPitStop.Duration != "36.604" {
		t.Errorf("expected first pit stop duration to be '36.604', got '%s'", firstPitStop.Duration)
	}
	if page.PageInfo.Limit != 30 {
		t.Errorf("expected page limit to be 30, got %d", page.PageInfo.Limit)
	}
	if page.PageInfo.Offset != 0 {
		t.Errorf("expected page offset to be 0, got %d", page.PageInfo.Offset)
	}
	if page.PageInfo.Total != 43 {
		t.Errorf("expected total pit stops to be 43, got %d", page.PageInfo.Total)
	}
}

func TestGetPitStops_EmptySeasonReturnsEmptyList(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/pitstops_2024_1_empty.json"))

	page, err := client.GetPitStops(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].PitStops) != 0 {
		t.Errorf("expected 0 pit stops, got %d", len(page.Races[0].PitStops))
	}
}

func TestGetPitStops_FaultyResponseReturnsNoError(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/pitstops_2024_1_with_faults.json"))

	page, err := client.GetPitStops(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].PitStops) != 1 {
		t.Errorf("expected 1 pit stop, got %d", len(page.Races[0].PitStops))
	}

	firstRace := page.Races[0]
	firstPitStop := firstRace.PitStops[0]
	if firstRace.RaceName != "" {
		t.Errorf("expected first race name to be empty, got '%s'", firstRace.RaceName)
	}
	if firstPitStop.DriverID != "" {
		t.Errorf("expected first pit stop driver ID to be empty, got '%s'", firstPitStop.DriverID)
	}
	if firstPitStop.Duration != "" {
		t.Errorf("expected first pit stop duration to be empty, got '%s'", firstPitStop.Duration)
	}
}

func TestIntegration_GetPitStops(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetPitStops(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races[0].PitStops) == 0 {
		t.Errorf("expected at least 1 pitstop, got %d", len(page.Races[0].PitStops))
	}
}

func TestIntegration_GetPitStops_InvalidSeasonReturnsNoError(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetPitStops(WithSeason(9999), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 0 {
		t.Errorf("expected 0 races for invalid season, got %d", len(page.Races))
	}
}

func TestIntegration_GetAllPitStops(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	races, err := client.GetAllPitStops(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(races) == 0 {
		t.Errorf("expected pit stops to be returned, got 0")
	}
}

func TestIntegration_GetAllPitStops_InvalidSeasonReturnsEmptyList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	races, err := client.GetAllPitStops(WithSeason(9999), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(races) != 0 {
		t.Errorf("expected 0 pit stops for invalid season, got %d", len(races))
	}
}
