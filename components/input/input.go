package input

type Input struct {
	inner *Component
}

func New(ops ...Option) *Input {
	i := &Input{inner: NewComponent()}

	i.Apply(ops...)
	return i
}

// Apply options on Select
func (i *Input) Apply(ops ...Option) *Input {
	if len(ops) > 0 {
		for _, option := range ops {
			option(i)
		}
	}
	return i
}

func (i *Input) Show() error {
	return i.inner.Start()
}

// Focus sets the Focus state on the model. When the model is in Focus it can
// receive keyboard input and the cursor will be hidden.
func (i *Input) Focus() {
	i.inner.Focus()
}

// Blur removes the Focus state on the model.  When the model is blurred it can
// not receive keyboard input and the cursor will be hidden.
func (i *Input) Blur() {
	i.inner.Blur()
}

// Quit Component
func (i *Input) Quit() {
	i.inner.Quit()
}

// Value returns the value of the text input.
func (i *Input) Value() string {
	return i.inner.Value()
}

// Focused returns the focus state on the model.
func (i *Input) Focused() bool {
	return i.inner.Focused()
}

// CursorMode returns the model's cursor mode. For available cursor modes, see
// type CursorMode.
func (i *Input) CursorMode() CursorMode {
	return i.inner.CursorMode()
}

// SetCursorMode sets the model's cursor mode. This method returns a command.
//
// For available cursor modes, see type CursorMode.
func (i *Input) SetCursorMode(model CursorMode) {
	i.inner.SetCursorMode(model)
}
