package confirm

import "github.com/fzdwx/infinite/components"

type Confirm struct {
	startUp *components.StartUp
	inner   *inner
}

func New(ops ...Option) *Confirm {

	i := newInner()

	c := &Confirm{
		inner:   i,
		startUp: components.NewStartUp(i),
	}

	c.Apply(ops...)

	return c
}

// Display Confirm component.
func (c *Confirm) Display() error {
	return c.startUp.Start()
}

// Value returns whether the user has chosen to confirm or deny.
func (c *Confirm) Value() bool {
	return c.inner.Value
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
