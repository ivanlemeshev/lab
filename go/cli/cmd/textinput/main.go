package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

// Set the session time (default 25 minutes)
// Set the break time (default 5 minutes)
// Set the long break time (default 15 minutes)
// Set the number of sessions (default 4)
// Set the number of sessions before the long break (default 4)

func main() {
	p := tea.NewProgram(initialSessionModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	p = tea.NewProgram(initialBreakModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
