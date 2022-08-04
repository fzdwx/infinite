package components

import tea "github.com/charmbracelet/bubbletea"

/*

 input.component
 selection.component
 spinner.component

*/

type (
	Components struct {
		tea.Model
		P       *tea.Program
		started bool
	}
)

// Start Components
func (c *Components) Start(ops ...tea.ProgramOption) error {
	c.P = tea.NewProgram(c, ops...)

	c.started = true
	return c.P.Start()
}

// Kill Components
func (c *Components) Kill() {
	if c.started {
		c.P.Kill()
		c.started = false
	}
}

// Send message to component
func (c *Components) Send(msg tea.Msg) {
	if c.started {
		c.P.Send(msg)
	}
}

func Seq(cmd ...tea.Cmd) tea.Cmd {
	return tea.Sequentially(cmd...)
}
