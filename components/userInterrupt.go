package components

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

var (
	InterruptKey = key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("^C", "kill program"),
	)

	OnUserInterrupt = func(p *tea.Program) {
		os.Exit(0)
	}
)
