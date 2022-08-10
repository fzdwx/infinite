package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type (
	/*
		Components, You can use these components directly:
			 	1. Input
			 	2. Selection
			 	3. Spinner
				4. Autocomplete
				5. Progress
		Or use them inline in your custom component,
		for how to embed them, you can refer to the implementation of `Confirm`.
	*/
	Components interface {
		tea.Model

		// SetProgram this method will be called back when the tea.Program starts.
		// please keep passing this method
		SetProgram(program *tea.Program)
	}
)

const (
	DefaultBlinkSpeed = time.Millisecond * 530
)

// Status About the state of the Component
type Status int

const (
	// Focus only use Input
	Focus Status = iota
	// Blur only use Input
	Blur
	// Quit component
	Quit
	// Normal ignore it
	Normal
)

// CursorMode describes the behavior of the cursor.
type CursorMode int

const (
	CursorBlink CursorMode = iota
	CursorStatic
	CursorHide
)

// String returns the cursor mode in a human-readable format. This method is
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

// EchoMode sets the Input behavior of the text Input field.
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
)
