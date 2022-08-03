package spinner

type Spinner struct {
	inner *innerSpinner
	err   error
}

func New() *Spinner {
	return &Spinner{
		inner: newInner(),
	}
}

func (s *Spinner) Show() *Spinner {
	go func() {
		s.err = s.inner.Start()
	}()
	return s
}

func (s *Spinner) Finish() error {
	s.inner.quited = true

	return s.err
}
