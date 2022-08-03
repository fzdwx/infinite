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

	return ms.Apply(ops...)
}

// Show startup MultiSelect
func (ms MultiSelect) Show(prompt ...string) ([]int, error) {
	ms.Apply(WithPrompt(prompt...))

	ms.inner.renderColor()

	err := ms.inner.Start()
	if err != nil {
		return nil, eris.Wrap(err, "start inner multi select fail")
	}

	return ms.inner.value(), nil
}

// Apply options on MultiSelect
func (ms *MultiSelect) Apply(ops ...Option) *MultiSelect {
	if len(ops) > 0 {
		for _, option := range ops {
			option(ms)
		}
	}
	return ms
}
