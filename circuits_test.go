package f1

import "testing"

func TestGetCircuits_ReturnsParsedCircuits(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/circuits_firstpage.json"))

	page, err := client.GetCircuits()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Circuits) == 0 {
		t.Errorf("expected at least 1 circuit, got 0")
	}

	first := page.Circuits[0]
	if first.CircuitID != "adelaide" {
		t.Errorf("expected first circuit to have CircuitID 'adelaide', got '%s'", first.CircuitID)
	}
	if first.CircuitName != "Adelaide Street Circuit" {
		t.Errorf("expected first circuit to have CircuitName 'Adelaide Street Circuit', got '%s'", first.CircuitName)
	}
	if first.Location.Country != "Australia" {
		t.Errorf("expected first circuit to be in country 'Australia', got '%s'", first.Location.Country)
	}
	if page.PageInfo.Total != 78 {
		t.Errorf("expected total circuits to be 78, got %d", page.PageInfo.Total)
	}
	if page.PageInfo.Limit != 30 {
		t.Errorf("expected page limit to be 30, got %d", page.PageInfo.Limit)
	}
	if page.PageInfo.Offset != 0 {
		t.Errorf("expected page offset to be 0, got %d", page.PageInfo.Offset)
	}
	if !page.PageInfo.HasNext() {
		t.Errorf("expected HasNext to be true, got false")
	}
	if page.PageInfo.NextOffset() != 30 {
		t.Errorf("expected NextOffset to be 30, got %d", page.PageInfo.NextOffset())
	}
}

func TestGetCircuits_EmptySeasonReturnsEmptyList(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/circuits_empty.json"))

	page, err := client.GetCircuits()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Circuits) != 0 {
		t.Errorf("expected 0 circuits, got %d", len(page.Circuits))
	}
}

func TestGetCircuits_FaultyResponseReturnsNoError(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/circuits_with_faults.json"))

	page, err := client.GetCircuits()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	first := page.Circuits[0]

	if first.CircuitID != "" {
		t.Errorf("expected first circuit to have empty CircuitID, got '%s'", first.CircuitID)
	}
	if first.CircuitName != "" {
		t.Errorf("expected first circuit to have empty CircuitName, got '%s'", first.CircuitName)
	}
	if first.Location.Country != "" {
		t.Errorf("expected first circuit to haveempty Location.Country, got '%s'", first.Location.Country)
	}
}

func TestIntegration_GetCircuits(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetCircuits()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Circuits) == 0 {
		t.Errorf("expected at least 1 circuit, got 0")
	}
}

func TestIntegration_GetCircuitsReturnsNoErrorOnInvalidSeason(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetCircuits(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Circuits) != 0 {
		t.Errorf("expected 0 circuits for invalid season, got %d", len(page.Circuits))
	}
}

func TestIntegration_GetAllCircuits(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	circuits, err := client.GetAllCircuits()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(circuits) == 0 {
		t.Errorf("expected at least 1 circuit, got 0")
	}
}

func TestIntegration_GetAllCircuitsReturnsNoErrorOnInvalidSeason(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	circuits, err := client.GetAllCircuits(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(circuits) != 0 {
		t.Errorf("expected 0 circuits for invalid season, got %d", len(circuits))
	}
}
