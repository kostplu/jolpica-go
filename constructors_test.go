package f1

import "testing"

func TestGetConstructors_ReturnsParsedConstructors(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/constructors_2024.json"))

	page, err := client.GetConstructors(WithSeason(2024))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Constructors) == 0 {
		t.Error("expected at least 1 constructor, got 0")
	}

	first := page.Constructors[0]
	if first.ConstructorID != "alpine" {
		t.Errorf("expected first constructor to have ConstructorID 'alpine', got '%s'", first.ConstructorID)
	}
	if first.Name != "Alpine F1 Team" {
		t.Errorf("expected first constructor to have Name 'Alpine F1 Team', got '%s'", first.Name)
	}
	if len(page.Constructors) != 10 {
		t.Errorf("expected 10 constructors, got %d", len(page.Constructors))
	}
	if page.PageInfo.Limit != 30 {
		t.Errorf("expected page limit to be 30, got %d", page.PageInfo.Limit)
	}
	if page.PageInfo.Offset != 0 {
		t.Errorf("expected page offset to be 0, got %d", page.PageInfo.Offset)
	}
	if page.PageInfo.Total != 10 {
		t.Errorf("expected total constructors to be 10, got %d", page.PageInfo.Total)
	}
	if page.PageInfo.HasNext() {
		t.Errorf("expected HasNext to be false, got true")
	}
	if page.PageInfo.NextOffset() != 30 {
		t.Errorf("expected NextOffset to be 30, got %d", page.PageInfo.NextOffset())
	}
}

func TestGetConstructors_EmptySeasonReturnsEmptyList(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/constructors_2024_empty.json"))

	page, err := client.GetConstructors()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Constructors) != 0 {
		t.Errorf("expected 0 constructors, got %d", len(page.Constructors))
	}
}

func TestGetConstructors_FaultyResponseReturnsNoError(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/constructors_2024_with_faults.json"))

	page, err := client.GetConstructors()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	first := page.Constructors[0]
	if first.ConstructorID != "" {
		t.Errorf("expected first constructor to have empty ConstructorID, got '%s'", first.ConstructorID)
	}
	if first.Nationality != "" {
		t.Errorf("expected first constructor to have empty National")
	}
}

func TestIntegration_GetConstructors(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetConstructors(WithSeason(2024))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Constructors) == 0 {
		t.Error("expected at least 1 constructor, got 0")
	}
}

func TestIntegration_GetConstructorsReturnsNoErrorOnInvalidSeason(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetConstructors(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Constructors) != 0 {
		t.Errorf("expected 0 constructors for invalid season, got %d", len(page.Constructors))
	}
}

func TestIntegration_GetAllConstructors(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	constructors, err := client.GetAllConstructors()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(constructors) == 0 {
		t.Error("expected at least 1 constructor, got 0")
	}
}

func TestIntegration_GetAllConstructorsReturnsNoErrorOnInvalidSeason(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	constructors, err := client.GetAllConstructors(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(constructors) != 0 {
		t.Errorf("expected 0 constructors for invalid season, got %d", len(constructors))
	}
}
