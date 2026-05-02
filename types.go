package f1

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type LapTime struct {
	time.Duration
}

func (l *LapTime) UnmarshalJSON(data []byte) error {
	// data arrives as `"1:23.456"` — with quotes
	s := strings.Trim(string(data), `"`)
	if s == "" || s == "\\N" || string(data) == "null" {
		return nil // null value from API
	}

	// split "1:23.456" into parts
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return fmt.Errorf("unexpected lap time format: %s", s)
	}

	minutes, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	seconds, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return err
	}

	total := time.Duration(minutes)*time.Minute +
		time.Duration(seconds*float64(time.Second))

	l.Duration = total
	return nil
}

type LapTimeObject struct {
	LapTime `json:"time"`
}

func (l *LapTimeObject) UnmarshalJSON(data []byte) error {
	var obj struct {
		Time LapTime `json:"time"`
	}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	l.LapTime = obj.Time
	return nil
}

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if s == "" || s == "\\N" || string(data) == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("invalid date format: %s", s)
	}

	d.Time = t
	return nil
}

type IntString int

func (i *IntString) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if s == "" || s == "\\N" || string(data) == "null" {
		return nil
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("invalid integer format: %s", s)
	}

	*i = IntString(n)
	return nil
}

type FloatString float64

func (f *FloatString) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if s == "" || s == "\\N" || string(data) == "null" {
		return nil
	}

	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("invalid float format: %s", s)
	}

	*f = FloatString(n)
	return nil
}
