package confirm

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/pkg/strx"
	"strings"
)

type switchIt int

type inner struct {
	selection       *components.Selection
	focusPrompt     string
	unFocusPrompt   string
	focusInterval   string
	unFocusInterval string
	keyMap          KeyMap
	Value           bool
	DefaultVal      bool
	choice          bool
	outputResult    bool
	program         *tea.Program
}

func newInner(selection *components.Selection) *inner {
	return &inner{selection: selection}
}

func (i *inner) Init() tea.Cmd {
	cmd := i.selection.Init()

	if i.DefaultVal {
		cmd = tea.Batch(cmd, func() tea.Msg {
			return switchIt(1)
		})
	}

	return cmd
}

func (i *inner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, i.keyMap.Switch):
			i.switchIt()
		case key.Matches(msg, i.keyMap.Choice):
			i.choice = true
			return i, tea.Quit
		case key.Matches(msg, i.keyMap.Quit):
			components.OnUserInterrupt(i.program)
		}
	case switchIt:
		i.switchIt()
	}

	return i, nil
}

func (i *inner) View() string {
	rows := strx.RemoveEmpty(strings.Split(i.selection.View(), strx.NewLine))
	row := strx.NewFluent().
		Write(i.prompt()).
		Write(i.interval()).
		Write(rows[1]).Space().Write("/").Space().Write(rows[2])

	if i.outputResult {
		row.NewLine()
	}

	return row.String()
}

func (i *inner) SetProgram(program *tea.Program) {
	i.program = program
}

func (i *inner) switchIt() {
	var msg tea.Msg

	if i.Value {
		msg = tea.KeyMsg{
			Type: tea.KeyUp,
		}
	} else {
		msg = tea.KeyMsg{Type: tea.KeyDown}
	}

	i.Value = !i.Value
	i.selection.Update(msg)
}

func (i *inner) interval() string {
	if i.choice {
		return i.unFocusInterval
	}
	return i.focusInterval
}

func (i *inner) prompt() string {
	if i.choice {
		return i.unFocusPrompt
	}
	return i.focusPrompt
}
