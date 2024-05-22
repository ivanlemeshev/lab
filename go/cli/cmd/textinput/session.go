package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionModel struct {
	textInput textinput.Model
	err       error
}

func initialSessionModel() sessionModel {
	ti := textinput.New()
	ti.Placeholder = "25"
	ti.Focus()
	ti.CharLimit = 2
	ti.Width = 2

	return sessionModel{
		textInput: ti,
		err:       nil, // ???
	}
}

func (m sessionModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m sessionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textInput.Value())
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m sessionModel) View() string {
	return fmt.Sprintf(
		"Set the session time:\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
