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
	InputDefaultPlaceholder         = strx.Empty
	InputPlaceholderIsDefault       = true
	InputDefaultBlinkSpeed          = time.Millisecond * 530
	InputDefaultEchoMode            = EchoNormal
	InputDefaultEchoCharacter       = '*'
	InputDefaultCharLimit           = 0
	InputDefaultPlaceholderStyle    = style.New().Fg(color.Gray)
	InputDefaultPromptStyle         = style.New()
	InputDefaultTextStyle           = style.New()
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
		FocusPrompt     string
		UnFocusPrompt   string

		OutputResult             bool
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

		Prompt      string
		Placeholder string

		// set the default value required
		PlaceholderIsDefault bool

		PromptStyle       *style.Style
		DefaultValueStyle *style.Style

		FocusSymbol     string
		UnFocusSymbol   string
		FocusInterval   string
		UnFocusInterval string

		FocusSymbolStyle     *style.Style
		UnFocusSymbolStyle   *style.Style
		FocusIntervalStyle   *style.Style
		UnFocusIntervalStyle *style.Style

		TextStyle   *style.Style
		CursorStyle *style.Style
		KeyMap      InputKeyMap
	}
)

// Focus sets the Focus state on the model. When the model is in Focus it can
// receive keyboard Input and the cursor will be hidden.
func (i *Input) Focus() {
	i.program.Send(Focus)
}

// Blur removes the Focus state on the model.  When the model is blurred it can
// not receive keyboard Input and the cursor will be hidden.
func (i *Input) Blur() {
	i.program.Send(Blur)
}

// Value returns the value of the text Input.
func (i *Input) Value() string {
	value := i.Model.Value()

	if len(value) == 0 && i.PlaceholderIsDefault {
		value = i.Placeholder
	}

	return value
}

// Cursor returns the cursor position.
// deprecated: use Position() instead.
func (i *Input) Cursor() int {
	return i.Model.Position()
}

func (i *Input) Position() int {
	return i.Model.Position()
}

// Blink returns whether or not to draw the cursor.
func (i *Input) Blink() bool {
	return i.Model.Cursor.Blink
}

// SetCursor moves the cursor to the given position. If the position is
// out of bounds the cursor will be moved to the start or end accordingly.
func (i *Input) SetCursor(pos int) {
	i.Model.SetCursor(pos)
}

// Focused returns the focus state on the model.
func (i *Input) Focused() bool {
	return i.Model.Focused()
}

// CursorStart moves the cursor to the start of the Input field.
func (i *Input) CursorStart() {
	i.Model.CursorStart()
}

// CursorEnd moves the cursor to the end of the Input field.
func (i *Input) CursorEnd() {
	i.Model.CursorEnd()
}

// Reset sets the Input to its default state with no Input. Returns whether
// or not the cursor blink should reset.
func (i *Input) Reset() {
	i.Model.Reset()
}

// CursorMode returns the model's cursor mode. For available cursor modes, see
// type CursorMode.
func (i *Input) CursorMode() CursorMode {
	return newCursorMode(i.Model.Cursor.Mode())
}

// SetCursorMode sets the model's cursor mode. This method returns a command.
//
// For available cursor modes, see type CursorMode.
func (i *Input) SetCursorMode(model CursorMode) {
	i.Model.Cursor.SetMode(model.Map())
}

func (i *Input) Init() tea.Cmd {
	i.FocusPrompt = strx.NewFluent().
		Style(i.FocusSymbolStyle, i.FocusSymbol).
		Style(i.PromptStyle, i.Prompt).
		Style(i.FocusIntervalStyle, i.FocusInterval).
		String()

	i.UnFocusPrompt = strx.NewFluent().
		Style(i.UnFocusSymbolStyle, i.UnFocusSymbol).
		Style(i.PromptStyle, i.Prompt).
		Style(i.UnFocusIntervalStyle, i.UnFocusInterval).
		String()

	i.Model.Prompt = i.FocusPrompt
	i.Model.Placeholder = i.Placeholder
	i.Model.Placeholder = i.Placeholder
	i.Model.Cursor.BlinkSpeed = i.BlinkSpeed
	i.Model.EchoMode = textinput.EchoMode(i.EchoMode)
	i.Model.EchoCharacter = i.EchoCharacter
	i.Model.TextStyle = i.TextStyle.Inner()
	i.Model.PlaceholderStyle = i.DefaultValueStyle.Inner()
	i.Model.Cursor.Style = i.CursorStyle.Inner()
	i.Model.CharLimit = i.CharLimit

	return tea.Batch(textinput.Blink, func() tea.Msg {
		return i.Status
	})
}

func (i *Input) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch {
		case key.Matches(msg, i.KeyMap.Confirm):
			// todo Verification function can be added
			return i.confirm()
		case key.Matches(msg, i.KeyMap.Quit):
			i.Status = Quit
			return i, tea.Quit
		}

	case Status:
		i.Status = msg
		switch msg {
		case Focus:
			cmds = append(cmds, i.Model.Focus())
		case Blur:
			i.Model.Blur()
		case Finish:
			return i.finish()
		default: // do nothing
		}
	case cleanRequired:
		i.cleanRequiredMsg(msg)
	}

	model, modelCmd := i.Model.Update(msg)
	i.Model = model
	cmds = append(cmds, modelCmd)

	return i, tea.Batch(cmds...)
}

func (i *Input) View() string {
	builder := strx.NewFluent().Write(i.Model.View())

	if i.showRequiredMsg {
		builder.NewLine().Write(i.RequiredMsg)
	}

	if IsFinish(i.Status) && i.OutputResult {
		builder.NewLine()
	}

	return builder.String()
}

func (i *Input) SetProgram(program *tea.Program) {
	i.program = program
}

// confirm input val, if the verification passes, it will exit.
func (i *Input) confirm() (tea.Model, tea.Cmd) {
	if i.shouldShowRequiredMsg() {
		i.showRequiredMsg = true
		i.cleanId++
		return i, tea.Tick(i.RequiredMsgKeepAliveTime, cleanRequiredMsg(i.cleanId))
	}

	return i.finish()
}

func (i *Input) finish() (tea.Model, tea.Cmd) {
	i.Model.Blur()
	i.Model.Prompt = i.UnFocusPrompt
	i.Status = Finish
	return i, tea.Quit
}

func (i *Input) shouldShowRequiredMsg() bool {
	return i.Required && len(i.Model.Value()) == 0
}

func (i *Input) cleanRequiredMsg(msg cleanRequired) {
	if int(msg) == i.cleanId {
		i.showRequiredMsg = false
	}
}
