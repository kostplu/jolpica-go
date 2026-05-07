# livetiming

> ⚠️ **Work in progress** — this package is under active development and the API is not stable.

A Go library for accessing and replaying Formula 1 live timing data from the official F1 static livetiming API.

## What it does

- Fetches historical session data (years, meetings, sessions) from `livetiming.formula1.com`
- Parses `.jsonStream` and `.z` compressed feed files
- Provides a replay engine that simulates real-time data playback at configurable speeds
- Includes a terminal UI (TUI) for interactive session selection and live track visualization

## Package structure

livetiming/
├── client.go # HTTP client with SQLite caching
├── types.go # Feed types (PositionZ, DriverList, etc.)
├── static/ # Historical data fetching and parsing
│ ├── feed.go # ParseFeed, StreamFeed, ReplayFeed
│ └── ...
└── cmd/tui/ # BubbleTea terminal UI
├── main.go
├── model.go
└── commands.go

## TUI

The interactive TUI allows you to browse available seasons, meetings, and sessions, then replay position data in real time on a terminal-rendered track map.

```bash
go run ./livetiming/cmd/tui/
```

Navigation: arrow keys to move, `enter` to select, `q` to quit.

## Roadmap

- [x] Feed parsing (plain and `.z` compressed)
- [x] Replay engine with configurable speed
- [x] Session selection TUI
- [x] Terminal track map rendering (Unicode block characters)
- [ ] Live car positions on track
- [ ] Driver list overlay (3-letter codes, team colors)
- [ ] Timing data panel (gaps, lap times, sector times)
- [ ] Pit stop and race control event overlays
- [ ] Kitty/Sixel graphics backend for supported terminals
- [ ] Live timing mode (current session, not just historical replay)

## Data sources

Session data is fetched from the official F1 livetiming static API. Circuit layout coordinates are sourced from the [MultiViewer](https://multiviewer.app/) API, which also powers [FastF1](https://github.com/theOehrly/Fast-F1).

## Notes

- Responses are cached locally in SQLite to avoid hammering the API during development
- The replay engine supports variable speed playback (`Speed: 1.0` = real time, `Speed: 10.0` = 10x)
- Early stream entries often have `X:0 Y:0 Z:0` — this is normal, cars haven't left the garage yet
