package ui

import (
	"blink/form"
	"blink/notification"
	"blink/sound"
	"blink/views"
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type (
	AppState   int
	TimerState int
)

const (
	DisplayingForm AppState = iota
	formSubmitted
	TimerStarted
)

const (
	TimerStopped TimerState = iota
	CountingMinutes
	CountingSeconds
)

const (
	minutesDuration = 20 * time.Minute
	minutesInterval = time.Second
	secondsDuration = 20 * time.Second
)

type MainModel struct {
	form          form.FormModel
	secondsTimer  views.SecondsModel
	minutesTimer  views.MinutesModel
	TimerModel    timer.Model
	state         AppState
	timerState    TimerState
	formSubmitted bool
}

func InitialModel() MainModel {
	minutesModel := views.NewMinutesModel(minutesDuration, minutesInterval)
	secondsModel := views.NewSecondsModel(20 * time.Second)

	return MainModel{
		timerState:    TimerStopped,
		state:         DisplayingForm,
		form:          form.CreateForm(),
		minutesTimer:  minutesModel.(views.MinutesModel),
		secondsTimer:  secondsModel.(views.SecondsModel),
		formSubmitted: false,
	}
}

func (m MainModel) Init() tea.Cmd {
	return tea.Batch(m.form.Form.Init(), m.minutesTimer.Init(), m.secondsTimer.Init())
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
		switch msg.Type {
		case tea.KeyEnter:
			if m.state == DisplayingForm {
				m.formSubmitted = true

				updatedModel, _ := m.form.Form.Update(msg)

				updatedForm, ok := updatedModel.(*huh.Form)

				if !ok {
					return m, nil
				}
				m.form.Form = updatedForm

				m.formSubmitted = true
				log.Printf("Form submitted, Test value: %s\n", form.Confirm)

				if form.Confirm == "no" {
					return m, tea.Quit
				}
				if form.Confirm == "yes" {
					m.state = TimerStarted
					m.timerState = CountingMinutes
					return m, m.minutesTimer.StartTimer(minutesDuration, minutesInterval)
				}

			}
		}
	}

	switch m.state {
	case DisplayingForm:
		var cmd tea.Cmd

		updatedModel, cmd := m.form.Form.Update(msg)

		updatedForm, ok := updatedModel.(*huh.Form)

		if !ok {
			return m, nil
		}

		m.form.Form = updatedForm
		return m, cmd
	// case formSubmitted:
	// 	log.Printf("What is Confirm in this case: %+v\n\n", form.Confirm)
	// 	m.state = TimerStarted
	// 	m.timerState = CountingMinutes
	// 	return m, m.minutesTimer.StartTimer(minutesDuration, minutesInterval)

	case TimerStarted:
		switch m.timerState {
		case CountingMinutes:
			updatedModel, cmd := m.minutesTimer.Update(msg)
			m.minutesTimer = updatedModel.(views.MinutesModel)

			if m.minutesTimer.Timer.Timedout() && !m.minutesTimer.Timer.Running() {
				notification.Notify("Rest your eyes", "Look at something 20 feet away!")
				sound.Sound()
				m.timerState = CountingSeconds
				return m, m.secondsTimer.Start()
			}
			return m, cmd

		case CountingSeconds:
			updatedModel, cmd := m.secondsTimer.Update(msg)
			m.secondsTimer = updatedModel.(views.SecondsModel)
			if m.secondsTimer.Timer.Timedout() && !m.secondsTimer.Timer.Running() {
				notification.Notify("Time's up!", "You can look at the screen again!")
				sound.Sound()
				m.timerState = CountingMinutes
				return m, m.minutesTimer.StartTimer(minutesDuration, minutesInterval)
			}

			return m, cmd
		}
	}

	return m, cmd
}

func (m MainModel) View() string {
	switch m.state {
	case DisplayingForm:
		paragraphStyle := lipgloss.NewStyle().
			// Foreground(lipgloss.Color("#7471F9")).
			// Bold(true).
			Padding(1, 2, 1, 2).
			Width(70).
			Align(lipgloss.Left)

		timeStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F780E2")).
			Bold(true)

		paragraphText := `This program aims to reduce eye strain by reminding you to follow the 20-20-20 rule: every 20 minutes, glance at something 20 feet away for 20 seconds. Expect a notification and a sound alert to signal each interval's start and end.`

		styledParagraph := paragraphStyle.Render(paragraphText)

		styled20 := timeStyle.Render("20")
		styledParagraph = strings.ReplaceAll(styledParagraph, "20", styled20)

		return styledParagraph + "\n" + m.form.Form.View()
	case TimerStarted:
		switch m.timerState {
		case CountingMinutes:
			style := lipgloss.NewStyle().
				Bold(true).
				BorderForeground(lipgloss.Color("#F780E2")).
				BorderStyle(lipgloss.NormalBorder()).
				PaddingTop(2).
				PaddingBottom(2).
				Width(40).Align(lipgloss.Center)

			footer := lipgloss.NewStyle().
				Faint(true).
				Align(lipgloss.Center)

			return lipgloss.Place(
				m.minutesTimer.Width,
				m.minutesTimer.Height,
				lipgloss.Center,
				lipgloss.Center,
				lipgloss.JoinVertical(
					lipgloss.Center,
					style.Render("A short 👀 break is on the horizon: \n\n"+m.minutesTimer.View()),
					footer.Render("To exit, press ctrl+c or q"),
				),
			)

			// return m.minutesTimer.View()
		case CountingSeconds:
			style := lipgloss.NewStyle().
				Bold(true).
				BorderForeground(lipgloss.Color("#F780E2")).
				BorderStyle(lipgloss.NormalBorder()).
				// Background(lipgloss.Color("#7D56F4")).
				PaddingTop(2).
				PaddingBottom(2).
				PaddingLeft(4).
				Width(40)
			return lipgloss.Place(
				m.minutesTimer.Width,
				m.minutesTimer.Height,
				lipgloss.Center,
				lipgloss.Center,
				lipgloss.JoinVertical(
					lipgloss.Center,
					style.Render(m.secondsTimer.View()),
					"To exit, press ctrl+c",
				),
			)
		}
	}

	return "No views"
}
