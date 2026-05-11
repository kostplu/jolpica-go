package livetiming

import (
	"bufio"
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func PrettyLogString(str string) {
	jsonData := []byte(str)
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, jsonData, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println(prettyJSON.String())
}

func PrettyLog(data any) {
	smth, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	PrettyLogString(string(smth))
}

func Str2Duration(s string) (time.Duration, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("unexpected lap time format: %s", s)
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("error parsing hours: %w", err)
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("error parsing minutes: %w", err)
	}

	seconds, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing seconds: %w", err)
	}

	total := time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute +
		time.Duration(seconds*float64(time.Second))

	return total, nil
}

func DecodeBase64AndDecompress(data []byte) (string, error) {
	// remove the first and last characters
	cleanData := strings.ReplaceAll(string(data), "\r", "")
	cleanData = strings.ReplaceAll(cleanData, "\n", "")
	if len(cleanData) < 2 {
		return "", fmt.Errorf("data is too short to remove first and last characters")
	}
	cleanData = cleanData[1 : len(cleanData)-1]
	decoded, err := base64.StdEncoding.DecodeString(cleanData)
	if err != nil {
		return "", fmt.Errorf("error decoding base64: %w", err)
	}

	r := flate.NewReader(bytes.NewBuffer(decoded))
	defer r.Close()

	decompressed, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("error decompressing data: %w", err)
	}
	return string(decompressed), nil
}

func ParseFeed[T any](data []byte) ([]StreamEntry[T], error) {
	var feed []StreamEntry[T]
	scanner := bufio.NewScanner(bytes.NewReader(data))
	f, err := os.Create("feed.json")
	if err != nil {
		return nil, fmt.Errorf("error creating file: %w", err)
	}
	defer f.Close()

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		idx := strings.IndexAny(line, `"{`)
		if idx == -1 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}
		timestamp, err := Str2Duration(line[:idx])
		if err != nil {
			return nil, fmt.Errorf("error parsing timestamp: %w", err)
		}
		raw := line[idx:]
		var jsonData string
		if raw[0] == '"' {
			jsonData, err = DecodeBase64AndDecompress([]byte(raw))
			if err != nil {
				return nil, fmt.Errorf("error decoding and decompressing data: %w", err)
			}
		} else {
			jsonData = raw
		}
		f.WriteString(jsonData + "\n")
		var result T
		err = json.Unmarshal([]byte(jsonData), &result)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
		}
		feed = append(feed, StreamEntry[T]{
			Timestamp: timestamp,
			Data:      result,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading data: %w", err)
	}
	return feed, nil
}

func StreamFeed[T any](entries []StreamEntry[T]) <-chan StreamEntry[T] {
	ch := make(chan StreamEntry[T], 100)

	go func() {
		defer close(ch)
		for _, entry := range entries {
			ch <- entry
		}
	}()

	return ch
}

func ReplayFeed[T any](in <-chan StreamEntry[T], config ReplayConfig) <-chan T {
	out := make(chan T, 10)

	go func() {
		defer close(out)

		currentTime := config.StartTime
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		var pending *StreamEntry[T]

		for {
			<-ticker.C
			currentTime += time.Duration(float64(100*time.Millisecond) * config.Speed)
			if pending == nil {
				entry, ok := <-in
				if !ok {
					return
				}
				pending = &entry
			}
			for pending != nil && pending.Timestamp <= currentTime {
				out <- pending.Data
				pending = nil

				entry, ok := <-in
				if !ok {
					return
				}
				pending = &entry
			}
		}
	}()

	return out
}
