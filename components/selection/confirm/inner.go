package confirm

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/pkg/strx"
	"strings"
)

type inner struct {
	selection       *components.Selection
	focusInterval   string
	unFocusInterval string
	keyMap          KeyMap
	Value           bool
	choice          bool
	outPutResult    bool
}

func newInner(selection *components.Selection) *inner {
	return &inner{selection: selection}
}

func (i *inner) Init() tea.Cmd {
	return i.selection.Init()
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
		}
	}

	return i, nil
}

func (i *inner) View() string {
	rows := strx.RemoveEmpty(strings.Split(i.selection.View(), strx.NewLine))
	row := strx.NewFluent().
		Write(rows[0]).
		Write(i.interval()).
		Write(rows[1]).Space().Write("/").Space().Write(rows[2])

	if i.outPutResult {
		row.NewLine()
	}

	return row.String()
}

func (i *inner) SetProgram(program *tea.Program) {
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
