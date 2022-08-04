package input

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/style"
	"time"
)

const defaultBlinkSpeed = time.Millisecond * 530

type Component struct {
	components.Components

	Status Status
	Model  textinput.Model

	/* option start */
	DefaultStatus   Status
	TickStatusDelay time.Duration

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

	// CharLimit is the maximum amount of characters this input element will
	// accept. If 0 or less, there's no limit.
	CharLimit int
	/* option end */
}

func NewComponent() *Component {
	c := &Component{
		Model:            textinput.New(),
		DefaultStatus:    Focus,
		TickStatusDelay:  time.Millisecond * 50,
		Prompt:           ">",
		Placeholder:      "",
		BlinkSpeed:       defaultBlinkSpeed,
		EchoMode:         EchoNormal,
		EchoCharacter:    '*',
		PromptStyle:      style.New(),
		TextStyle:        style.New(),
		BackgroundStyle:  style.New(),
		PlaceholderStyle: style.New(),
		CursorStyle:      style.New(),
		CharLimit:        0,
	}

	c.Components = components.Components{Model: c}

	return c
}

// Focus sets the Focus state on the model. When the model is in Focus it can
// receive keyboard input and the cursor will be hidden.
func (c *Component) Focus() {
	c.Status = Focus
}

// Blur removes the Focus state on the model.  When the model is blurred it can
// not receive keyboard input and the cursor will be hidden.
func (c *Component) Blur() {
	c.Status = Blur
}

// Quit Component
func (c *Component) Quit() {
	c.Status = Quit
}

// Value returns the value of the text input.
func (c *Component) Value() string {
	return c.Model.Value()
}

// Focused returns the focus state on the model.
func (c *Component) Focused() bool {
	return c.Model.Focused()
}

// CursorMode returns the model's cursor mode. For available cursor modes, see
// type CursorMode.
func (c Component) CursorMode() CursorMode {
	return newCursorMode(c.Model.CursorMode())
}

// SetCursorMode sets the model's cursor mode. This method returns a command.
//
// For available cursor modes, see type CursorMode.
func (c *Component) SetCursorMode(model CursorMode) {
	c.Model.SetCursorMode(model.Map())
}

func (c *Component) Init() tea.Cmd {

	c.Status = c.DefaultStatus
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
		return c.DefaultStatus
	})
}

func (c *Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case Status:
		switch msg {
		case Focus:
			cmds = append(cmds, c.Model.Focus())
		case Blur:
			c.Model.Blur()
		case Quit:
			return c, tea.Quit
		}
		cmds = append(cmds, c.tickStatus(c.Status))
	}

	_, modelCmd := c.Model.Update(msg)

	cmds = append(cmds, modelCmd)
	return c, tea.Batch(cmds...)
}

func (c *Component) View() string {
	return c.Model.View()
}

func (c *Component) tickStatus(status Status) tea.Cmd {
	return tea.Tick(c.TickStatusDelay, func(t time.Time) tea.Msg {
		return status
	})
}
