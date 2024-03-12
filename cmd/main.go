package main

import (
	"blink/ui"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.InitialModel())

	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			log.Fatalf("fatal: %v", err)
			os.Exit(1)
		}

		defer f.Close()
	}

	if _, err := p.Run(); err != nil {
		fmt.Printf("🙈 there's been an error: %v", err)
		os.Exit(1)
	}
}
