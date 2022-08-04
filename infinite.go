package inf

import (
	"github.com/fzdwx/infinite/components/selectd/multiselect"
	"github.com/fzdwx/infinite/components/selectd/singleselect"
	"github.com/fzdwx/infinite/components/spinner"
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
