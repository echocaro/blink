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

	t.Run("Changes MinutesFirstRunComplete to true", func(t *testing.T) {
		minutesModel := model.(views.MinutesModel)
		minutesModel.StartTimer(duration, interval)

		if minutesModel.MinutesFirstRunComplete != true {
			t.Errorf("Expected MinutesFirstRunComplete to be true, got %v", minutesModel.MinutesFirstRunComplete)
		}
	})
	t.Run("Resets timer if MinutesFirstRunComplete is true", func(t *testing.T) {
		minutesModel := model.(views.MinutesModel)
		minutesModel.MinutesFirstRunComplete = true

		minutesModel.StartTimer(duration, interval)

		// the Reset func should be triggered on the third run
		// which means the id of the Timer should be 3
		if minutesModel.Timer.ID() != 3 {
			t.Errorf("Expected ID to be 3, got %v", minutesModel.Timer.ID())
		}
	})
}
