package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type model struct {
	buf strings.Builder
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.buf.WriteString(msg.String())
		m.buf.WriteString("\n")
	}

	return m, nil
}

func (m model) View() string {
	return m.buf.String()
}

func main() {
	if _, err := tea.NewProgram(model{buf: strings.Builder{}}).Run(); err != nil {
		return
	}
}
