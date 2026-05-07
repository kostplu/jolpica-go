# jolpica-go

A Go client library for the [Jolpica F1 API](https://jolpi.ca) —
the open-source successor to the Ergast Formula 1 API.

![Tests](https://github.com/kostplu/jolpica-go/actions/workflows/test.yml/badge.svg)

## Installation

```bash
go get github.com/kostplu/jolpica-go
```

Requires Go 1.25 or later.

## Quick start

```go
package main

import (
    "fmt"
    f1 "github.com/kostplu/jolpica-go"
)

func main() {
    client := f1.NewClient()

    drivers, err := client.GetDrivers(f1.WithSeason(2024))
    if err != nil {
        panic(err)
    }

    for _, d := range drivers.Drivers {
        fmt.Printf("%s %s (%s)\n", d.GivenName, d.FamilyName, d.Code)
    }
}
```

## Live Timing

The [`livetiming`](./livetiming/README.md) package provides access to historical F1 session data with a replay engine and an interactive terminal UI. Browse seasons and races, then watch car positions replay on a terminal-rendered track map.

> Still under construction — see the [livetiming README](./livetiming/README.md) for the current roadmap.

## Usage

### Creating a client

```go
// default client
client := f1.NewClient()

// with caching — strongly recommended
// historical data never changes, caching avoids hammering Jolpica's servers
client := f1.NewClient(
    f1.WithCache("~/.cache/jolpica-go/cache.db", 24*time.Hour),
)

// with a custom timeout
client := f1.NewClient(
    f1.WithTimeout(15 * time.Second),
)
```

### Filtering

All methods accept functional options to filter results:

```go
// by season
client.GetDrivers(f1.WithSeason(2024))

// by season and constructor
client.GetDrivers(f1.WithSeason(2024), f1.WithConstructor("ferrari"))

// by season, round, and driver
client.GetResults(
    f1.WithSeason(2024),
    f1.WithRound(1),
    f1.WithDriver("hamilton"),
)
```

### Pagination

Each method returns a page with pagination metadata:

```go
page, err := client.GetDrivers(
    f1.WithSeason(2024),
    f1.WithLimit(10),
)

fmt.Println(page.PageInfo.Total)      // total results available
fmt.Println(page.PageInfo.HasMore())  // true if more pages exist

// fetch next page
next, err := client.GetDrivers(
    f1.WithSeason(2024),
    f1.WithLimit(10),
    f1.WithOffset(page.PageInfo.NextOffset()),
)
```

To fetch all pages automatically:

```go
drivers, err := client.GetAllDrivers(f1.WithSeason(2024))
```

### Endpoints

| Method                             | Description            |
| ---------------------------------- | ---------------------- |
| `GetDrivers(opts...)`              | Drivers                |
| `GetConstructors(opts...)`         | Constructors           |
| `GetRaces(opts...)`                | Race schedule          |
| `GetResults(opts...)`              | Race results           |
| `GetQualifying(opts...)`           | Qualifying results     |
| `GetDriverStandings(opts...)`      | Driver standings       |
| `GetConstructorStandings(opts...)` | Constructors standings |
| `GetCircuits(opts...)`             | Circuit information    |
| `GetLaps(opts...)`                 | Lap times              |
| `GetPitStops(opts...)`             | Pit stop data          |
| `GetSprintResults(opts...)`        | Sprint race results    |
| `GetStatus(opts...)`               | Status codes           |

Each method has a corresponding `GetAll` variant that handles pagination automatically.

### Typed values

All values are returned as proper Go types — no manual parsing needed:

```go
results, _ := client.GetResults(f1.WithSeason(2024), f1.WithRound(1))
race := results.Races[0]

// positions and points are ints and floats, not strings
winner := race.Results[0]
fmt.Println(int(winner.Position))        // 1
fmt.Println(float64(winner.Points))      // 25.0

// lap times are time.Duration
fmt.Println(winner.FastestLap.Time.Duration.Seconds()) // 83.456

// dates are time.Time
driver := winner.Driver
fmt.Println(driver.DateOfBirth.Year())   // 1985
```

### Examples

**Driver standings after a specific round:**

```go
standings, err := client.GetStandings(
    f1.WithSeason(2024),
    f1.WithRound(10),
)
for _, s := range standings.Standings[0].DriverStandings {
    fmt.Printf("P%d %s — %s pts\n",
        int(s.Position),
        s.Driver.Code,
        s.Points,
    )
}
```

**Qualifying results for a specific race:**

```go
page, err := client.GetQualifying(
    f1.WithSeason(2024),
    f1.WithRound(1),
)
race := page.Races[0]
for _, q := range race.QualifyingResults {
    fmt.Printf("P%d %s — Q1: %s\n",
        int(q.Position),
        q.Driver.Code,
        q.Q1.Duration,
    )
}
```

**All races Hamilton won:**

```go
results, err := client.GetAllResults(f1.WithDriver("hamilton"))
var wins []Race
for _, race := range results {
    if len(race.Results) > 0 && int(race.Results[0].Position) == 1 {
        wins = append(wins, race)
    }
}
fmt.Printf("Hamilton wins: %d\n", len(wins))
```

## Caching

jolpica-go uses SQLite for caching. Historical F1 data never changes,
so caching is strongly recommended both for performance and to be
respectful of Jolpica's free service.

```go
client := f1.NewClient(
    f1.WithCache("./f1-cache.db", 24*time.Hour),
)
```

Cache TTL guidance:

- Historical seasons (any completed season) — `24*time.Hour` or longer
- Current season race results — `1*time.Hour`
- Current season standings — `30*time.Minute`

## Rate limiting

Jolpica enforces a limit of 500 requests per hour. jolpica-go's built-in
caching reduces the number of requests made, and `GetAll` methods
include a small delay between pages to avoid bursting.
Please be a considerate user of Jolpica's free service.

## Terms of use

By using this library you are indirectly using the Jolpica F1 API.
Please review [Jolpica's terms of use](https://jolpi.ca) before
building applications with this library.

## License

MIT
