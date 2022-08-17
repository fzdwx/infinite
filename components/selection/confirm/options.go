package confirm

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/style"
	"github.com/fzdwx/infinite/theme"
)

var (
	Yes      = "Yes"
	No       = "No"
	ShowHelp = true

	FocusSymbol     = theme.DefaultTheme.FocusSymbol
	UnFocusSymbol   = theme.DefaultTheme.UnFocusSymbol
	FocusInterval   = theme.DefaultTheme.FocusInterval
	UnFocusInterval = theme.DefaultTheme.UnFocusInterval
	Prompt          = "Are you handsome?"

	FocusSymbolStyle     = theme.DefaultTheme.FocusSymbolStyle
	UnFocusSymbolStyle   = theme.DefaultTheme.UnFocusSymbolStyle
	FocusIntervalStyle   = theme.DefaultTheme.FocusIntervalStyle
	UnFocusIntervalStyle = theme.DefaultTheme.UnFocusIntervalStyle
	PromptStyle          = style.New().Bold().Fg(color.White)
	ChoiceStyle          = theme.DefaultTheme.ChoiceTextStyle.Underline()
)

// DefaultKeyBinding the Confirm default key binding.
func DefaultKeyBinding() KeyMap {
	return KeyMap{
		Switch: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "switch item"),
		),
		Choice: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "status item"),
		),
		Quit: components.InterruptKey,
	}
}

type Option func(c *Confirm)

// WithDisableShowHelp  disable show help info.
func WithDisableShowHelp() Option {
	return func(c *Confirm) {
		c.ShowHelp = false
	}
}

// WithDefaultYes  set Confirm default val is `true`
func WithDefaultYes() Option {
	return func(c *Confirm) {
		c.DefaultVal = true
	}
}

// WithDisableOutputResult on finish confirm, output result.
func WithDisableOutputResult() Option {
	return func(c *Confirm) {
		c.OutputResult = false
	}
}

// WithKeyMap replace keymap. default is DefaultKeyBinding
func WithKeyMap(keymap KeyMap) Option {
	return func(c *Confirm) {
		c.KeyMap = keymap
	}
}

// WithYes default is Yes
func WithYes(yes string) Option {
	return func(c *Confirm) {
		c.Yes = yes
	}
}

// WithNo default is No
func WithNo(no string) Option {
	return func(c *Confirm) {
		c.No = no
	}
}

// WithFocusSymbol replace FocusSymbol, default is FocusSymbol
func WithFocusSymbol(s string) Option {
	return func(c *Confirm) {
		c.FocusSymbol = s
	}
}

// WithUnFocusSymbol replace UnFocusSymbol, default is UnFocusSymbol
func WithUnFocusSymbol(s string) Option {
	return func(c *Confirm) {
		c.UnFocusSymbol = s
	}
}

// WithFocusInterval replace FocusInterval, default is FocusInterval
func WithFocusInterval(no string) Option {
	return func(c *Confirm) {
		c.FocusInterval = no
	}
}

// WithUnFocusInterval replace UnFocusInterval, default is UnFocusInterval
func WithUnFocusInterval(no string) Option {
	return func(c *Confirm) {
		c.UnFocusInterval = no
	}
}

// WithPrompt replace Prompt, default is Prompt
func WithPrompt(prompt string) Option {
	return func(c *Confirm) {
		c.Prompt = prompt
	}
}

// WithFocusSymbolStyle replace FocusSymbolStyle, default is FocusSymbolStyle
func WithFocusSymbolStyle(s *style.Style) Option {
	return func(c *Confirm) {
		c.FocusSymbolStyle = s
	}
}

// WithUnFocusSymbolStyle replace UnFocusSymbolStyle, default is UnFocusSymbolStyle
func WithUnFocusSymbolStyle(s *style.Style) Option {
	return func(c *Confirm) {
		c.UnFocusSymbolStyle = s
	}
}

// WithFocusIntervalStyle replace FocusIntervalStyle, default is FocusIntervalStyle
func WithFocusIntervalStyle(s *style.Style) Option {
	return func(c *Confirm) {
		c.FocusIntervalStyle = s
	}
}

// WithUnFocusIntervalStyle replace UnFocusIntervalStyle, default is UnFocusIntervalStyle
func WithUnFocusIntervalStyle(s *style.Style) Option {
	return func(c *Confirm) {
		c.FocusIntervalStyle = s
	}
}

// WithPromptStyle replace PromptStyle, default is PromptStyle
func WithPromptStyle(s *style.Style) Option {
	return func(c *Confirm) {
		c.PromptStyle = s
	}
}

// WithChoiceStyle replace ChoiceStyle, default is ChoiceStyle
func WithChoiceStyle(s *style.Style) Option {
	return func(c *Confirm) {
		c.ChoiceStyle = s
	}
}
