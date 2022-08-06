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
		Components

		Model textinput.Model

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

	c.Components = Components{Model: c}

	return c
}

// Focus sets the Focus state on the model. When the model is in Focus it can
// receive keyboard input and the cursor will be hidden.
func (c *Input) Focus() {
	c.Send(Focus)
}

// Blur removes the Focus state on the model.  When the model is blurred it can
// not receive keyboard input and the cursor will be hidden.
func (c *Input) Blur() {
	c.Send(Blur)
}

// Quit Input
func (c *Input) Quit() {
	c.Send(Quit)
}

// Value returns the value of the text input.
func (c *Input) Value() string {
	return c.Model.Value()
}

// Cursor returns the cursor position.
func (c *Input) Cursor() int {
	return c.Model.Cursor()
}

// Blink returns whether or not to draw the cursor.
func (c *Input) Blink() bool {
	return c.Model.Blink()
}

// SetCursor moves the cursor to the given position. If the position is
// out of bounds the cursor will be moved to the start or end accordingly.
func (c *Input) SetCursor(pos int) {
	c.Model.SetCursor(pos)
}

// Focused returns the focus state on the model.
func (c *Input) Focused() bool {
	return c.Model.Focused()
}

// CursorStart moves the cursor to the start of the input field.
func (c *Input) CursorStart() {
	c.Model.CursorStart()
}

// CursorEnd moves the cursor to the end of the input field.
func (c *Input) CursorEnd() {
	c.Model.CursorEnd()
}

// Reset sets the input to its default state with no input. Returns whether
// or not the cursor blink should reset.
func (c *Input) Reset() bool {
	return c.Model.Reset()
}

// CursorMode returns the model's cursor mode. For available cursor modes, see
// type CursorMode.
func (c *Input) CursorMode() CursorMode {
	return newCursorMode(c.Model.CursorMode())
}

// SetCursorMode sets the model's cursor mode. This method returns a command.
//
// For available cursor modes, see type CursorMode.
func (c *Input) SetCursorMode(model CursorMode) {
	c.Model.SetCursorMode(model.Map())
}

func (c *Input) Init() tea.Cmd {

	c.Model.Prompt = c.Prompt
	c.Model.Placeholder = c.Placeholder
	c.Model.BlinkSpeed = c.BlinkSpeed
	c.Model.EchoMode = textinput.EchoMode(c.EchoMode)
	c.Model.EchoCharacter = c.EchoCharacter
	c.Model.PromptStyle = c.PromptStyle
	c.Model.TextStyle = c.TextStyle
	c.Model.BackgroundStyle = c.BackgroundStyle
	c.Model.PlaceholderStyle = c.PlaceholderStyle
	c.Model.CursorStyle = c.CursorStyle
	c.Model.CharLimit = c.CharLimit

	return tea.Batch(textinput.Blink, func() tea.Msg {
		return c.Status
	})
}

func (c *Input) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, c.QuitKey):
			c.Model.Blur()
			return c, tea.Quit
		}
	case Status:
		c.Status = msg
		switch msg {
		case Focus:
			cmds = append(cmds, c.Model.Focus())
		case Blur:
			c.Model.Blur()
		case Quit:
			c.Model.Blur()
			return c, tea.Quit
		}
	}

	model, modelCmd := c.Model.Update(msg)
	c.Model = model
	cmds = append(cmds, modelCmd)

	return c, tea.Batch(cmds...)
}

func (c *Input) View() string {
	return c.Model.View()
}
