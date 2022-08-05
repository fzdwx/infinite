package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

/*

 input.component
 selection.component
 spinner.component

*/

const GlobalTickStatusDelay = time.Millisecond * 10

type (
	/*
		Components, You can use these components directly:
			 	1.input.component
			 	2.selection.component
			 	3.spinner.component
		Or use them inline in your custom component,
		for how to embed them, you can refer to the implementation of `Confirm`.
	*/
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
