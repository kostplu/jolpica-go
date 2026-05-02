package f1

import (
	"testing"
)

func TestGetDriverStandings_ReturnsParsedConstructorStandings(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/driverstandings_2024.json"))

	page, err := client.GetConstructorStandings()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.StandingsLists) == 0 {
		t.Error("expected at least 1 standings list, got 0")
	}
	if len(page.StandingsLists[0].DriverStandings) == 0 {
		t.Error("expected at least 1 driver standing, got 0")
	}

	first := page.StandingsLists[0].DriverStandings[0]
	if first.Driver.DriverID != "max_verstappen" {
		t.Errorf("expected first driver to have DriverID 'max_verstappen', got '%s'", first.Driver.DriverID)
	}
	if first.Driver.PermanentNumber != 3 {
		t.Errorf("expected first driver to have PermanentNumber 3, got '%d'", first.Driver.PermanentNumber)
	}
	if first.Constructors[0].ConstructorID != "red_bull" {
		t.Errorf("expected first driver to have ConstructorID 'red_bull', got '%s'", first.Constructors[0].ConstructorID)
	}
	if first.Wins != 9 {
		t.Errorf("expected first driver to have 9 wins, got %d", first.Wins)
	}
	if page.PageInfo.Limit != 30 {
		t.Errorf("expected page limit to be 30, got %d", page.PageInfo.Limit)
	}
	if page.PageInfo.Offset != 0 {
		t.Errorf("expected page offset to be 0, got %d", page.PageInfo.Offset)
	}
	if page.PageInfo.Total != 24 {
		t.Errorf("expected total driver standings to be 24, got %d", page.PageInfo.Total)
	}
	if page.PageInfo.HasNext() {
		t.Errorf("expected HasNext to be false, got true")
	}
	if page.PageInfo.NextOffset() != 30 {
		t.Errorf("expected NextOffset to be 30, got %d", page.PageInfo.NextOffset())
	}
}

func TestGetDriverStandings_EmptySeasonReturnsEmptyList(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/driverstandings_2024_empty.json"))

	page, err := client.GetDriverStandings()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.StandingsLists) == 0 {
		t.Error("expected at least 1 standings list, got 0")
	}
	if len(page.StandingsLists[0].DriverStandings) != 0 {
		t.Errorf("expected 0 driver standings, got %d", len(page.StandingsLists[0].DriverStandings))
	}
}

func TestGetDriverStandings_FaultyResponseReturnsNoError(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/driverstandings_2024_with_faults.json"))

	page, err := client.GetDriverStandings()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	first := page.StandingsLists[0].DriverStandings[0]
	if first.Wins != 0 {
		t.Errorf("expected first driver to have 0 wins, got %d", first.Wins)
	}
	if first.Points != 0 {
		t.Errorf("expected first driver to have 0 points, got %f", first.Points)
	}
	if first.Driver.Nationality != "" {
		t.Errorf("expected first driver to have empty Nationality, got '%s'", first.Driver.Nationality)
	}
	if !first.Driver.DateOfBirth.IsZero() {
		t.Errorf("expected first driver to have empty DateOfBirth, got '%s'", first.Driver.DateOfBirth)
	}
	if first.Constructors[0].Name != "" {
		t.Errorf("expected first constructor to have empty Name, got '%s'", first.Constructors[0].Name)
	}
}

func TestIntegration_GetDriverStandings(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetDriverStandings(WithSeason(2024))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.StandingsLists) == 0 {
		t.Error("expected at least 1 standings list, got 0")
	}
	if len(page.StandingsLists[0].DriverStandings) == 0 {
		t.Error("expected at least 1 driver standing, got 0")
	}
}

func TestIntegration_GetDriverStandingsReturnsNoErrorOnInvalidSeason(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetDriverStandings(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.StandingsLists) != 0 {
		t.Errorf("expected 0 standings lists for invalid season, got %d", len(page.StandingsLists))
	}
}

func TestIntegration_GetAllDriverStandings(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	
	client := NewClient()
	driverStandings, err := client.GetAllDriverStandings(WithSeason(2024))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(driverStandings) == 0 {
		t.Error("expected at least 1 driver standing, got 0")
	}
}

func TestIntegration_GetAllDriverStandingsReturnsNoErrorOnInvalidSeason(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	driverStandings, err := client.GetAllDriverStandings(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(driverStandings) != 0 {
		t.Errorf("expected 0 driver standings for invalid season, got %d", len(driverStandings))
	}
}