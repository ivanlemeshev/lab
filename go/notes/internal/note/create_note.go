package note

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type NewNote struct{}

type CreateNoteModel struct {
	textarea textarea.Model
	err      error
}

func NewCreateNoteModel() CreateNoteModel {
	ti := textarea.New()
	ti.Placeholder = "Once upon a time..."
	ti.ShowLineNumbers = false

	width, height, err := getTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		width = 80
		height = 24
	}
	fmt.Println("Terminal size:", width, "characters wide,", height, "characters tall")
	ti.SetWidth(width)
	ti.SetHeight(height - 10)

	ti.Focus()

	return CreateNoteModel{
		textarea: ti,
		err:      nil,
	}
}

func (m CreateNoteModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m CreateNoteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	case NewNote:
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m CreateNoteModel) View() string {
	return fmt.Sprintf(
		"Tell me a story.\n\n%s\n\n%s",
		m.textarea.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
}

func getTerminalSize() (int, int, error) {
	fd := int(os.Stdout.Fd())
	var dims struct {
		rows uint16
		cols uint16
		x    uint16
		y    uint16
	}

	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&dims)))
	if err != 0 {
		return 0, 0, err
	}

	return int(dims.cols), int(dims.rows), nil
}
