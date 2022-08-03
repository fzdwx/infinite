package singleselect

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/fzdwx/infinite/multiselect"
)

type KeyMap struct {
	Up      key.Binding
	Down    key.Binding
	Choice  key.Binding
	Confirm key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Choice, k.Confirm}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},        // first column
		{k.Choice, k.Confirm}, // second column
	}
}

func (k KeyMap) MapToMulti() multiselect.KeyMap {
	return multiselect.KeyMap{
		Up:      k.Up,
		Down:    k.Down,
		Choice:  k.Choice,
		Confirm: k.Confirm,
	}
}

var DefaultKeyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "move down"),
	),
	Choice: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "choice it"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("ctrl+c", "enter"),
		key.WithHelp("ctrl+c/enter", "quit"),
	),
}
