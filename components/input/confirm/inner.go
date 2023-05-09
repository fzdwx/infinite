package confirm

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/pkg/strx"
	"github.com/fzdwx/infinite/style"
	"github.com/fzdwx/infinite/theme"
)

// inner is confirm impl
type inner struct {
	input   *components.Input
	program *tea.Program

	// the KeyMap of Confirm
	KeyMap KeyMap
	Help   help.Model
	// display help? default is false
	DisplayHelp bool
	// default is false
	Value      bool
	ValueStyle *style.Style

	FocusStyle   *Style
	UnFocusStyle *Style

	// OutputResult is whether to output the result to the screen
	OutputResult bool

	status components.Status
}

func newInner() *inner {
	i := &inner{
		input:        components.NewInput(),
		KeyMap:       DefaultKeyMap(),
		Help:         help.New(),
		DisplayHelp:  false,
		Value:        false,
		ValueStyle:   theme.DefaultTheme.ChoiceTextStyle.Underline(),
		status:       components.Normal,
		OutputResult: true,
		UnFocusStyle: UnFocusStyle(),
		FocusStyle:   FocusStyle(),
	}

	i.input.Prompt = "Are you handsome?"

	return i
}

// Init confirm
func (i *inner) Init() tea.Cmd {
	focusPrompt := strx.NewFluent().
		Style(i.FocusStyle.SymbolStyle, i.FocusStyle.Symbol).
		Style(i.FocusStyle.PromptStyle, i.input.Prompt).
		Style(i.FocusStyle.NoticeStyle, i.FocusStyle.Notice).
		Style(i.FocusStyle.IntervalStyle, i.FocusStyle.Interval).
		String()

	unFocusPrompt := strx.NewFluent().
		Style(i.UnFocusStyle.SymbolStyle, i.UnFocusStyle.Symbol).
		Style(i.UnFocusStyle.PromptStyle, i.input.Prompt).
		Style(i.UnFocusStyle.NoticeStyle, i.UnFocusStyle.Notice).
		Style(i.UnFocusStyle.IntervalStyle, i.UnFocusStyle.Interval[:len(i.UnFocusStyle.Interval)-1]).
		String()

	i.input.OutputResult = false
	i.input.Init()

	i.input.Model.Prompt = focusPrompt
	i.input.UnFocusPrompt = unFocusPrompt

	return components.FocusCmd
}

func (i *inner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msgCast := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msgCast, i.KeyMap.Quit):
			return i, tea.Quit
		case key.Matches(msgCast, i.KeyMap.Yes):
			msg = i.finish()
			i.Value = true
		case key.Matches(msgCast, i.KeyMap.No):
			msg = i.finish()
			i.Value = false
		default:
			// discard, maybe output some error msg to user?
			msg = nil
		}
	}

	_, cmd := i.input.Update(msg)
	return i, cmd
}

func (i *inner) View() string {
	if components.IsFinish(i.status) && !i.OutputResult {
		return strx.Empty
	}

	builder := strx.NewFluent().Write(i.input.View())

	if components.IsFinish(i.status) {
		builder.Style(i.ValueStyle, strx.BoolMapYesOrNo(i.Value))
	}

	if !i.DisplayHelp {
		builder.NewLine().Write(i.Help.View(i.KeyMap))
	}

	if components.IsFinish(i.status) && i.OutputResult {
		builder.NewLine()
	}

	return builder.String()
}

func (i *inner) SetProgram(program *tea.Program) {
	i.program = program
	i.input.SetProgram(program)
}

func (i *inner) finish() tea.Msg {
	i.status = components.Finish
	return i.status
}
