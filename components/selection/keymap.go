package selection

import "github.com/charmbracelet/bubbles/key"

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

var DefaultMultiKeyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "move down"),
	),
	Choice: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "choice it"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "confirm and quit"),
	),
}

var DefaultSingleKeyMap = KeyMap{
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
