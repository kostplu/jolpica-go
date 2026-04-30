package f1

import "testing"

func TestBuildURL(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		opts     []Option
		expected string
	}{
		{
			name:     "no options",
			endpoint: "drivers",
			opts:     []Option{},
			expected: "/drivers.json?limit=30&offset=0",
		},
		{
			name:     "season only",
			endpoint: "drivers",
			opts:     []Option{WithSeason(2024)},
			expected: "/2024/drivers.json?limit=30&offset=0",
		},
		{
			name:     "season and constructor",
			endpoint: "drivers",
			opts:     []Option{WithSeason(2024), WithConstructor("mercedes")},
			expected: "/2024/constructors/mercedes/drivers.json?limit=30&offset=0",
		},
		{
			name:     "season and driver",
			endpoint: "drivers",
			opts:     []Option{WithSeason(2024), WithDriver("hamilton")},
			expected: "/2024/drivers/hamilton/drivers.json?limit=30&offset=0",
		},
		{
			name:     "season, round and driver",
			endpoint: "results",
			opts:     []Option{WithSeason(2024), WithRound(5), WithDriver("hamilton")},
			expected: "/2024/5/drivers/hamilton/results.json?limit=30&offset=0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildPath(tt.endpoint, tt.opts)
			if got != tt.expected {
				t.Errorf("expected URL path '%s', got '%s'", tt.expected, got)
			}
		})
	}
}
