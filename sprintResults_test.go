package f1

import "testing"

func TestGetSprintResults_ReturnsParsedSprintResults(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/sprintResults_2024_5.json"))

	page, err := client.GetSprintResults(WithSeason(2024), WithRound(5))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].SprintResults) != 20 {
		t.Errorf("expected 20 sprint results, got %d", len(page.Races[0].SprintResults))
	}

	firstRace := page.Races[0]
	firstSprintResult := firstRace.SprintResults[0]
	if firstRace.Round != 5 {
		t.Errorf("expected first race to have round 5, got %d", firstRace.Round)
	}
	if firstRace.Circuit.CircuitName != "Shanghai International Circuit" {
		t.Errorf("expected first race to be at Shanghai International Circuit, got '%s'", firstRace.Circuit.CircuitName)
	}
	if firstRace.Circuit.Location.Country != "China" {
		t.Errorf("expected first race to be in China, got '%s'", firstRace.Circuit.Location.Country)
	}
	if firstSprintResult.Driver.DriverID != "max_verstappen" {
		t.Errorf("expected first sprint result to be for driver 'max_verstappen', got '%s'", firstSprintResult.Driver.DriverID)
	}
	if firstSprintResult.Constructor.ConstructorID != "red_bull" {
		t.Errorf("expected first sprint result to be for constructor 'red_bull', got '%s'", firstSprintResult.Constructor.ConstructorID)
	}
	if firstSprintResult.Position != 1 {
		t.Errorf("expected first sprint result to have position 1, got %d", firstSprintResult.Position)
	}
	if firstSprintResult.Points != 8 {
		t.Errorf("expected first sprint result to have 8 points, got %f", firstSprintResult.Points)
	}
	if firstSprintResult.Status != "Finished" {
		t.Errorf("expected first sprint result to have status 'Finished', got '%s'", firstSprintResult.Status)
	}
	if page.PageInfo.Limit != 30 {
		t.Errorf("expected page limit to be 30, got %d", page.PageInfo.Limit)
	}
	if page.PageInfo.Offset != 0 {
		t.Errorf("expected page offset to be 0, got %d", page.PageInfo.Offset)
	}
	if page.PageInfo.Total != 20 {
		t.Errorf("expected total sprint results to be 20, got %d", page.PageInfo.Total)
	}
}

func TestGetSprintResults_EmptyRaceReturnsEmptyList(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/sprintResults_2024_5_empty.json"))

	page, err := client.GetSprintResults(WithSeason(2024), WithRound(5))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].SprintResults) != 0 {
		t.Errorf("expected 0 sprint results, got %d", len(page.Races[0].SprintResults))
	}
}

func TestGetSprintResults_FaultyResponseReturnsNoError(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/sprintResults_2024_5_with_faults.json"))

	page, err := client.GetSprintResults(WithSeason(2024), WithRound(5))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].SprintResults) != 1 {
		t.Errorf("expected 1 sprint result, got %d", len(page.Races[0].SprintResults))
	}

	firstRace := page.Races[0]
	firstSprintResult := firstRace.SprintResults[0]
	if firstRace.RaceName != "" {
		t.Errorf("expected first race name to be empty, got '%s'", firstRace.RaceName)
	}
	if firstSprintResult.Driver.GivenName != "" {
		t.Errorf("expected first sprint result driver given name to be empty, got '%s'", firstSprintResult.Driver.GivenName)
	}
	if firstSprintResult.Status != "" {
		t.Errorf("expected first sprint result status to be empty, got '%s'", firstSprintResult.Status)
	}
}

func TestIntegration_GetSprintResults(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetSprintResults(WithSeason(2024), WithRound(5))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races[0].SprintResults) == 0 {
		t.Errorf("expected at least 1 sprint result, got %d", len(page.Races[0].SprintResults))
	}
}

func TestIntegration_GetSprintResults_InvalidSeasonReturnsNoError(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetSprintResults(WithSeason(9999), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 0 {
		t.Errorf("expected 0 races for invalid season, got %d", len(page.Races))
	}
}

func TestIntegration_GetAllSprintResults(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	races, err := client.GetAllSprintResults(WithSeason(2024), WithRound(5))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(races) == 0 {
		t.Errorf("expected sprint results to be returned, got 0")
	}
}

func TestIntegration_GetAllSprintResults_InvalidSeasonReturnsEmptyList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	races, err := client.GetAllSprintResults(WithSeason(9999), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(races) != 0 {
		t.Errorf("expected 0 sprint results for invalid season, got %d", len(races))
	}
}
