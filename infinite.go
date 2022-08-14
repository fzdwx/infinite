package infinite

import (
	"github.com/fzdwx/infinite/components/input/confirm"
	"github.com/fzdwx/infinite/components/input/text"
	"github.com/fzdwx/infinite/components/progress"
	cs "github.com/fzdwx/infinite/components/selection/confirm"
	"github.com/fzdwx/infinite/components/selection/multiselect"
	"github.com/fzdwx/infinite/components/selection/singleselect"
	"github.com/fzdwx/infinite/components/spinner"
)

// NewMultiSelect new multi select
func NewMultiSelect(choices []string, ops ...multiselect.Option) *multiselect.Select {
	return multiselect.New(choices, ops...)
}

// NewSingleSelect new single select
func NewSingleSelect(choices []string, ops ...singleselect.Option) *singleselect.Select {
	return singleselect.New(choices, ops...)
}

// NewSpinner new spinner
func NewSpinner(ops ...spinner.Option) *spinner.Spinner {
	return spinner.New(ops...)
}

// NewText new text input
func NewText(ops ...text.Option) *text.Text {
	return text.New(ops...)
}

// NewConfirm new confirm
func NewConfirm(ops ...confirm.Option) *confirm.Confirm {
	return confirm.New(ops...)
}

// NewConfirmWithSelection new confirm with Selection
func NewConfirmWithSelection(ops ...cs.Option) *cs.Confirm {
	return cs.WithSelection(ops...)
}

// NewProgressGroup new progress group with count.
func NewProgressGroup(processCnt int) *progress.Group {
	return progress.NewGroupWithCount(processCnt)
}
