package f1

import (
	"encoding/json"
	"testing"
	"time"
)

func TestLapTime_ParseLapTime(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
	}{
		{`"1:23.456"`, 83456 * time.Millisecond},
		{`"0:58.123"`, 58123 * time.Millisecond},
		{`""`, 0}, // empty
	}

	for _, tt := range tests {
		var lt LapTime
		if err := json.Unmarshal([]byte(tt.input), &lt); err != nil {
			t.Errorf("input %s: expected no error, got %v", tt.input, err)
		}
		if lt.Duration != tt.expected {
			t.Errorf("input %s: expected duration %v, got %v", tt.input, tt.expected, lt.Duration)
		}
	}
}

func TestDate_ParseDate(t *testing.T) {
	var d Date
	if err := json.Unmarshal([]byte(`"2024-07-01"`), &d); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if d.Year() != 2024 || d.Month() != time.July || d.Day() != 1 {
		t.Errorf("expected date 2024-07-01, got %s", d.Format("2006-01-02"))
	}
}

func TestIntString_ParseIntString(t *testing.T) {
	var i IntString
	if err := json.Unmarshal([]byte(`"42"`), &i); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if int(i) != 42 {
		t.Errorf("expected integer 42, got %d", i)
	}
}

func TestFloatString_ParseFloatString(t *testing.T) {
	var f FloatString
	if err := json.Unmarshal([]byte(`"3.14"`), &f); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if float64(f) != 3.14 {
		t.Errorf("expected float 3.14, got %f", f)
	}
}
