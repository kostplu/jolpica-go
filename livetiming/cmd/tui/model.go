package main

import (
	"fmt"
	"time"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/kostplu/jolpica-go/livetiming"
)

type state int

const (
	stateSelectYear state = iota
	stateSelectMeeting
	stateSelectSession
	stateReplaying
)

// -- list.Item implementations --

type yearItem struct {
	path, name string
	year       int
}

func (i yearItem) Title() string       { return i.name }
func (i yearItem) Description() string { return "" }
func (i yearItem) FilterValue() string { return i.name }

type meetingItem struct{ meeting livetiming.Meeting }

func (i meetingItem) Title() string       { return i.meeting.Name }
func (i meetingItem) Description() string { return fmt.Sprintf("Round %d", i.meeting.Number) }
func (i meetingItem) FilterValue() string { return i.meeting.Name }

type sessionItem struct{ session livetiming.Session }

func (i sessionItem) Title() string       { return i.session.Name }
func (i sessionItem) Description() string { return i.session.Type }
func (i sessionItem) FilterValue() string { return i.session.Name }

// -- messages --

type (
	yearsLoadedMsg    struct{ items []list.Item }
	meetingsLoadedMsg struct{ items []list.Item }
	sessionsLoadedMsg struct{ items []list.Item }
	mapLoaded         struct {
		trackX, trackY []int
		degrees        float64
		frames         <-chan livetiming.PositionZ
		carLocations   []livetiming.CarLocation
	}
	updateLocations struct {
		timestamp    time.Time
		carLocations []livetiming.CarLocation
	}
	replayDoneMsg struct{}
	tickMsg       struct{}
	errMsg        struct{ err error }
)

// -- model --

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	state               state
	client              *livetiming.Client
	list                list.Model
	data                string
	selectedYear        int
	selectedYearPath    string
	selectedSessionPath string
	selectedSession     livetiming.Session
	width               int
	height              int
	trackX              []int
	trackY              []int
	degrees             float64
	carLocations        []livetiming.CarLocation
	frames              <-chan livetiming.PositionZ
	timestamp           time.Time
}

func NewModel(client *livetiming.Client) model {
	l := list.New(nil, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select Year"
	return model{
		state:  stateSelectYear,
		client: client,
		list:   l,
	}
}

func (m model) Init() tea.Cmd {
	return fetchYears(m.client)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
		if msg.String() == "enter" && !m.list.SettingFilter() {
			switch m.state {
			case stateSelectYear:
				item, ok := m.list.SelectedItem().(yearItem)
				if !ok {
					break
				}
				m.selectedYear = item.year
				m.selectedYearPath = item.path
				m.list.Title = "Loading..."
				m.list.SetItems(nil)
				return m, fetchMeetings(m.client, item.path)

			case stateSelectMeeting:
				item, ok := m.list.SelectedItem().(meetingItem)
				if !ok {
					break
				}
				m.list.Title = "Loading..."
				m.list.SetItems(nil)
				return m, fetchSessions(m.client, item.meeting)

			case stateSelectSession:
				item, ok := m.list.SelectedItem().(sessionItem)
				if !ok {
					break
				}
				m.selectedSession = item.session
				m.selectedSessionPath = item.session.Path
				return m, startReplay(m.client, m.selectedSessionPath, item.session, m.selectedYear)
			}
		}

	case yearsLoadedMsg:
		m.list.Title = "Select Year"
		m.list.SetItems(msg.items)

	case meetingsLoadedMsg:
		m.state = stateSelectMeeting
		m.list.Title = "Select Meeting"
		m.list.SetItems(msg.items)

	case sessionsLoadedMsg:
		m.state = stateSelectSession
		m.list.Title = "Select Session"
		m.list.SetItems(msg.items)

	case mapLoaded:
		m.state = stateReplaying
		m.trackX = msg.trackX
		m.trackY = msg.trackY
		m.degrees = msg.degrees
		m.frames = msg.frames
		m.carLocations = msg.carLocations
		return m, waitForFrames(m.frames, m.carLocations)

	case updateLocations:
		m.timestamp = msg.timestamp
		m.carLocations = msg.carLocations
		return m, tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
			return tickMsg{}
		})

	case tickMsg:
		return m, waitForFrames(m.frames, m.carLocations)

	case replayDoneMsg:

		return m, tea.Quit
	case errMsg:
		// TODO: show error in UI
		return m, tea.Quit
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() tea.View {
	if m.state == stateReplaying {
		v := tea.NewView(renderTrack(m.trackX, m.trackY, m.width, m.height, m.degrees, m.carLocations, m.timestamp))
		v.AltScreen = true
		return v
	}
	v := tea.NewView(docStyle.Render(m.list.View()))
	v.AltScreen = true
	return v
}
