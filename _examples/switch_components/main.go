package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/components"
)

func main() {
	p := tea.NewProgram(&myModel{})
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

type myModel struct {
	s1 *components.Selection
	s2 *components.Selection

	current components.Components
}

func (m *myModel) Init() tea.Cmd {
	m.s1 = components.NewSelection([]string{"a", "b", "c"})
	m.s2 = components.NewSelection([]string{"1", "2", "3"})
	m.current = m.s1

	return func() tea.Msg {
		c1 := m.s1.Init()
		c2 := m.s2.Init()
		return tea.Batch(c1, c2)
	}
}

func (m *myModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			m.current = m.s2
			return m, nil
		case "right":
			m.current = m.s1
			return m, nil
		case "q":
			return m, tea.Quit
		}
	}

	updatedM, cmd := m.current.Update(msg)
	m.current = updatedM.(components.Components)

	return m, cmd
}

func (m *myModel) View() string {
	return m.current.View()
}
