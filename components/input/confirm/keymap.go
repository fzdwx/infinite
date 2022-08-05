package confirm

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit key.Binding
	Yes  key.Binding
	No   key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Yes, k.No, k.Quit,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

var (
	DefaultKeyMap = KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("^c", "quit confirm"),
		),
		Yes: key.NewBinding(
			key.WithKeys("y"),
			key.WithHelp("y", "yes"),
		),
		No: key.NewBinding(
			key.WithKeys("N"),
			key.WithHelp("N", "no"),
		),
	}
)
