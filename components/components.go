package components

import tea "github.com/charmbracelet/bubbletea"

type (
	Components struct {
		tea.Model
		P *tea.Program
	}
)

// Start Components
func (c *Components) Start(ops ...tea.ProgramOption) error {
	c.P = tea.NewProgram(c, ops...)

	return c.P.Start()
}

// Kill Components
func (c *Components) Kill() {
	c.P.Kill()
}
