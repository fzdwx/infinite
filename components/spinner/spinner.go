package spinner

import (
	"fmt"
	"github.com/fzdwx/infinite/components"
)

type Spinner struct {
	inner   *components.Spinner
	startUp *components.StartUp
	err     error
}

// New Spinner
func New(ops ...Option) *Spinner {
	inner := components.NewSpinner()
	s := &Spinner{
		inner:   inner,
		startUp: components.NewStartUp(inner),
	}

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
func (s *Spinner) Display(prompt ...string) *Spinner {
	go func() {
		s.err = s.startUp.Start()
		s.refresh(prompt...)
	}()
	return s
}

// Finish quit Spinner
func (s *Spinner) Finish(prompt ...string) error {
	s.refresh(prompt...)
	s.inner.Quit()

	return s.err
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
