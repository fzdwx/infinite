package multiselect

import (
	"github.com/fzdwx/infinite/components/selection"
	"github.com/rotisserie/eris"
)

type Select struct {
	inner *selection.Component
}

func New(choices []string, ops ...Option) *Select {
	ms := &Select{
		inner: selection.NewComponent(choices),
	}

	return ms.Apply(ops...)
}

// Show startup Select
func (ms Select) Show(prompt ...string) ([]int, error) {
	ms.Apply(WithPrompt(prompt...))

	ms.inner.RenderColor()

	err := ms.inner.Start()
	if err != nil {
		return nil, eris.Wrap(err, "start inner selection fail")
	}

	return ms.inner.Value(), nil
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
