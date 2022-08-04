package spinner

type Spinner struct {
	inner *InnerSpinner
	err   error
}

func New(ops ...Option) *Spinner {
	s := &Spinner{
		inner: NewInner(),
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

func (s *Spinner) Show() *Spinner {
	go func() {
		s.err = s.inner.Start()
	}()
	return s
}

func (s *Spinner) Finish() error {
	s.inner.Quited = true

	return s.err
}
