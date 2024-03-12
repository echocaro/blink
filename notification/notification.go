package notification

import (
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func Notify(title, message string) tea.Msg {
	cmd := exec.Command("osascript", "-e", `display notification "`+message+`" with title "`+title+`"`)
	return cmd.Run()
}
