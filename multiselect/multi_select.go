package multiselect

import (
	"github.com/rotisserie/eris"
)

type MultiSelect struct {
	inner *innerMultiSelect
}

func New(choices []string, ops ...Option) *MultiSelect {
	ms := &MultiSelect{
		inner: newInnerSelect(choices),
	}

	return ms.apply(ops...)
}

// Show startup MultiSelect
func (ms MultiSelect) Show(prompt ...string) ([]int, error) {
	ms.apply(withPrompt(prompt...))

	ms.inner.renderColor()

	err := ms.inner.Start()
	if err != nil {
		return nil, eris.Wrap(err, "start inner multi select fail")
	}

	return ms.inner.Selected(), nil
}

// apply options on MultiSelect
func (ms *MultiSelect) apply(ops ...Option) *MultiSelect {
	if len(ops) > 0 {
		for _, option := range ops {
			option(ms)
		}
	}
	return ms
}
