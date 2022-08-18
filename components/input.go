package components

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/pkg/strx"
	"github.com/fzdwx/infinite/style"
	"time"
)

var (
	InputDefaultRequired            = false
	InputDefaultRequiredMsg         = style.New().Fg(color.Red).Render("please input text!")
	InputDefaultRequiredMsgKeepTime = time.Second * 3
	InputDefaultStatus              = Focus
	InputDefaultPrompt              = "> "
	InputDefaultValue               = strx.Empty
	InputDefaultBlinkSpeed          = time.Millisecond * 530
	InputDefaultEchoMode            = EchoNormal
	InputDefaultEchoCharacter       = '*'
	InputDefaultCharLimit           = 0
	InputDefaultPlaceholderStyle    = style.New().Fg(color.Gray)
	InputDefaultPromptStyle         = style.New()
	InputDefaultTextStyle           = style.New()
	InputDefaultBackgroundStyle     = style.New()
	InputDefaultCursorStyle         = style.New()
	cleanRequiredMsg                = func(i int) func(t time.Time) tea.Msg {
		return func(t time.Time) tea.Msg {
			return cleanRequired(i)
		}
	}
)

func InputDefaultKeyMap() InputKeyMap {
	return InputKeyMap{
		Confirm: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "confirm input value"),
		),
		Quit: InterruptKey,
	}
}

type cleanRequired int

type InputKeyMap struct {
	Confirm key.Binding
	Quit    key.Binding
}

type (
	// Input the Input component.
	Input struct {
		Model           textinput.Model
		program         *tea.Program
		showRequiredMsg bool
		cleanId         int

		Required                 bool
		RequiredMsg              string
		RequiredMsgKeepAliveTime time.Duration

		BlinkSpeed    time.Duration
		Status        Status
		EchoMode      EchoMode
		EchoCharacter rune
		// CharLimit is the maximum amount of characters this Input element will
		// accept. If 0 or less, there's no limit.
		CharLimit int

		Prompt       string
		DefaultValue string

		PromptStyle       *style.Style
		DefaultValueStyle *style.Style

		TextStyle       *style.Style
		BackgroundStyle *style.Style
		CursorStyle     *style.Style
		KeyMap          InputKeyMap
	}
)

// Focus sets the Focus state on the model. When the model is in Focus it can
// receive keyboard Input and the cursor will be hidden.
func (in *Input) Focus() {
	in.program.Send(Focus)
}

// Blur removes the Focus state on the model.  When the model is blurred it can
// not receive keyboard Input and the cursor will be hidden.
func (in *Input) Blur() {
	in.program.Send(Blur)
}

// Value returns the value of the text Input.
func (in *Input) Value() string {
	value := in.Model.Value()

	if len(value) == 0 {
		value = in.DefaultValue
	}

	return value
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

// CursorStart moves the cursor to the start of the Input field.
func (in *Input) CursorStart() {
	in.Model.CursorStart()
}

// CursorEnd moves the cursor to the end of the Input field.
func (in *Input) CursorEnd() {
	in.Model.CursorEnd()
}

// Reset sets the Input to its default state with no Input. Returns whether
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
	in.Model.Placeholder = in.DefaultValue
	in.Model.BlinkSpeed = in.BlinkSpeed
	in.Model.EchoMode = textinput.EchoMode(in.EchoMode)
	in.Model.EchoCharacter = in.EchoCharacter
	in.Model.PromptStyle = in.PromptStyle.Inner()
	in.Model.TextStyle = in.TextStyle.Inner()
	in.Model.BackgroundStyle = in.BackgroundStyle.Inner()
	in.Model.PlaceholderStyle = in.DefaultValueStyle.Inner()
	in.Model.CursorStyle = in.CursorStyle.Inner()
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
		case key.Matches(msg, in.KeyMap.Confirm):
			// todo Verification function can be added
			return in.confirm()
		case key.Matches(msg, in.KeyMap.Quit):
			OnUserInterrupt(in.program)
		}
	case Status:
		in.Status = msg
		switch msg {
		case Focus:
			cmds = append(cmds, in.Model.Focus())
		case Blur:
			in.Model.Blur()
		case Finish:
			return in.finish()
		}
	case cleanRequired:
		in.cleanRequiredMsg(msg)
	}

	model, modelCmd := in.Model.Update(msg)
	in.Model = model
	cmds = append(cmds, modelCmd)

	return in, tea.Batch(cmds...)
}

func (in *Input) View() string {
	builder := strx.NewFluent().Write(in.Model.View())

	if in.showRequiredMsg {
		builder.NewLine().Write(in.RequiredMsg)
	}

	return builder.String()
}

func (in *Input) SetProgram(program *tea.Program) {
	in.program = program
}

// confirm input val, if the verification passes, it will exit.
func (in *Input) confirm() (tea.Model, tea.Cmd) {
	if in.shouldShowRequiredMsg() {
		in.showRequiredMsg = true
		in.cleanId++
		return in, tea.Tick(in.RequiredMsgKeepAliveTime, cleanRequiredMsg(in.cleanId))
	}

	return in.finish()
}

func (in *Input) finish() (tea.Model, tea.Cmd) {
	in.Model.Blur()
	return in, tea.Quit
}

func (in *Input) shouldShowRequiredMsg() bool {
	return in.Required && len(in.Model.Value()) == 0
}

func (in *Input) cleanRequiredMsg(msg cleanRequired) {
	if int(msg) == in.cleanId {
		in.showRequiredMsg = false
	}
}
