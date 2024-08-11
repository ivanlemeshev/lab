package app

import (
	"fmt"
	"notes/internal/note"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	createNoteView sessionState = iota + 1
)

type keyMap struct {
	Help    key.Binding
	Quit    key.Binding
	NewNote key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit, k.NewNote},
	}
}

var keys = keyMap{
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	NewNote: key.NewBinding(
		key.WithKeys("n", "n"),
		key.WithHelp("n", "create a new note"),
	),
}

type AppModel struct {
	state      sessionState
	name       string
	keys       keyMap
	help       help.Model
	createNote tea.Model
	texts      []tea.Cmd
	newState   bool
}

func New(name string) tea.Model {
	return AppModel{
		name:       name,
		keys:       keys,
		help:       help.New(),
		createNote: note.NewCreateNoteModel(),
		newState:   false,
	}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.NewNote):
			m.state = createNoteView
			m.newState = true
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	switch m.state {
	case createNoteView:
		// m.createNote = note.NewCreateNoteModel()

		if m.newState {
			msg = func(m tea.Model) tea.Msg {
				return note.NewNote{}
			}(m)
			m.newState = false
		}

		newModel, newCmd := m.createNote.Update(msg)
		noteModel, ok := newModel.(note.CreateNoteModel)
		if !ok {
			panic("unexpected model type")
		}
		m.createNote = noteModel
		cmd = newCmd
	default:

	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m AppModel) View() string {
	switch m.state {
	case createNoteView:
		return m.createNote.View()
	default:
		text := fmt.Sprintf("Hi. This %s application", m.name)
		helpView := m.help.View(m.keys)
		height := 8 - strings.Count(text, "\n") - strings.Count(helpView, "\n")
		return "\n" + text + strings.Repeat("\n", height) + helpView
	}
}
