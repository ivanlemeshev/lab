package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type breakModel struct {
	textInput textinput.Model
	err       error
}

func initialBreakModel() breakModel {
	ti := textinput.New()
	ti.Placeholder = "5"
	ti.Focus()
	ti.CharLimit = 2
	ti.Width = 2

	return breakModel{
		textInput: ti,
		err:       nil, // ???
	}
}

func (m breakModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m breakModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m breakModel) View() string {
	return fmt.Sprintf(
		"Set the break time:\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
