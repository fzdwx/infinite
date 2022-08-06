package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/emoji"
	"time"
)

type (
	/*
		Components, You can use these components directly:
			 	1. Input
			 	2. Selection
			 	3. Spinner
		Or use them inline in your custom component,
		for how to embed them, you can refer to the implementation of `Confirm`.
	*/
	Components interface {
		tea.Model

		// SetProgram this method will be called back when the tea.Program starts.
		// please keep passing this method
		SetProgram(program *tea.Program)
	}

	StartUp struct {
		P       *tea.Program
		started bool
	}
)

// NewStartUp new StartUp
func NewStartUp(c Components, ops ...tea.ProgramOption) *StartUp {
	program := tea.NewProgram(c, ops...)

	c.SetProgram(program)

	return &StartUp{
		P:       program,
		started: false,
	}
}

func (s *StartUp) Start() error {
	s.started = true
	return s.P.Start()
}

// Kill Components
func (s *StartUp) Kill() {
	if s.started {
		s.started = false
		s.P.Kill()
	}
}

// Send message to component
func (s *StartUp) Send(msg tea.Msg) {
	if s.started {
		s.P.Send(msg)
	}
}

const (
	DefaultBlinkSpeed = time.Millisecond * 530
)

type (
	// Shape the Spinner Shape
	Shape struct {
		Frames []string
		FPS    time.Duration
	}
)

// Some spinners to choose from. You could also make your own.
var (
	Line = Shape{
		Frames: []string{"|", "/", "-", "\\"},
		FPS:    time.Second / 10, //nolint:gomnd
	}
	Dot = Shape{
		Frames: []string{"⣾ ", "⣽ ", "⣻ ", "⢿ ", "⡿ ", "⣟ ", "⣯ ", "⣷ "},
		FPS:    time.Second / 10, //nolint:gomnd
	}
	MiniDot = Shape{
		Frames: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		FPS:    time.Second / 12, //nolint:gomnd
	}
	Jump = Shape{
		Frames: []string{"⢄", "⢂", "⢁", "⡁", "⡈", "⡐", "⡠"},
		FPS:    time.Second / 10, //nolint:gomnd
	}
	Pulse = Shape{
		Frames: []string{"█", "▓", "▒", "░"},
		FPS:    time.Second / 8, //nolint:gomnd
	}
	Points = Shape{
		Frames: []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●"},
		FPS:    time.Second / 7, //nolint:gomnd
	}
	Globe = Shape{
		Frames: []string{"🌍", "🌎", "🌏"},
		FPS:    time.Second / 4, //nolint:gomnd
	}
	Moon = Shape{
		Frames: []string{"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘"},
		FPS:    time.Second / 8, //nolint:gomnd
	}
	Monkey = Shape{
		Frames: []string{"🙈", "🙉", "🙊"},
		FPS:    time.Second / 3, //nolint:gomnd
	}
	Meter = Shape{
		Frames: []string{
			"▱▱▱",
			"▰▱▱",
			"▰▰▱",
			"▰▰▰",
			"▰▰▱",
			"▰▱▱",
			"▱▱▱",
		},
		FPS: time.Second / 7, //nolint:gomnd
	}
	Hamburger = Shape{
		Frames: []string{"☱", "☲", "☴", "☲"},
		FPS:    time.Second / 3, //nolint:gomnd
	}
	Running = Shape{
		Frames: []string{emoji.Walking, emoji.Running},
		FPS:    time.Second / 6, //nolint:gomnd
	}
)

type (
	// Status About the state of the Component
	Status int

	// CursorMode describes the behavior of the cursor.
	CursorMode int

	// EchoMode sets the input behavior of the text input field.
	EchoMode int
)

const (
	// Focus only use Input
	Focus Status = iota
	// Blur only use Input
	Blur
	// Quit component
	Quit
	// Normal ignore it
	Normal

	CursorBlink CursorMode = iota
	CursorStatic
	CursorHide

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
