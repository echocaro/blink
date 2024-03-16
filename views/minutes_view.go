package views

import (
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MinutesModel struct {
	duration                time.Duration
	Timer                   timer.Model
	MinutesFirstRunComplete bool
	Width                   int
	Height                  int
}

func NewMinutesModel(duration, interval time.Duration) tea.Model {
	tm := timer.NewWithInterval(duration, interval)

	return MinutesModel{
		duration:                duration,
		Timer:                   tm,
		MinutesFirstRunComplete: false,
	}
}

func (m MinutesModel) Init() tea.Cmd {
	return nil
}

func (m *MinutesModel) StartTimer(duration, interval time.Duration) tea.Cmd {
	if m.MinutesFirstRunComplete {
		return m.Reset(duration, interval)
	} else {
		m.MinutesFirstRunComplete = true
		return m.start()
	}
}

func (m *MinutesModel) start() tea.Cmd {
	return m.Timer.Start()
}

func (m *MinutesModel) Reset(duration, interval time.Duration) tea.Cmd {
	m.Timer = timer.NewWithInterval(duration, interval)
	return m.Timer.Start()
}

func (m MinutesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}
	m.Timer, cmd = m.Timer.Update(msg)
	return m, cmd
}

func (m MinutesModel) View() string {
	return lipgloss.Place(
		m.Width,
		m.Height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.Timer.View(),
		),
	)
	// return m.Timer.View()
}
