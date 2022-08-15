package confirm

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/pkg/strx"
	"github.com/fzdwx/infinite/style"
)

// Confirm with components.Selection
type Confirm struct {
	startUp *components.StartUp
	inner   *inner

	KeyMap          KeyMap
	Yes             string
	No              string
	Symbol          string
	Prompt          string
	FocusInterval   string
	UnFocusInterval string
	OutputResult    bool
	DefaultVal      bool
	SymbolStyle     *style.Style
	PromptStyle     *style.Style
	ChoiceStyle     *style.Style

	ops []Option
}

type KeyMap struct {
	Switch key.Binding
	Choice key.Binding
}

// WithSelection new Confirm with components.Selection
func WithSelection(ops ...Option) *Confirm {
	c := &Confirm{
		Yes:             Yes,
		No:              No,
		KeyMap:          DefaultKeyBinding(),
		SymbolStyle:     SymbolStyle,
		Symbol:          Symbol,
		Prompt:          Prompt,
		PromptStyle:     PromptStyle,
		ChoiceStyle:     ChoiceStyle,
		FocusInterval:   FocusInterval,
		UnFocusInterval: UnFocusInterval,
		OutputResult:    true,
		DefaultVal:      false,
		ops:             ops,
	}
	return c
}

// Display Confirm
func (c *Confirm) Display() (bool, error) {
	for _, op := range c.ops {
		op(c)
	}

	c.init()

	c.startUp = components.NewStartUp(c.inner)
	err := c.startUp.Start()
	return c.inner.Value, err
}

// init Adjust the `components.Selection` to fit the ` Confirm ` scene
func (c *Confirm) init() {
	c.inner = newInner(components.NewSelection([]string{c.No, c.Yes}))
	c.inner.focusInterval = c.FocusInterval
	c.inner.unFocusInterval = c.UnFocusInterval
	c.inner.selection.EnableFilter = false
	c.inner.selection.ShowHelp = false
	c.inner.selection.ChoiceTextStyle = c.ChoiceStyle
	c.inner.outputResult = c.OutputResult
	c.inner.selection.Prompt = strx.NewFluent().
		Style(c.SymbolStyle, c.Symbol).
		Style(c.PromptStyle, c.Prompt).
		String()

	c.inner.selection.RowRender = func(CursorSymbol string, HintSymbol string, choice string) string {
		return choice
	}

	// default true
	c.inner.DefaultVal = c.DefaultVal

	c.inner.keyMap = c.KeyMap
}
