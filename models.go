package f1

type mrData struct {
	Series string `json:"series"`
	URL    string `json:"url"`
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	Total  string `json:"total"`
}

type apiResponse struct {
	MRData struct {
		mrData
		DriverTable      *driverTable      `json:"DriverTable,omitempty"`
		RaceTable        *raceTable        `json:"RaceTable,omitempty"`
		SeasonTable      *seasonTable      `json:"SeasonTable,omitempty"`
		CircuitTable     *circuitTable     `json:"CircuitTable,omitempty"`
		ConstructorTable *constructorTable `json:"ConstructorTable,omitempty"`
		StandingsTable   *standingsTable   `json:"StandingsTable,omitempty"`
		StatusTable      *statusTable      `json:"StatusTable,omitempty"`
	} `json:"MRData"`
}

type Driver struct {
	DriverID        string    `json:"driverId"`
	PermanentNumber IntString `json:"permanentNumber"`
	Code            string    `json:"code"`
	URL             string    `json:"url"`
	GivenName       string    `json:"givenName"`
	FamilyName      string    `json:"familyName"`
	DateOfBirth     Date      `json:"dateOfBirth"`
	Nationality     string    `json:"nationality"`
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
	Date Date   `json:"date"`
	Time string `json:"time"`
}

type Timing struct {
	DriverID string  `json:"driverId"`
	Time     LapTime `json:"time"`
	Position string  `json:"position"`
}

type Lap struct {
	Number  IntString `json:"number"`
	Timings []Timing  `json:"Timings"`
}

type PitStop struct {
	DriverID string    `json:"driverId"`
	Lap      IntString `json:"lap"`
	Stop     string    `json:"stop"`
	Time     string    `json:"time"`
	Duration string    `json:"duration"`
}

type QualifyingResult struct {
	Number      IntString   `json:"number"`
	Position    IntString   `json:"position"`
	Driver      Driver      `json:"Driver"`
	Constructor Constructor `json:"Constructor"`
	Q1          LapTime     `json:"Q1"`
	Q2          LapTime     `json:"Q2"`
	Q3          LapTime     `json:"Q3"`
}

type Result struct {
	Number       IntString   `json:"number"`
	Position     IntString   `json:"position"`
	PositionText string      `json:"positionText"`
	Points       FloatString `json:"points"`
	Driver       Driver      `json:"Driver"`
	Constructor  Constructor `json:"Constructor"`
	Grid         IntString   `json:"grid"`
	Laps         IntString   `json:"laps"`
	Status       string      `json:"status"`
	Time         struct {
		Millis string `json:"millis"`
		Time   string `json:"time"`
	} `json:"Time"`
	FastestLap FastestLap `json:"FastestLap"`
}

type LapTimeObject struct {
	LapTime `json:"time"`
}

type FastestLap struct {
	Rank IntString     `json:"rank"`
	Lap  IntString     `json:"lap"`
	Time LapTimeObject `json:"Time"`
}

type Race struct {
	Season            IntString          `json:"season"`
	Round             IntString          `json:"round"`
	URL               string             `json:"url"`
	RaceName          string             `json:"raceName"`
	Circuit           Circuit            `json:"Circuit"`
	Date              Date               `json:"date"`
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
	Season IntString `json:"season"`
	Round  IntString `json:"round,omitempty"`
	Races  []Race    `json:"Races"`
}

type Season struct {
	Season IntString `json:"season"`
	URL    string    `json:"url"`
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
	Season       IntString     `json:"season"`
	Constructors []Constructor `json:"Constructors"`
}

type DriverStanding struct {
	Position     IntString     `json:"position"`
	PositionText string        `json:"positionText"`
	Points       FloatString   `json:"points"`
	Wins         IntString     `json:"wins"`
	Driver       Driver        `json:"Driver"`
	Constructors []Constructor `json:"Constructors"`
}

type ConstructorStanding struct {
	Position     IntString   `json:"position"`
	PositionText string      `json:"positionText"`
	Points       FloatString `json:"points"`
	Wins         IntString   `json:"wins"`
	Constructor  Constructor `json:"Constructor"`
}

type StandingsList struct {
	Season               IntString             `json:"season"`
	Round                IntString             `json:"round"`
	DriverStandings      []DriverStanding      `json:"DriverStandings,omitempty"`
	ConstructorStandings []ConstructorStanding `json:"ConstructorStandings,omitempty"`
}

type standingsTable struct {
	Season         IntString       `json:"season"`
	Round          IntString       `json:"round"`
	StandingsLists []StandingsList `json:"StandingsLists"`
}

type Status struct {
	StatusID string    `json:"statusId"`
	Count    IntString `json:"count"`
	Status   string    `json:"status"`
}

type statusTable struct {
	Status []Status `json:"Status"`
}

type PageInfo struct {
	Total  int
	Offset int
	Limit  int
}

func (p PageInfo) HasNext() bool {
	return p.Offset+p.Limit < p.Total
}

func (p PageInfo) NextOffset() int {
	return p.Offset + p.Limit
}
