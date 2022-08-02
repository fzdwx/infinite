package inf

import (
	tea "github.com/charmbracelet/bubbletea"
)

type (
	Components interface {
		// Start components
		//
		// example:
		//
		//
		// func (is *innerMultiSelect) Start() error {
		//	return startUp(is)
		// }
		Start() error

		// Init is the first function that will be called. It returns an optional
		// initial command. To not perform an initial command return nil.
		Init() tea.Cmd

		// Update is called when a message is received. Use it to inspect messages
		// and, in response, update the model and/or send a command.
		Update(tea.Msg) (tea.Model, tea.Cmd)

		// View renders the program's UI, which is just a string. The view is
		// rendered after every Update.
		View() string
	}
)

func startUp(c Components) error {
	return tea.NewProgram(c).Start()
}

func NewMultiSelect(choices []string, ops ...MultiSelectOption) *MultiSelect {

	ms := &MultiSelect{
		inner: newInnerSelect(choices),
	}

	return ms.apply(ops...)
}
