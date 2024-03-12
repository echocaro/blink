package views_test

import (
	"blink/views"
	"testing"
	"time"
)

func TestNewMinutesModel(t *testing.T) {
	t.Run("Initiates a timer with correct values", func(t *testing.T) {
		duration := 20 * time.Minute
		interval := 1 * time.Second

		model := views.NewMinutesModel(duration, interval)

		minutesModel := model.(views.MinutesModel)

		if minutesModel.Timer.Timeout != duration {
			t.Errorf("expected duration %v, got %v", duration, minutesModel.Timer.Timeout)
		}

		if minutesModel.Timer.Interval != interval {
			t.Errorf("expected timer duration %v, got %v", interval, minutesModel.Timer.Interval)
		}
	})
}

func TestMinutesStartTimer(t *testing.T) {
	duration := 5 * time.Second
	interval := 1 * time.Second

	model := views.NewMinutesModel(duration, interval)
	minutesModel := model.(views.MinutesModel)

	t.Run("Changes MinutesFirstRunComplete to true", func(t *testing.T) {
		minutesModel.StartTimer(duration, interval)

		if minutesModel.MinutesFirstRunComplete != true {
			t.Errorf("Expected MinutesFirstRunComplete to be true, got %v", minutesModel.MinutesFirstRunComplete)
		}
	})
}
