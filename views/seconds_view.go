package views

import (
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type SecondsModel struct {
	duration                time.Duration
	Timer                   timer.Model
	SecondsFirstRunComplete bool
	Width                   int
	Height                  int
}

func NewSecondsModel(duration time.Duration, interval time.Duration) tea.Model {
	tm := timer.NewWithInterval(duration, interval)

	return SecondsModel{
		duration:                duration,
		Timer:                   tm,
		SecondsFirstRunComplete: false,
	}
}

func (s *SecondsModel) Start() tea.Cmd {
	if s.SecondsFirstRunComplete {
		return s.Reset()
	} else {
		s.SecondsFirstRunComplete = true
		return s.Timer.Start()
	}
}

func (s *SecondsModel) Reset() tea.Cmd {
	s.Timer = timer.NewWithInterval(20*time.Second, time.Second)
	return s.Timer.Init()
}

func (s SecondsModel) Init() tea.Cmd {
	return s.Timer.Init()
}

func (s SecondsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.Width = msg.Width
		s.Height = msg.Height
	}

	s.Timer, cmd = s.Timer.Update(msg)
	return s, cmd
}

func (s SecondsModel) View() string {
	return s.Timer.View()
}
