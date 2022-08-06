package multiselect

import (
	"github.com/fzdwx/infinite/components"
	"github.com/rotisserie/eris"
)

type Select struct {
	startUp *components.StartUp
	inner   *components.Selection
}

func New(choices []string, ops ...Option) *Select {
	inner := components.NewSelection(choices)
	startUp := components.NewStartUp(inner)

	ms := &Select{
		inner:   inner,
		startUp: startUp,
	}

	return ms.Apply(ops...)
}

// Display startup Select
func (ms *Select) Display(prompt ...string) ([]int, error) {
	ms.Apply(WithPrompt(prompt...))

	ms.inner.RenderColor()

	err := ms.startUp.Start()
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
