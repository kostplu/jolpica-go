package f1

import (
	"testing"
)

func TestGetDrivers_ReturnsParsedDrivers(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/drivers_2024.json"))

	page, err := client.GetDrivers()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Drivers) == 0 {
		t.Error("expected at least one driver, got 0")
	}

	first := page.Drivers[0]
	if first.DriverID != "albon" {
		t.Errorf("expected first driver ID to be 'albon', got '%s'", first.DriverID)
	}
	if first.DateOfBirth.Year() != 1996 {
		t.Errorf("expected first driver to be born in 1996, got %d", first.DateOfBirth.Year())
	}
	if first.DateOfBirth.Month() != 3 {
		t.Errorf("expected first driver to be born in March, got %d", first.DateOfBirth.Month())
	}
	if first.DateOfBirth.Day() != 23 {
		t.Errorf("expected first driver to be born on the 23rd, got %d", first.DateOfBirth.Day())
	}
	if first.DriverID != "albon" {
		t.Errorf("expected first driver ID to be 'albon', got '%s'", first.DriverID)
	}
	if first.Nationality != "Thai" {
		t.Errorf("expected first driver Nationality to be 'Thai', got '%s'", first.Nationality)
	}
	if first.PermanentNumber != 23 {
		t.Errorf("expected first driver PermanentNumber to be 23, got %d", first.PermanentNumber)
	}
	if len(page.Drivers) != 25 {
		t.Errorf("expected 20 drivers, got %d", len(page.Drivers))
	}
	if page.PageInfo.Limit != 30 {
		t.Errorf("expected page limit to be 30, got %d", page.PageInfo.Limit)
	}
	if page.PageInfo.Offset != 0 {
		t.Errorf("expected page offset to be 0, got %d", page.PageInfo.Offset)
	}
	if page.PageInfo.Total != 25 {
		t.Errorf("expected total drivers to be 25, got %d", page.PageInfo.Total)
	}
	if page.PageInfo.HasNext() {
		t.Errorf("expected HasNext to be false, got true")
	}
	if page.PageInfo.NextOffset() != 30 {
		t.Errorf("expected NextOffset to be 30, got %d", page.PageInfo.NextOffset())
	}
}

func TestGetDrivers_EmptySeasonReturnsEmptyList(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/drivers_2024_empty.json"))

	page, err := client.GetDrivers()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Drivers) != 0 {
		t.Errorf("expected 0 drivers, got %d", len(page.Drivers))
	}
}

func TestGetDrivers_FaultyResponseReturnsNoError(t *testing.T) {
	client := newTestClient(t, fixtureHandler(t, "testdata/drivers_2024_with_faults.json"))

	page, err := client.GetDrivers()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	first := page.Drivers[0]

	if first.FamilyName != "" {
		t.Errorf("expected first driver to have empty FamilyName, got '%s'", first.FamilyName)
	}
	if first.PermanentNumber != 0 {
		t.Errorf("expected second driver to have nil PermanentNumber, got %d", first.PermanentNumber)
	}
	if first.DriverID != "" {
		t.Errorf("expected third driver to have empty DriverID, got '%s'", first.DriverID)
	}
}

func TestIntegration_GetDriversReturnsNoErrorOnInvalidSeason(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetDrivers(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Drivers) != 0 {
		t.Errorf("expected 0 drivers for invalid season, got %d", len(page.Drivers))
	}
}

func TestIntegration_GetDrivers(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	page, err := client.GetDrivers(WithSeason(2024))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(page.Drivers) == 0 {
		t.Error("expected at least one driver, got 0")
	}
}

func TestIntegration_GetAllDrivers(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	drivers, err := client.GetAllDrivers(WithSeason(2024))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(drivers) == 0 {
		t.Error("expected at least one driver, got 0")
	}
}

func TestIntegration_GetAllDriversReturnsNoErrorOnInvalidSeason(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	client := NewClient()
	drivers, err := client.GetAllDrivers(WithSeason(9999))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(drivers) != 0 {
		t.Errorf("expected 0 drivers for invalid season, got %d", len(drivers))
	}
}
