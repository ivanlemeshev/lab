package main

import (
	"log"
	"notes/internal/app"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	a := app.New("Notes")
	p := tea.NewProgram(a, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
