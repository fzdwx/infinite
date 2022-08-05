package confirm

type Confirm struct {
	inner *inner
}

func New(ops ...Option) *Confirm {
	c := &Confirm{inner: newInner()}

	c.Apply(ops...)

	return c
}

// Display Confirm component.
func (c *Confirm) Display() error {
	return c.inner.Start()
}

// Value returns whether the user has chosen to confirm or deny.
func (c Confirm) Value() bool {
	return c.inner.value
}

// Apply options on Confirm
func (c *Confirm) Apply(ops ...Option) *Confirm {
	if len(ops) > 0 {
		for _, option := range ops {
			option(c)
		}
	}
	return c
}
