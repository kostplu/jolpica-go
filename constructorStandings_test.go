package f1

import "testing"

func TestGetConstructorStandings_ReturnsParsedConstructorStandings(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/constructorstandings_2024.json"))

	page, err := client.GetConstructorStandings()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.StandingsLists) == 0 {
		t.Error("expected at least 1 standings list, got 0")
	}
	if len(page.StandingsLists[0].ConstructorStandings) == 0 {
		t.Error("expected at least 1 constructor standing, got 0")
	}

	first := page.StandingsLists[0].ConstructorStandings[0]
	if first.Constructor.ConstructorID != "mclaren" {
		t.Errorf("expected first constructor to have ConstructorID 'mclaren', got '%s'", first.Constructor.ConstructorID)
	}
	if first.Constructor.Name != "McLaren" {
		t.Errorf("expected first constructor to have Name 'McLaren', got '%s'", first.Constructor.Name)
	}
	if first.Wins != 6 {
		t.Errorf("expected first constructor to have 6 wins, got %d", first.Wins)
	}
	if first.Points != 666 {
		t.Errorf("expected first constructor to have 666 points, got %f", first.Points)
	}
	if page.PageInfo.Limit != 30 {
		t.Errorf("expected page limit to be 30, got %d", page.PageInfo.Limit)
	}
	if page.PageInfo.Offset != 0 {
		t.Errorf("expected page offset to be 0, got %d", page.PageInfo.Offset)
	}
	if page.PageInfo.Total != 10 {
		t.Errorf("expected total constructor standings to be 10, got %d", page.PageInfo.Total)
	}
	if page.PageInfo.HasNext() {
		t.Errorf("expected HasNext to be false, got true")
	}
	if page.PageInfo.NextOffset() != 30 {
		t.Errorf("expected NextOffset to be 30, got %d", page.PageInfo.NextOffset())
	}
}

func TestGetConstructorStandings_EmptySeasonReturnsEmptyList(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/constructorstandings_2024_empty.json"))

	page, err := client.GetConstructorStandings()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.StandingsLists) == 0 {
		t.Error("expected at least 1 standings list, got 0")
	}
	if len(page.StandingsLists[0].ConstructorStandings) != 0 {
		t.Errorf("expected 0 constructor standings, got %d", len(page.StandingsLists[0].ConstructorStandings))
	}
}

func TestGetConstructorStandings_FaultyResponseReturnsNoError(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/constructorstandings_2024_with_faults.json"))

	page, err := client.GetConstructorStandings()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	first := page.StandingsLists[0].ConstructorStandings[0]
	if first.Constructor.ConstructorID != "" {
		t.Errorf("expected first constructor to have empty ConstructorID, got '%s'", first.Constructor.ConstructorID)
	}
	if first.Wins != 0 {
		t.Errorf("expected first constructor to have 0 wins, got %d", first.Wins)
	}
	if first.Points != 0 {
		t.Errorf("expected first constructor to have 0 points, got %f", first.Points)
	}
}

func TestIntegration_GetConstructorStandings(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetConstructorStandings(WithSeason(2024))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.StandingsLists) == 0 {
		t.Error("expected at least 1 standings list, got 0")
	}
	if len(page.StandingsLists[0].ConstructorStandings) == 0 {
		t.Error("expected at least 1 constructor standing, got 0")
	}
}

func TestIntegration_GetConstructorStandingsReturnsNoErrorOnInvalidSeason(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetConstructorStandings(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.StandingsLists) != 0 {
		t.Errorf("expected 0 standings lists for invalid season, got %d", len(page.StandingsLists))
	}
}

func TestIntegration_GetAllConstructorStandings(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	constructorStandings, err := client.GetAllConstructorStandings(WithSeason(2024))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(constructorStandings) == 0 {
		t.Error("expected at least 1 constructor standing, got 0")
	}
}

func TestIntegration_GetAllConstructorStandingsReturnsNoErrorOnInvalidSeason(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	constructorStandings, err := client.GetAllConstructorStandings(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(constructorStandings) != 0 {
		t.Errorf("expected 0 constructor standings for invalid season, got %d", len(constructorStandings))
	}
}
