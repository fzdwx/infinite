package confirm

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/pkg/strx"
	"github.com/fzdwx/infinite/style"
)

type KeyMap struct {
	Switch key.Binding
	Choice key.Binding
	Quit   key.Binding
}

// Confirm with components.Selection
type Confirm struct {
	startUp *components.StartUp
	inner   *inner

	KeyMap       KeyMap
	Yes          string
	No           string
	OutputResult bool
	DefaultVal   bool

	FocusSymbol     string
	UnFocusSymbol   string
	FocusInterval   string
	UnFocusInterval string
	Prompt          string

	FocusSymbolStyle     *style.Style
	UnFocusSymbolStyle   *style.Style
	FocusIntervalStyle   *style.Style
	UnFocusIntervalStyle *style.Style
	PromptStyle          *style.Style
	ChoiceStyle          *style.Style

	ops []Option
}

// WithSelection new Confirm with components.Selection
func WithSelection(ops ...Option) *Confirm {
	c := &Confirm{
		Yes:          Yes,
		No:           No,
		KeyMap:       DefaultKeyBinding(),
		OutputResult: true,
		DefaultVal:   false,
		ops:          ops,

		FocusSymbol:          FocusSymbol,
		UnFocusSymbol:        UnFocusSymbol,
		FocusInterval:        FocusInterval,
		UnFocusInterval:      UnFocusInterval,
		Prompt:               Prompt,
		FocusSymbolStyle:     FocusSymbolStyle,
		UnFocusSymbolStyle:   UnFocusSymbolStyle,
		FocusIntervalStyle:   FocusIntervalStyle,
		UnFocusIntervalStyle: UnFocusIntervalStyle,
		PromptStyle:          PromptStyle,
		ChoiceStyle:          ChoiceStyle,
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
	// default true
	c.inner.DefaultVal = c.DefaultVal
	c.inner.keyMap = c.KeyMap
	c.inner.focusPrompt = strx.NewFluent().Style(c.FocusSymbolStyle, c.FocusSymbol).Style(c.PromptStyle, c.Prompt).String()
	c.inner.unFocusPrompt = strx.NewFluent().Style(c.UnFocusSymbolStyle, c.UnFocusSymbol).Style(c.PromptStyle, c.Prompt).String()
	c.inner.selection.EnableFilter = false
	c.inner.selection.ShowHelp = false
	c.inner.selection.ChoiceTextStyle = c.ChoiceStyle
	c.inner.outputResult = c.OutputResult
	c.inner.selection.RowRender = func(CursorSymbol string, HintSymbol string, choice string) string {
		return choice
	}
}
