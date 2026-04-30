package f1

import "testing"

func TestGetStatus_ReturnsParsedStatus(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/status_firstpage.json"))

	page, err := client.GetStatus()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Statuses) != 30 {
		t.Errorf("expected 30 status entries, got %d", len(page.Statuses))
	}

	first := page.Statuses[0]
	if first.StatusID != "1" {
		t.Errorf("expected first status ID to be '1', got '%s'", first.StatusID)
	}
	if first.Status != "Finished" {
		t.Errorf("expected first status to be 'Finished', got '%s'", first.Status)
	}
	if first.Count != 8035 {
		t.Errorf("expected first status Count to be 8035, got %d", first.Count)
	}
	if len(page.Statuses) != 30 {
		t.Errorf("expected 30 status entries, got %d", len(page.Statuses))
	}
	if page.PageInfo.Limit != 30 {
		t.Errorf("expected page limit to be 30, got %d", page.PageInfo.Limit)
	}
	if page.PageInfo.Offset != 0 {
		t.Errorf("expected page offset to be 0, got %d", page.PageInfo.Offset)
	}
	if page.PageInfo.Total != 136 {
		t.Errorf("expected total status entries to be 136, got %d", page.PageInfo.Total)
	}
	if !page.PageInfo.HasNext() {
		t.Errorf("expected HasNext to be true, got false")
	}
	if page.PageInfo.NextOffset() != 30 {
		t.Errorf("expected NextOffset to be 30, got %d", page.PageInfo.NextOffset())
	}
}

func TestGetStatus_EmptyResponseReturnsEmptyList(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/status_empty.json"))

	page, err := client.GetStatus()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Statuses) != 0 {
		t.Errorf("expected 0 status entries, got %d", len(page.Statuses))
	}
}

func TestGetStatus_FaultyResponseReturnsNoError(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/status_with_faults.json"))

	page, err := client.GetStatus()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	first := page.Statuses[0]

	if first.StatusID != "" {
		t.Errorf("expected first status to have empty ID, got '%s'", first.StatusID)
	}
	if first.Status != "" {
		t.Errorf("expected first status to have empty Status, got '%s'", first.Status)
	}
	if first.Count != 0 {
		t.Errorf("expected first status to have Count 0, got %d", first.Count)
	}
}

func TestIntegration_GetStatus(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetStatus()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Statuses) == 0 {
		t.Errorf("expected at least 1 status entry, got 0")
	}
}

func TestIntegration_GetStatusReturnsNoErrorOnInvalidSeason(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetStatus(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Statuses) != 0 {
		t.Errorf("expected 0 status entries for invalid season, got %d", len(page.Statuses))
	}
}

func TestIntegration_GetAllStatus(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	statuses, err := client.GetAllStatus()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(statuses) == 0 {
		t.Errorf("expected at least 1 status entry, got 0")
	}
}

func TestIntegration_GetAllStatusReturnsNoErrorOnInvalidSeason(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	statuses, err := client.GetAllStatus(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(statuses) != 0 {
		t.Errorf("expected 0 status entries for invalid season, got %d", len(statuses))
	}
}
