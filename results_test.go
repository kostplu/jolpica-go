package f1

import "testing"

func TestGetResults_ReturnsParsedResults(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/results_2024_1.json"))

	page, err := client.GetResults(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].Results) != 20 {
		t.Errorf("expected 20 results, got %d", len(page.Races[0].Results))
	}

	firstRace := page.Races[0]
	firstResult := firstRace.Results[0]
	if firstRace.Round != 1 {
		t.Errorf("expected first race to have round 1, got %d", firstRace.Round)
	}
	if firstRace.Circuit.CircuitName != "Bahrain International Circuit" {
		t.Errorf("expected first race to be at Bahrain International Circuit, got '%s'", firstRace.Circuit.CircuitName)
	}
	if firstRace.Circuit.Location.Country != "Bahrain" {
		t.Errorf("expected first race to be in Bahrain, got '%s'", firstRace.Circuit.Location.Country)
	}
	if firstResult.Driver.DriverID != "max_verstappen" {
		t.Errorf("expected first result to be for driver 'max_verstappen', got '%s'", firstResult.Driver.DriverID)
	}
	if firstResult.Constructor.ConstructorID != "red_bull" {
		t.Errorf("expected first result to be for constructor 'red_bull', got '%s'", firstResult.Constructor.ConstructorID)
	}
	if firstResult.Position != 1 {
		t.Errorf("expected first result to have position 1, got %d", firstResult.Position)
	}
	if firstResult.Points != 26 {
		t.Errorf("expected first result to have 26 points, got %f", firstResult.Points)
	}
	if firstResult.Status != "Finished" {
		t.Errorf("expected first result to have status 'Finished', got '%s'", firstResult.Status)
	}
	if page.PageInfo.Limit != 30 {
		t.Errorf("expected page limit to be 30, got %d", page.PageInfo.Limit)
	}
	if page.PageInfo.Offset != 0 {
		t.Errorf("expected page offset to be 0, got %d", page.PageInfo.Offset)
	}
	if page.PageInfo.Total != 20 {
		t.Errorf("expected total results to be 20, got %d", page.PageInfo.Total)
	}
}

func TestGetResults_EmptyRaceReturnsEmptyList(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/results_2024_1_empty.json"))

	page, err := client.GetResults(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].Results) != 0 {
		t.Errorf("expected 0 results, got %d", len(page.Races[0].Results))
	}
}

func TestGetResults_FaultyResponseReturnsNoError(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/results_2024_1_with_faults.json"))

	page, err := client.GetResults(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].Results) != 1 {
		t.Errorf("expected 1 result, got %d", len(page.Races[0].Results))
	}

	firstRace := page.Races[0]
	firstResult := firstRace.Results[0]
	if firstRace.RaceName != "" {
		t.Errorf("expected first race name to be empty, got '%s'", firstRace.RaceName)
	}
	if firstResult.Driver.GivenName != "" {
		t.Errorf("expected first result driver given name to be empty, got '%s'", firstResult.Driver.GivenName)
	}
	if firstResult.Status != "" {
		t.Errorf("expected first result status to be empty, got '%s'", firstResult.Status)
	}
}

func TestIntegration_GetResults(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetResults(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races[0].Results) == 0 {
		t.Errorf("expected at least 1 result, got %d", len(page.Races[0].Results))
	}
}

func TestIntegration_GetResults_InvalidSeasonReturnsNoError(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetResults(WithSeason(9999), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 0 {
		t.Errorf("expected 0 races for invalid season, got %d", len(page.Races))
	}
}

func TestIntegration_GetAllResults(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	races, err := client.GetAllResults(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(races) == 0 {
		t.Errorf("expected results to be returned, got 0")
	}
}

func TestIntegration_GetAllResults_InvalidSeasonReturnsEmptyList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	races, err := client.GetAllResults(WithSeason(9999), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(races) != 0 {
		t.Errorf("expected 0 results for invalid season, got %d", len(races))
	}
}
