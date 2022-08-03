package inf

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/components/selectd/multiselect"
	"github.com/fzdwx/infinite/components/selectd/singleselect"
)

type (
	Components interface {
		// Start components
		//
		// example:
		//
		// func (is *innerMultiSelect) Start() error {
		//	return tea.NewProgram(is).Start()
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

func NewMultiSelect(choices []string, ops ...multiselect.Option) *multiselect.Select {
	return multiselect.New(choices, ops...)
}

func NewSingleSelect(choices []string, ops ...singleselect.Option) *singleselect.Select {
	return singleselect.New(choices, ops...)
}
