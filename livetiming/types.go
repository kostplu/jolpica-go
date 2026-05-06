package livetiming

import "time"

type YearsResponse struct {
	Years []Year `json:"Years"`
}

type Year struct {
	Year int    `json:"Year"`
	Path string `json:"Path"`
}

type MeetingsResponse struct {
	Year     int       `json:"Year"`
	Meetings []Meeting `json:"Meetings"`
}

type Meeting struct {
	Sessions     []Session `json:"Sessions"`
	Key          int       `json:"Key"`
	Code         string    `json:"Code"`
	Number       int       `json:"Number"`
	Location     string    `json:"Location"`
	OfficialName string    `json:"OfficialName"`
	Name         string    `json:"Name"`
	Country      Country   `json:"Country"`
	Circuit      Circuit   `json:"Circuit"`
}

type Session struct {
	Key       int    `json:"Key"`
	Type      string `json:"Type"`
	Number    int    `json:"Number"`
	Name      string `json:"Name"`
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
	GmtOffset string `json:"GmtOffset"`
	Path      string `json:"Path"`
}

type Country struct {
	Key  int    `json:"Key"`
	Code string `json:"Code"`
	Name string `json:"Name"`
}

type Circuit struct {
	Key       int    `json:"Key"`
	ShortName string `json:"ShortName"`
}

type FeedsResponse struct {
	Feeds Feeds `json:"Feeds"`
}

type Feeds struct {
	SessionInfo            Feed `json:"SessionInfo"`
	ArchiveStatus          Feed `json:"ArchiveStatus"`
	TrackStatus            Feed `json:"TrackStatus"`
	SessionData            Feed `json:"SessionData"`
	ContentStreams         Feed `json:"ContentStreams"`
	AudioStreams           Feed `json:"AudioStreams"`
	ExtrapolatedClock      Feed `json:"ExtrapolatedClock"`
	ChampionshipPrediction Feed `json:"ChampionshipPrediction"`
	TyreStintSeries        Feed `json:"TyreStintSeries"`
	SessionStatus          Feed `json:"SessionStatus"`
	TimingDataF1           Feed `json:"TimingDataF1"`
	TimingData             Feed `json:"TimingData"`
	DriverList             Feed `json:"DriverList"`
	LapSeries              Feed `json:"LapSeries"`
	TopThree               Feed `json:"TopThree"`
	TimingAppData          Feed `json:"TimingAppData"`
	TimingStats            Feed `json:"TimingStats"`
	Heartbeat              Feed `json:"Heartbeat"`
	WeatherData            Feed `json:"WeatherData"`
	WeatherDataSeries      Feed `json:"WeatherDataSeries"`
	PositionZ              Feed `json:"Position.z"`
	LapCount               Feed `json:"LapCount"`
	DriverRaceInfo         Feed `json:"DriverRaceInfo"`
	CarDataZ               Feed `json:"CarData.z"`
	TlaRcm                 Feed `json:"TlaRcm"`
	RaceControlMessages    Feed `json:"RaceControlMessages"`
	TeamRadio              Feed `json:"TeamRadio"`
	CurrentTyres           Feed `json:"CurrentTyres"`
	PitLaneTimeCollection  Feed `json:"PitLaneTimeCollection"`
	PitStop                Feed `json:"PitStop"`
	PitStopSeries          Feed `json:"PitStopSeries"`
	OvertakeSeries         Feed `json:"OvertakeSeries"`
	DriverTracker          Feed `json:"DriverTracker"`
}

type Feed struct {
	KeyFramePath string `json:"KeyFramePath"`
	StreamPath   string `json:"StreamPath"`
}

type SessionInfo struct {
	Meeting Meeting `json:"Meeting"`
	Session
}

type ArchiveStatus struct {
	Status string `json:"Status"`
}

type TrackStatus struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
}

type SessionData struct {
	Series       []SeriesEntry       `json:"Series"`
	StatusSeries []StatusSeriesEntry `json:"StatusSeries"`
}

type SeriesEntry struct {
	Utc time.Time `json:"Utc"`
	Lap int       `json:"Lap"`
}

type StatusSeriesEntry struct {
	Utc         time.Time `json:"Utc"`
	TrackStatus string    `json:"TrackStatus"`
}

type ContentStreams struct {
	Streams []Stream `json:"Streams"`
}

type Stream struct {
	Type     string `json:"Type"`
	Name     string `json:"Name"`
	Language string `json:"Language"`
	URI      string `json:"Uri"`
	Utc      string `json:"Utc"`
}

type AudioStreams struct {
	Streams []Stream `json:"Streams"`
}

type ExtrapolatedClock struct {
	Utc           time.Time `json:"Utc"`
	Remaining     string    `json:"Remaining"`
	Extrapolating bool      `json:"Extrapolating"`
}

type ChampionshipPrediction struct {
	Drivers map[string]DriverPrediction `json:"Drivers"`
	Teams   map[string]TeamPrediction   `json:"Teams"`
}

type DriverPrediction struct {
	RacingNumber      string `json:"RacingNumber"`
	CurrentPosition   int    `json:"CurrentPosition"`
	PredictedPosition int    `json:"PredictedPosition"`
	CurrentPoints     int    `json:"CurrentPoints"`
	PredictedPoints   int    `json:"PredictedPoints"`
}

type TeamPrediction struct {
	TeamName          string `json:"TeamName"`
	CurrentPosition   int    `json:"CurrentPosition"`
	PredictedPosition int    `json:"PredictedPosition"`
	CurrentPoints     int    `json:"CurrentPoints"`
	PredictedPoints   int    `json:"PredictedPoints"`
}

type TyreStintSeries struct {
	Stints map[string][]TyreEntry `json:"Stints"`
}

type TyreEntry struct {
	Compound        string `json:"Compound"`
	New             string `json:"New"`
	TyresNotChanged string `json:"TyresNotChanged"`
	TotalLaps       int    `json:"TotalLaps"`
	StartLaps       int    `json:"StartLaps"`
}

type SessionStatus struct {
	Status string `json:"Status"`
}

type TimingDataF1 struct {
	Lines    map[string]TimingDataF1DriverLine `json:"Lines"`
	Withheld bool                              `json:"Withheld"`
}

type TimingDataF1DriverLine struct {
	GapToLeader             string `json:"GapToLeader"`
	IntervalToPositionAhead struct {
		Value    string `json:"Value"`
		Catching bool   `json:"Catching"`
	} `json:"IntervalToPositionAhead"`
	Line             int                   `json:"Line"`
	Position         string                `json:"Position"`
	ShowPosition     bool                  `json:"ShowPosition"`
	RacingNumber     string                `json:"RacingNumber"`
	Retired          bool                  `json:"Retired"`
	InPit            bool                  `json:"InPit"`
	PitOut           bool                  `json:"PitOut"`
	Stopped          bool                  `json:"Stopped"`
	Status           int                   `json:"Status"`
	NumberOfLaps     int                   `json:"NumberOfLaps"`
	NumberOfPitstops int                   `json:"NumberOfPitstops"`
	Sectors          []SectorEntry         `json:"Sectors"`
	Speeds           map[string]SpeedEntry `json:"Speeds"`
	BestLapTime      struct {
		Value string `json:"Value"`
		Lap   int    `json:"Lap"`
	} `json:"BestLapTime"`
	LastLapTime struct {
		Value           string `json:"Value"`
		Status          int    `json:"Status"`
		OverallFastest  bool   `json:"OverallFastest"`
		PersonalFastest bool   `json:"PersonalFastest"`
	} `json:"LastLapTime"`
}

type SectorEntry struct {
	Stopped       bool   `json:"Stopped"`
	PreviousValue string `json:"PreviousValue"`
	Segments      []struct {
		Status int `json:"Status"`
	} `json:"Segments"`
	Value           string `json:"Value"`
	Status          int    `json:"Status"`
	OverallFastest  bool   `json:"OverallFastest"`
	PersonalFastest bool   `json:"PersonalFastest"`
}

type SpeedEntry struct {
	Value           string `json:"Value"`
	Status          int    `json:"Status"`
	OverallFastest  bool   `json:"OverallFastest"`
	PersonalFastest bool   `json:"PersonalFastest"`
}

type TimingData struct {
	Lines    map[string]TimingDataDriverLine `json:"Lines"`
	Withheld bool                            `json:"Withheld"`
}

type TimingDataDriverLine struct {
	TimingDataF1DriverLine
}

type DriverList map[string]DriverEntry

type DriverEntry struct {
	RacingNumber  string `json:"RacingNumber"`
	BroadcastName string `json:"BroadcastName"`
	FullName      string `json:"FullName"`
	Tla           string `json:"Tla"`
	Line          int    `json:"Line"`
	TeamName      string `json:"TeamName"`
	TeamColour    string `json:"TeamColour"`
	FirstName     string `json:"FirstName"`
	LastName      string `json:"LastName"`
	Reference     string `json:"Reference"`
	NameFormat    string `json:"NameFormat"`
	HeadshotURL   string `json:"HeadshotUrl"`
	CountryCode   string `json:"CountryCode"`
}

type LapCount struct {
	CurrentLap int `json:"CurrentLap"`
	TotalLaps  int `json:"TotalLaps"`
}

type DriverRaceInfo map[string]DriverRaceInfoEntry

type DriverRaceInfoEntry struct {
	RacingNumber  string `json:"RacingNumber"`
	Position      string `json:"Position"`
	Gap           string `json:"Gap"`
	Interval      string `json:"Interval"`
	PitStops      int    `json:"PitStops"`
	Catching      int    `json:"Catching"`
	OvertakeState int    `json:"OvertakeState"`
	IsOut         bool   `json:"IsOut"`
}

type CurrentTyres struct {
	Tyres map[string]TyreEntry `json:"Tyres"`
}

type TimingAppData struct {
	Lines map[string]TimingAppDataEntry `json:"Lines"`
}

type TimingAppDataEntry struct {
	RacingNumber string       `json:"RacingNumber"`
	Line         int          `json:"Line"`
	GridPos      string       `json:"GridPos"`
	Stints       []StintEntry `json:"Stints"`
}

type StintEntry struct {
	LapTime         string `json:"LapTime"`
	LapNumber       int    `json:"LapNumber"`
	LapFlags        int    `json:"LapFlags"`
	Compound        string `json:"Compound"`
	New             string `json:"New"`
	TyresNotChanged string `json:"TyresNotChanged"`
	TotalLaps       int    `json:"TotalLaps"`
	StartLaps       int    `json:"StartLaps"`
}

type DriverTracker struct {
	Withheld bool                 `json:"Withheld"`
	Lines    []DriverTrackerEntry `json:"Lines"`
}

type DriverTrackerEntry struct {
	Position        string `json:"Position"`
	ShowPosition    bool   `json:"ShowPosition"`
	RacingNumber    string `json:"RacingNumber"`
	LapTime         string `json:"LapTime"`
	LapState        int    `json:"LapState"`
	DiffToAhead     string `json:"DiffToAhead"`
	DiffToLeader    string `json:"DiffToLeader"`
	OverallFastest  bool   `json:"OverallFastest"`
	PersonalFastest bool   `json:"PersonalFastest"`
}

type LapSeries map[string]LapSeriesEntry

type LapSeriesEntry struct {
	RacingNumber string   `json:"RacingNumber"`
	LapPosition  []string `json:"LapPosition"`
}

type TopThree struct {
	Withheld bool            `json:"Withheld"`
	Lines    []TopThreeEntry `json:"Lines"`
}

type TopThreeEntry struct {
	Position        string `json:"Position"`
	ShowPosition    bool   `json:"ShowPosition"`
	RacingNumber    string `json:"RacingNumber"`
	Tla             string `json:"Tla"`
	BroadcastName   string `json:"BroadcastName"`
	FullName        string `json:"FullName"`
	FirstName       string `json:"FirstName"`
	LastName        string `json:"LastName"`
	Reference       string `json:"Reference"`
	Team            string `json:"Team"`
	TeamColour      string `json:"TeamColour"`
	LapTime         string `json:"LapTime"`
	LapState        int    `json:"LapState"`
	DiffToAhead     string `json:"DiffToAhead"`
	DiffToLeader    string `json:"DiffToLeader"`
	OverallFastest  bool   `json:"OverallFastest"`
	PersonalFastest bool   `json:"PersonalFastest"`
}

type TimingStats struct {
	Withheld bool                        `json:"Withheld"`
	Lines    map[string]TimingStatsEntry `json:"Lines"`
}

type TimingStatsEntry struct {
	Line               int    `json:"Line"`
	RacingNumber       string `json:"RacingNumber"`
	PeronalBestLapTime struct {
		Lap      int    `json:"Lap"`
		Position int    `json:"Position"`
		Value    string `json:"Value"`
	} `json:"PersonalBestLapTime"`
	BestSectors []struct {
		Position int    `json:"Position"`
		Value    string `json:"Value"`
	} `json:"BestSectors"`
	BestSpeeds map[string]struct {
		Position int    `json:"Position"`
		Value    string `json:"Value"`
	} `json:"BestSpeeds"`
}

type WeatherData struct {
	AirTemp       string `json:"AirTemp"`
	Humidity      string `json:"Humidity"`
	Pressure      string `json:"Pressure"`
	Rainfall      string `json:"Rainfall"`
	TrackTemp     string `json:"TrackTemp"`
	WindDirection string `json:"WindDirection"`
	WindSpeed     string `json:"WindSpeed"`
}

type WeatherDataSeries struct {
	Series []struct {
		Timestamp time.Time   `json:"Timestamp"`
		Weather   WeatherData `json:"Weather"`
	}
}

type Heartbeat struct {
	Utc time.Time `json:"Utc"`
}

type TlaRcm struct {
	Timestamp time.Time `json:"Timestamp"`
	Message   string    `json:"Message"`
}

type RaceControlMessages struct {
	Messages []RaceControlMessageEntry `json:"Messages"`
}

type RaceControlMessageEntry struct {
	Utc      time.Time `json:"Utc"`
	Lap      int       `json:"Lap"`
	Category string    `json:"Category"`
	Flag     string    `json:"Flag"`
	Scope    string    `json:"Scope"`
	Message  string    `json:"Message"`
}

type OvertakeSeries struct {
	Overtakes map[string][]OvertakeEntry `json:"Overtakes"`
}

type OvertakeEntry struct {
	Timestamp time.Time `json:"Timestamp"`
	Count     int       `json:"Count"`
}

type PitLaneTimeCollection struct {
	PitTimes struct{} `json:"PitTimes"`
}

type PitStop struct {
	RacingNumber string `json:"RacingNumber"`
	PitStopTime  string `json:"PitStopTime"`
	PitLaneTime  string `json:"PitLaneTime"`
	Lap          string `json:"Lap"`
}

type PitStopSeries struct {
	PitTimes map[string][]struct {
		Timestamp time.Time `json:"Timestamp"`
		PitStop   PitStop   `json:"PitStop"`
	} `json:"PitTimes"`
}

type PositionZ struct {
	Position []PositionZEntry `json:"Position"`
}

type PositionZEntry struct {
	Timestamp time.Time `json:"Timestamp"`
	Entries   map[string]struct {
		Status string `json:"Status"`
		X      int    `json:"X"`
		Y      int    `json:"Y"`
		Z      int    `json:"Z"`
	} `json:"Entries"`
}

type CarDataZ struct {
	Entries []CarDataZEntries `json:"Entries"`
}

type CarDataZEntries struct {
	Utc  time.Time `json:"Utc"`
	Cars map[string]struct {
		Channels map[string]int `json:"Channels"`
	} `json:"Cars"`
}

type TeamRadio struct {
	Captures []TeamRadioCapture `json:"Captures"`
}

type TeamRadioCapture struct {
	Utc          time.Time `json:"Utc"`
	RacingNumber string    `json:"RacingNumber"`
	Path         string    `json:"Path"`
}

type StreamEntry[T any] struct {
	Timestamp time.Duration
	Data      T
}
