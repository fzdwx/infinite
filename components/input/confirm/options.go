package confirm

type Option func(confirm *Confirm)

// WithDefaultYes the confirm default use yes.
func WithDefaultYes() Option {
	return func(c *Confirm) {
		c.inner.value = true
	}
}
