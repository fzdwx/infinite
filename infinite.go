package inf

import (
	"github.com/fzdwx/infinite/components/input/confirm"
	"github.com/fzdwx/infinite/components/input/text"
	"github.com/fzdwx/infinite/components/selection/multiselect"
	"github.com/fzdwx/infinite/components/selection/singleselect"
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

func NewText(ops ...text.Option) *text.Text {
	return text.New(ops...)
}

func NewConfirm(ops ...confirm.Option) *confirm.Confirm {
	return confirm.New(ops...)
}
