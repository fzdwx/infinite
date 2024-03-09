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

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Switch, k.Choice}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Switch, k.Choice},
		{k.Quit},
	}
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
	ShowHelp     bool

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
		ShowHelp:     true,
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
func (i *Confirm) Display() (bool, error) {
	for _, op := range i.ops {
		op(i)
	}

	i.init()

	i.startUp = components.NewStartUp(i.inner)
	_, err := i.startUp.Run()
	return i.inner.Value, err
}

// Status get confirm current status
func (i *Confirm) Status() components.Status {
	return i.inner.status
}

// init Adjust the `components.Selection` to fit the ` Confirm ` scene
func (i *Confirm) init() {
	i.inner = newInner(components.NewSelection([]string{i.No, i.Yes}))
	i.inner.focusInterval = i.FocusInterval
	i.inner.unFocusInterval = i.UnFocusInterval
	i.inner.ShowHelp = i.ShowHelp
	// default true
	i.inner.DefaultVal = i.DefaultVal
	i.inner.keyMap = i.KeyMap
	i.inner.focusPrompt = strx.NewFluent().Style(i.FocusSymbolStyle, i.FocusSymbol).Style(i.PromptStyle, i.Prompt).String()
	i.inner.unFocusPrompt = strx.NewFluent().Style(i.UnFocusSymbolStyle, i.UnFocusSymbol).Style(i.PromptStyle, i.Prompt).String()
	i.inner.selection.Keymap.ToggleFilter.SetEnabled(false)
	i.inner.selection.ShowHelp = false
	i.inner.selection.ShowPaginator = false
	i.inner.selection.Paginator.PerPage = 2
	i.inner.selection.ChoiceTextStyle = i.ChoiceStyle
	i.inner.outputResult = i.OutputResult
	i.inner.selection.RowRender = func(CursorSymbol string, HintSymbol string, choice string) string {
		return choice
	}
}
