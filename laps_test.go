package f1

import "testing"

func TestGetLaps_ReturnsParsedLaps(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/laps_2024_1_firstpage.json"))

	page, err := client.GetLaps(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].Laps) != 2 {
		t.Errorf("expected 2 laps, got %d", len(page.Races[0].Laps))
	}

	firstRace := page.Races[0]
	firstLap := firstRace.Laps[0]
	if firstRace.Round != 1 {
		t.Errorf("expected first race to have round 1, got %d", firstRace.Round)
	}
	if firstRace.Circuit.CircuitName != "Bahrain International Circuit" {
		t.Errorf("expected first race to be at Bahrain International Circuit, got '%s'", firstRace.Circuit.CircuitName)
	}
	if firstRace.Circuit.Location.Country != "Bahrain" {
		t.Errorf("expected first race to be in Bahrain, got '%s'", firstRace.Circuit.Location.Country)
	}
	if firstLap.Number != 1 {
		t.Errorf("expected first lap to have number 1, got %d", firstLap.Number)
	}
	if firstLap.Timings[0].DriverID != "max_verstappen" {
		t.Errorf("expected first timing to be for driver 'max_verstappen', got '%s'", firstLap.Timings[0].DriverID)
	}
	if page.PageInfo.Limit != 30 {
		t.Errorf("expected page limit to be 30, got %d", page.PageInfo.Limit)
	}
	if page.PageInfo.Offset != 0 {
		t.Errorf("expected page offset to be 0, got %d", page.PageInfo.Offset)
	}
	if page.PageInfo.Total != 1129 {
		t.Errorf("expected total laps to be 1129, got %d", page.PageInfo.Total)
	}
}

func TestGetLaps_EmptySeasonReturnsEmptyList(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/laps_2024_1_empty.json"))

	page, err := client.GetLaps(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].Laps) != 0 {
		t.Errorf("expected 0 laps, got %d", len(page.Races[0].Laps))
	}
}

func TestGetLaps_FaultyResponseReturnsNoError(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/laps_2024_1_with_faults.json"))

	page, err := client.GetLaps(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].Laps) != 1 {
		t.Errorf("expected 1 lap, got %d", len(page.Races[0].Laps))
	}

	firstRace := page.Races[0]
	firstLap := firstRace.Laps[0]

	if firstRace.Round != 1 {
		t.Errorf("expected first race to have round 1, got %d", firstRace.Round)
	}
	if firstRace.Circuit.CircuitName != "" {
		t.Errorf("expected CircuitName to be empty, got '%s'", firstRace.Circuit.CircuitName)
	}
	if firstRace.Circuit.Location.Country != "" {
		t.Errorf("expected Location.Country to be empty, got '%s'", firstRace.Circuit.Location.Country)
	}
	if firstLap.Number != 0 {
		t.Errorf("expected first lap to have number 0, got %d", firstLap.Number)
	}
	if firstLap.Timings[0].DriverID != "" {
		t.Errorf("expected first timing to have empty DriverID, got '%s'", firstLap.Timings[0].DriverID)
	}
}

func TestIntegration_GetLaps(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetLaps(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races[0].Laps) == 0 {
		t.Errorf("expected 2 laps, got %v", len(page.Races[0].Laps))
	}
}

func TestIntegration_GetLaps_InvalidSeasonReturnsEmptyList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetLaps(WithSeason(9999), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 0 {
		t.Errorf("expected 0 races for invalid season, got %d", len(page.Races))
	}
}

func TestIntegration_GetAllLaps(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	races, err := client.GetAllLaps(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(races) == 0 {
		t.Errorf("expected laps to be returned, got 0")
	}
}

func TestIntegration_GetAllLaps_InvalidSeasonReturnsEmptyList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	races, err := client.GetAllLaps(WithSeason(9999), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(races) != 0 {
		t.Errorf("expected 0 laps for invalid season, got %d", len(races))
	}
}
