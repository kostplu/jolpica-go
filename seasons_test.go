package f1

import "testing"

func TestGetSeasons_ReturnsParsedSeasons(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/seasons.json"))

	page, err := client.GetSeasons()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Seasons) != 30 {
		t.Errorf("expected 30 seasons, got %d", len(page.Seasons))
	}

	firstSeason := page.Seasons[0]
	if firstSeason.Season != 1950 {
		t.Errorf("expected first season to be 1950, got %d", firstSeason.Season)
	}
	if firstSeason.URL != "https://en.wikipedia.org/wiki/1950_Formula_One_season" {
		t.Errorf("expected first season URL to be '1950_Formula_One_season', got '%s'", firstSeason.URL)
	}
	if page.PageInfo.Limit != 30 {
		t.Errorf("expected page limit to be 30, got %d", page.PageInfo.Limit)
	}
	if page.PageInfo.Offset != 0 {
		t.Errorf("expected page offset to be 0, got %d", page.PageInfo.Offset)
	}
	if page.PageInfo.Total != 77 {
		t.Errorf("expected total seasons to be 77, got %d", page.PageInfo.Total)
	}
}

func TestGetSeasons_EmptyReturnsEmptyList(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/seasons_empty.json"))

	page, err := client.GetSeasons()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Seasons) != 0 {
		t.Errorf("expected 0 seasons, got %d", len(page.Seasons))
	}
}

func TestGetSeasons_FaultyResponseReturnsNoError(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/seasons_with_faults.json"))

	page, err := client.GetSeasons()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Seasons) != 1 {
		t.Errorf("expected 1 season, got %d", len(page.Seasons))
	}

	firstSeason := page.Seasons[0]
	if firstSeason.URL != "" {
		t.Errorf("expected first season URL to be empty, got '%s'", firstSeason.URL)
	}
}

func TestIntegration_GetSeasons(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetSeasons()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Seasons) == 0 {
		t.Errorf("expected at least 1 season, got %d", len(page.Seasons))
	}
}

func TestIntegration_GetAllSeasons(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	seasons, err := client.GetAllSeasons()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(seasons) == 0 {
		t.Errorf("expected seasons to be returned, got 0")
	}
}
