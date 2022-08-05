package input

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Status About the state of the Component
type Status int

// Available Component Status.
const (
	Focus Status = iota
	Blur
	Quit
)

func FocusCmd() tea.Msg {
	return Focus
}

func BlurCmd() tea.Msg {
	return Blur
}

func QuitCmd() tea.Msg {
	return Quit
}

// CursorMode describes the behavior of the cursor.
type CursorMode int

// Available cursor modes.
const (
	CursorBlink CursorMode = iota
	CursorStatic
	CursorHide
)

// String returns a the cursor mode in a human-readable format. This method is
// provisional and for informational purposes only.
func (c CursorMode) String() string {
	return [...]string{
		"blink",
		"static",
		"hidden",
	}[c]
}

func (c CursorMode) Map() textinput.CursorMode {
	switch c {
	case CursorBlink:
		return textinput.CursorBlink
	case CursorStatic:
		return textinput.CursorStatic
	case CursorHide:
		return textinput.CursorHide
	}

	panic(fmt.Sprintf("unknow cursorMode :%d", c))
}

func newCursorMode(other textinput.CursorMode) CursorMode {
	switch other {
	case textinput.CursorBlink:
		return CursorBlink
	case textinput.CursorStatic:
		return CursorStatic
	case textinput.CursorHide:
		return CursorHide
	}

	panic(fmt.Sprintf("unknow cursorMode :%s", other))
}

// EchoMode sets the input behavior of the text input field.
type EchoMode int

const (
	// EchoNormal displays text as is. This is the default behavior.
	EchoNormal EchoMode = iota

	// EchoPassword displays the EchoCharacter mask instead of actual
	// characters.  This is commonly used for password fields.
	EchoPassword

	// EchoNone displays nothing as characters are entered. This is commonly
	// seen for password fields on the command line.
	EchoNone

	// EchoOnEdit.
)
