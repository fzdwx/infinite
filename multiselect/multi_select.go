package multiselect

import (
	"github.com/rotisserie/eris"
)

type Select struct {
	inner *innerMultiSelect
}

func New(choices []string, ops ...Option) *Select {
	ms := &Select{
		inner: newInnerSelect(choices),
	}

	return ms.Apply(ops...)
}

// Show startup Select
func (ms Select) Show(prompt ...string) ([]int, error) {
	ms.Apply(WithPrompt(prompt...))

	ms.inner.renderColor()

	err := ms.inner.Start()
	if err != nil {
		return nil, eris.Wrap(err, "start inner select fail")
	}

	return ms.inner.value(), nil
}

// Apply options on Select
func (ms *Select) Apply(ops ...Option) *Select {
	if len(ops) > 0 {
		for _, option := range ops {
			option(ms)
		}
	}
	return ms
}
