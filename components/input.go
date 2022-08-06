package components

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/style"
	"time"
)

type (
	// Input the input component.
	Input struct {
		Model   textinput.Model
		program *tea.Program

		/* option start */
		Status Status

		Prompt        string
		Placeholder   string
		BlinkSpeed    time.Duration
		EchoMode      EchoMode
		EchoCharacter rune

		PromptStyle      lipgloss.Style
		TextStyle        lipgloss.Style
		BackgroundStyle  lipgloss.Style
		PlaceholderStyle lipgloss.Style
		CursorStyle      lipgloss.Style

		// default is disable
		QuitKey key.Binding
		// CharLimit is the maximum amount of characters this input element will
		// accept. If 0 or less, there's no limit.
		CharLimit int
		/* option end */
	}
)

// NewInput constructor
func NewInput() *Input {
	c := &Input{
		Model:            textinput.New(),
		Status:           Focus,
		Prompt:           "> ",
		Placeholder:      "",
		BlinkSpeed:       DefaultBlinkSpeed,
		EchoMode:         EchoNormal,
		EchoCharacter:    '*',
		PlaceholderStyle: style.New().Foreground(color.Gray),
		CharLimit:        0,
		QuitKey:          key.NewBinding(),
	}

	return c
}

// Focus sets the Focus state on the model. When the model is in Focus it can
// receive keyboard input and the cursor will be hidden.
func (in *Input) Focus() {
	in.program.Send(Focus)
}

// Blur removes the Focus state on the model.  When the model is blurred it can
// not receive keyboard input and the cursor will be hidden.
func (in *Input) Blur() {
	in.program.Send(Blur)
}

// Quit Input
func (in *Input) Quit() {
	in.program.Send(Quit)
}

// Value returns the value of the text input.
func (in *Input) Value() string {
	return in.Model.Value()
}

// Cursor returns the cursor position.
func (in *Input) Cursor() int {
	return in.Model.Cursor()
}

// Blink returns whether or not to draw the cursor.
func (in *Input) Blink() bool {
	return in.Model.Blink()
}

// SetCursor moves the cursor to the given position. If the position is
// out of bounds the cursor will be moved to the start or end accordingly.
func (in *Input) SetCursor(pos int) {
	in.Model.SetCursor(pos)
}

// Focused returns the focus state on the model.
func (in *Input) Focused() bool {
	return in.Model.Focused()
}

// CursorStart moves the cursor to the start of the input field.
func (in *Input) CursorStart() {
	in.Model.CursorStart()
}

// CursorEnd moves the cursor to the end of the input field.
func (in *Input) CursorEnd() {
	in.Model.CursorEnd()
}

// Reset sets the input to its default state with no input. Returns whether
// or not the cursor blink should reset.
func (in *Input) Reset() bool {
	return in.Model.Reset()
}

// CursorMode returns the model's cursor mode. For available cursor modes, see
// type CursorMode.
func (in *Input) CursorMode() CursorMode {
	return newCursorMode(in.Model.CursorMode())
}

// SetCursorMode sets the model's cursor mode. This method returns a command.
//
// For available cursor modes, see type CursorMode.
func (in *Input) SetCursorMode(model CursorMode) {
	in.Model.SetCursorMode(model.Map())
}

func (in *Input) Init() tea.Cmd {

	in.Model.Prompt = in.Prompt
	in.Model.Placeholder = in.Placeholder
	in.Model.BlinkSpeed = in.BlinkSpeed
	in.Model.EchoMode = textinput.EchoMode(in.EchoMode)
	in.Model.EchoCharacter = in.EchoCharacter
	in.Model.PromptStyle = in.PromptStyle
	in.Model.TextStyle = in.TextStyle
	in.Model.BackgroundStyle = in.BackgroundStyle
	in.Model.PlaceholderStyle = in.PlaceholderStyle
	in.Model.CursorStyle = in.CursorStyle
	in.Model.CharLimit = in.CharLimit

	return tea.Batch(textinput.Blink, func() tea.Msg {
		return in.Status
	})
}

func (in *Input) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, in.QuitKey):
			in.Model.Blur()
			return in, tea.Quit
		}
	case Status:
		in.Status = msg
		switch msg {
		case Focus:
			cmds = append(cmds, in.Model.Focus())
		case Blur:
			in.Model.Blur()
		case Quit:
			in.Model.Blur()
			return in, tea.Quit
		}
	}

	model, modelCmd := in.Model.Update(msg)
	in.Model = model
	cmds = append(cmds, modelCmd)

	return in, tea.Batch(cmds...)
}

func (in *Input) View() string {
	return in.Model.View()
}

func (in *Input) SetProgram(program *tea.Program) {
	in.program = program
}
