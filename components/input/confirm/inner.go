package confirm

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/strx"
	"github.com/fzdwx/infinite/style"
)

// inner is confirm impl
type inner struct {
	input   *components.Input
	program tea.Program

	/* option start */
	// the KeyMap of Confirm
	KeyMap KeyMap
	Help   help.Model
	// display help? default is false
	DisplayHelp bool
	// default is false
	Value       bool
	Notice      string
	NoticeStyle lipgloss.Style
	Symbol      string
	SymbolStyle lipgloss.Style
	/* option end */
}

func newInner() *inner {
	i := &inner{
		input:       components.NewInput(),
		KeyMap:      DefaultKeyMap,
		Help:        help.New(),
		DisplayHelp: false,
		Value:       false,
		Notice:      " ( y/N ) ",
		NoticeStyle: style.New(),
		Symbol:      "?",
		SymbolStyle: style.New().Foreground(color.Special),
	}

	i.input.Prompt = "Are you handsome?"

	return i
}

// Init confirm
func (i *inner) Init() tea.Cmd {
	i.input.Prompt = strx.NewFluent().
		Write(i.SymbolStyle.Render(i.Symbol)).
		Write(i.input.Prompt).
		Write(i.NoticeStyle.Render(i.Notice)).
		String()

	i.input.Init()

	return components.FocusCmd
}

func (i *inner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msgCast := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msgCast, i.KeyMap.Quit):
			msg = components.Quit
		case key.Matches(msgCast, i.KeyMap.Yes):
			msg = components.Quit
			i.Value = true
		case key.Matches(msgCast, i.KeyMap.No):
			msg = components.Quit
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
	if i.DisplayHelp {
		return strx.NewFluent().
			Write(i.input.View()).
			NewLine().
			Write(i.Help.View(i.KeyMap)).
			String()
	}

	return i.input.View()
}

func (i *inner) SetProgram(program *tea.Program) {
	i.input.SetProgram(program)
}
