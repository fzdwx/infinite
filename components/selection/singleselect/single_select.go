package singleselect

import (
	"fmt"
	"os"

	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/selection/multiselect"
)

type Select struct {
	inner *multiselect.Select
}

// New single Select
func New(choices []string, ops ...Option) *Select {
	s := &Select{inner: multiselect.New(choices)}

	s.mapMultiToSingle()

	s.Apply(ops...)
	return s
}

// Display startup Select
func (s *Select) Display(prompt ...string) (int, error) {
	hints, err := s.inner.Display(prompt...)

	if err != nil {
		return -1, err
	}

	if len(hints) == 0 {
		os.Exit(0)
	}

	return hints[0], nil
}

// Apply options on Select
func (s *Select) Apply(ops ...Option) *Select {
	if len(ops) > 0 {
		for _, option := range ops {
			option(s)
		}
	}
	return s
}

func (s *Select) mapMultiToSingle() {
	var ops []Option

	// replace row render
	ops = append(ops, WithRowRender(func(cursorSymbol string, hintSymbol string, choice string) string {
		return fmt.Sprintf("%s %s", cursorSymbol, choice)
	}))

	// replace prompt
	ops = append(ops, WithPrompt("Please selection your option:"))

	// replace key binding
	ops = append(ops, WithKeyBinding(components.DefaultSingleKeyMap()))

	s.Apply(ops...)
}
