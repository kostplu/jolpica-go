package f1

import (
	"testing"
	"time"
)

func TestGetQualifying_ReturnsParsedQualifying(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/qualifying_2024_1.json"))

	page, err := client.GetQualifyingResults(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].QualifyingResults) != 20 {
		t.Errorf("expected 20 qualifying results, got %d", len(page.Races[0].QualifyingResults))
	}

firstRace := page.Races[0]
firstQualifying := firstRace.QualifyingResults[0]
if firstRace.Round != 1 {
		t.Errorf("expected first race to have round 1, got %d", firstRace.Round)
	}
	if firstRace.Circuit.CircuitName != "Bahrain International Circuit" {
		t.Errorf("expected first race to be at Bahrain International Circuit, got '%s'", firstRace.Circuit.CircuitName)
	}
	if firstRace.Circuit.Location.Country != "Bahrain" {
		t.Errorf("expected first race to be in Bahrain, got '%s'", firstRace.Circuit.Location.Country)
	}
	if firstQualifying.Driver.DriverID != "max_verstappen" {
		t.Errorf("expected first qualifying result to be for driver 'max_verstappen', got '%s'", firstQualifying.Driver.DriverID)
	}
	if firstQualifying.Constructor.ConstructorID != "red_bull" {
		t.Errorf("expected first qualifying result to be for constructor 'red_bull', got '%s'", firstQualifying.Constructor.ConstructorID)
	}
	if firstQualifying.Q1 != (LapTime{Duration: time.Duration(60)*time.Minute + time.Duration(30031*time.Millisecond)}) {
		t.Errorf("expected first qualifying Q1 time to be '1:30.031', got '%s'", firstQualifying.Q1)
	}
	if page.PageInfo.Limit != 30 {
		t.Errorf("expected page limit to be 30, got %d", page.PageInfo.Limit)
	}
	if page.PageInfo.Offset != 0 {
		t.Errorf("expected page offset to be 0, got %d", page.PageInfo.Offset)
	}
	if page.PageInfo.Total != 20 {
		t.Errorf("expected total qualifying results to be 20, got %d", page.PageInfo.Total)
	}
}

func TestGetQualifying_EmptySeasonReturnsEmptyList(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/qualifying_2024_1_empty.json"))

	page, err := client.GetQualifyingResults(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].QualifyingResults) != 0 {
		t.Errorf("expected 0 qualifying results, got %d", len(page.Races[0].QualifyingResults))
	}
}

func TestGetQualifying_FaultyResponseReturnsNoError(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/qualifying_2024_1_with_faults.json"))

	page, err := client.GetQualifyingResults(WithSeason(-1), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].QualifyingResults) != 1 {
		t.Errorf("expected 1 qualifying result, got %d", len(page.Races[0].QualifyingResults))
	}
	
	firstRace := page.Races[0]
	firstQualifying := firstRace.QualifyingResults[0]
	if firstRace.Circuit.CircuitID != "" {
			t.Errorf("expected first race to have empty circuit ID, got '%s'", firstRace.Circuit.CircuitID)
	}
	if firstQualifying.Driver.GivenName != "" {
		t.Errorf("expected first qualifying result to have empty driver given name, got '%s'", firstQualifying.Driver.GivenName)
	}
	if firstQualifying.Q1 != (LapTime{}) {
		t.Errorf("expected first qualifying Q1 time to be empty, got '%s'", firstQualifying.Q1)
	}
}

func TestIntegration_GetQualifying(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()

	page, err := client.GetQualifyingResults(WithSeason(2024), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 1 {
		t.Errorf("expected 1 race, got %d", len(page.Races))
	}
	if len(page.Races[0].QualifyingResults) == 0 {
		t.Errorf("expected at least 1 qualifying result, got %d", len(page.Races[0].QualifyingResults))
	}
}

func TestIntegration_GetQualifying_InvalidSeasonReturnsNoError(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()

	page, err := client.GetQualifyingResults(WithSeason(9999), WithRound(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Races) != 0 {
		t.Errorf("expected 0 races for invalid season, got %d", len(page.Races))
	}
}

func TestIntegration_GetAllQualifying(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()

	races, err := client.GetAllQualifyingResults(WithSeason(2024))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(races) == 0 {
		t.Errorf("expected at least 1 race, got %d", len(races))
	}
}

func TestIntegration_GetAllQualifying_InvalidSeasonReturnsEmptyList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	races, err := client.GetAllQualifyingResults(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(races) != 0 {
		t.Errorf("expected 0 races for invalid season, got %d", len(races))
	}
}