package inf

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/components/selectd/multiselect"
	"github.com/fzdwx/infinite/components/selectd/singleselect"
	"github.com/fzdwx/infinite/components/spinner"
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

		tea.Model
	}
)

func NewMultiSelect(choices []string, ops ...multiselect.Option) *multiselect.Select {
	return multiselect.New(choices, ops...)
}

func NewSingleSelect(choices []string, ops ...singleselect.Option) *singleselect.Select {
	return singleselect.New(choices, ops...)
}

func NewSpinner(ops ...spinner.Option) *spinner.Spinner {
	return spinner.New(ops...)
}
