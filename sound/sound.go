package sound

import (
	"log"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func Sound() tea.Msg {
	cmd := exec.Command("afplay", "sound/sound.mp3")
	if err := cmd.Run(); err != nil {
		log.Printf("Error playing sound: %v", err)
	}

	return nil

}
