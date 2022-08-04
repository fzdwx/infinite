package input

type Text struct {
	inner *Component
}

func New(ops ...Option) *Text {
	i := &Text{inner: NewComponent()}

	i.Apply(ops...)
	return i
}

// Apply options on Select
func (i *Text) Apply(ops ...Option) *Text {
	if len(ops) > 0 {
		for _, option := range ops {
			option(i)
		}
	}
	return i
}

func (i *Text) Show() error {
	return i.inner.Start()
}

// Focus sets the Focus state on the model. When the model is in Focus it can
// receive keyboard input and the cursor will be hidden.
func (i *Text) Focus() {
	i.inner.Focus()
}

// Blur removes the Focus state on the model.  When the model is blurred it can
// not receive keyboard input and the cursor will be hidden.
func (i *Text) Blur() {
	i.inner.Blur()
}

// Quit Component
func (i *Text) Quit() {
	i.inner.Quit()
}

// Value returns the value of the text input.
func (i *Text) Value() string {
	return i.inner.Value()
}

// Focused returns the focus state on the model.
func (i *Text) Focused() bool {
	return i.inner.Focused()
}

// CursorMode returns the model's cursor mode. For available cursor modes, see
// type CursorMode.
func (i *Text) CursorMode() CursorMode {
	return i.inner.CursorMode()
}

// SetCursorMode sets the model's cursor mode. This method returns a command.
//
// For available cursor modes, see type CursorMode.
func (i *Text) SetCursorMode(model CursorMode) {
	i.inner.SetCursorMode(model)
}
