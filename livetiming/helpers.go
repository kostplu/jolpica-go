package livetiming

import (
	"bufio"
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
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
