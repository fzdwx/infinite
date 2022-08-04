package components

import tea "github.com/charmbracelet/bubbletea"

type (
	Component struct {
		tea.Model
		P *tea.Program
	}
)

// Start Component
func (c *Component) Start(ops ...tea.ProgramOption) error {
	c.P = tea.NewProgram(c, ops...)

	return c.P.Start()
}

// Kill Component
func (c *Component) Kill() {
	c.P.Kill()
}
