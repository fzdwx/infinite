package components

import (
	"github.com/charmbracelet/bubbles/key"
)

var (
	InterruptKey = key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("^C", "kill program"),
	)
)
