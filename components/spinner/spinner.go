package spinner

import (
	"errors"
	"fmt"
	"github.com/fzdwx/infinite/components"
)

type Spinner struct {
	*components.PrintHelper
	inner   *components.Spinner
	startUp *components.StartUp
}

var (
	spinnerRunnerIsRequired = errors.New("runner is required")
)

// New Spinner
func New(ops ...Option) *Spinner {
	inner := components.NewSpinner()
	s := &Spinner{
		inner:   inner,
		startUp: components.NewStartUp(inner),
	}

	s.PrintHelper = inner.PrintHelper

	s.Apply(ops...)

	return s
}

// Apply options on Select
func (s *Spinner) Apply(ops ...Option) *Spinner {
	if len(ops) > 0 {
		for _, option := range ops {
			option(s)
		}
	}
	return s
}

// Display Spinner
func (s *Spinner) Display(runner func(spinner *Spinner)) error {
	if runner == nil {
		return spinnerRunnerIsRequired
	}

	go func() {
		runner(s)
		s.Finish()
	}()

	return s.startUp.Start()
}

// Finish quit Spinner
func (s *Spinner) Finish(prompt ...string) {
	s.refresh(prompt...)
	s.inner.Quit()
}

// Refresh Spinner prompt
func (s *Spinner) Refresh(prompt string) {
	s.refresh(prompt)
}

// Refreshf Spinner prompt
func (s *Spinner) Refreshf(format string, a ...any) {
	s.refresh(fmt.Sprintf(format, a...))
}

func (s *Spinner) refresh(prompt ...string) {
	if len(prompt) > 0 {
		s.inner.RefreshPrompt(prompt[0])
	}
}
