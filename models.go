package f1

type mrData struct {
	Series string `json:"series"`
	URL    string `json:"url"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Total  int    `json:"total"`
}

// type apiResponse[T any] struct {
// 	MRData struct {
// 		mrData
// 		DriverTable      *driverTable      `json:"DriverTable,omitempty"`
// 		RaceTable        *raceTable        `json:"RaceTable,omitempty"`
// 		SeasonTable      *seasonTable      `json:"SeasonTable,omitempty"`
// 		CircuitTable     *circuitTable     `json:"CircuitTable,omitempty"`
// 		ConstructorTable *constructorTable `json:"ConstructorTable,omitempty"`
// 		StandingsTable   *standingsTable   `json:"StandingsTable,omitempty"`
// 		StatusTable      *statusTable      `json:"StatusTable,omitempty"`
// 	} `json:"MRData"`
// }

type Driver struct {
	DriverID        string `json:"driverId"`
	PermanentNumber string `json:"permanentNumber"`
	Code            string `json:"code"`
	URL             string `json:"url"`
	GivenName       string `json:"givenName"`
	FamilyName      string `json:"familyName"`
	DateOfBirth     string `json:"dateOfBirth"`
	Nationality     string `json:"nationality"`
}

type driverTable struct {
	Season  string   `json:"season"`
	Drivers []Driver `json:"Drivers"`
}

type Location struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"long"`
	Locality  string `json:"locality"`
	Country   string `json:"country"`
}

type Circuit struct {
	CircuitID   string   `json:"circuitId"`
	URL         string   `json:"url"`
	CircuitName string   `json:"circuitName"`
	Location    Location `json:"Location"`
}

type Session struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

type Timing struct {
	DriverID string `json:"driverId"`
	Time     string `json:"time"`
	Position string `json:"position"`
}

type Lap struct {
	Number  string   `json:"number"`
	Timings []Timing `json:"Timings"`
}

type PitStop struct {
	DriverID string `json:"driverId"`
	Lap      string `json:"lap"`
	Stop     string `json:"stop"`
	Time     string `json:"time"`
	Duration string `json:"duration"`
}

type QualifyingResult struct {
	Number      string      `json:"number"`
	Position    string      `json:"position"`
	Driver      Driver      `json:"Driver"`
	Constructor Constructor `json:"Constructor"`
	Q1          string      `json:"Q1,omitempty"`
	Q2          string      `json:"Q2,omitempty"`
	Q3          string      `json:"Q3,omitempty"`
}

type Result struct {
	Number       string      `json:"number"`
	Position     string      `json:"position"`
	PositionText string      `json:"positionText"`
	Points       string      `json:"points"`
	Driver       Driver      `json:"Driver"`
	Constructor  Constructor `json:"Constructor"`
	Grid         string      `json:"grid"`
	Laps         string      `json:"laps"`
	Status       string      `json:"status"`
	Time         struct {
		Millis string `json:"millis"`
		Time   string `json:"time"`
	} `json:"Time,omitempty"`
	FastestLap struct {
		Rank string `json:"rank"`
		Lap  string `json:"lap"`
		Time struct {
			Time string `json:"time"`
		} `json:"Time"`
	} `json:"FastestLap,omitempty"`
}

type Race struct {
	Season            string             `json:"season"`
	Round             string             `json:"round"`
	URL               string             `json:"url"`
	RaceName          string             `json:"raceName"`
	Circuit           Circuit            `json:"Circuit"`
	Date              string             `json:"date"`
	Time              string             `json:"time"`
	Results           []Result           `json:"Results,omitempty"`
	SprintResults     []Result           `json:"SprintResults,omitempty"`
	PitStops          []PitStop          `json:"PitStops,omitempty"`
	Laps              []Lap              `json:"Laps,omitempty"`
	FistPractice      Session            `json:"FirstPractice"`
	SecondPractice    Session            `json:"SecondPractice"`
	ThirdPractice     Session            `json:"ThirdPractice"`
	Qualifying        Session            `json:"Qualifying"`
	QualifyingResults []QualifyingResult `json:"QualifyingResults,omitempty"`
}

type raceTable struct {
	Season string `json:"season"`
	Round  string `json:"round,omitempty"`
	Races  []Race `json:"Races"`
}

type Season struct {
	Season string `json:"season"`
	URL    string `json:"url"`
}

type seasonTable struct {
	Seasons []Season `json:"Seasons"`
}

type circuitTable struct {
	Circuits []Circuit `json:"Circuits"`
}

type Constructor struct {
	ConstructorID string `json:"constructorId"`
	URL           string `json:"url"`
	Name          string `json:"name"`
	Nationality   string `json:"nationality"`
}

type constructorTable struct {
	Season       string        `json:"season"`
	Constructors []Constructor `json:"Constructors"`
}

type DriverStanding struct {
	Position     string        `json:"position"`
	PositionText string        `json:"positionText"`
	Points       string        `json:"points"`
	Wins         string        `json:"wins"`
	Driver       Driver        `json:"Driver"`
	Constructors []Constructor `json:"Constructor"`
}

type ConstructorStanding struct {
	Position     string      `json:"position"`
	PositionText string      `json:"positionText"`
	Points       string      `json:"points"`
	Wins         string      `json:"wins"`
	Constructor  Constructor `json:"Constructor"`
}

type StandingsList struct {
	Season               string                `json:"season"`
	Round                string                `json:"round"`
	DriverStandings      []DriverStanding      `json:"DriverStandings,omitempty"`
	ConstructorStandings []ConstructorStanding `json:"ConstructorStandings,omitempty"`
}

type standingsTable struct {
	Season         string          `json:"season"`
	Round          string          `json:"round"`
	StandingsLists []StandingsList `json:"StandingsLists"`
}

type Status struct {
	StatusID string `json:"statusId"`
	Count    string `json:"count"`
	Status   string `json:"status"`
}

type statusTable struct {
	Status []Status `json:"Status"`
}
