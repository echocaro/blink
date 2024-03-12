package sound

import (
	"log"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func Sound() tea.Msg {
	// log.Printf("We are here: %v", path)
	// return func() tea.Msg {
	cmd := exec.Command("afplay", "sound/mixkit-retro-game-notification-212.wav")
	// log.Printf("What is the err: %v", err)
	if err := cmd.Run(); err != nil {
		log.Printf("Error playing sound: %v", err)
	}

	return nil
	// }
}
