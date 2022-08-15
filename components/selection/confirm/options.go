package confirm

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/style"
	"github.com/fzdwx/infinite/theme"
)

var (
	Yes             = "Yes"
	No              = "No"
	Symbol          = "? "
	Prompt          = "Are you handsome?"
	FocusInterval   = style.New().Fg(color.Gray).Render(" » ")
	UnFocusInterval = style.New().Fg(color.White).Bold().Render(" ... ")
	SymbolStyle     = style.New().Fg(color.Special)
	PromptStyle     = style.New().Bold()
	ChoiceStyle     = theme.DefaultTheme.ChoiceTextStyle.Underline()
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
			key.WithHelp("enter", "choice item"),
		),
	}
}

type Option func(c *Confirm)

// WithDefaultYes  set Confirm default val is `true`
func WithDefaultYes() Option {
	return func(c *Confirm) {
		c.DefaultVal = true
	}
}

func WithDisableOutputResult() Option {
	return func(c *Confirm) {
		c.OutputResult = false
	}
}

func WithKeyMap(keymap KeyMap) Option {
	return func(c *Confirm) {
		c.KeyMap = keymap
	}
}

func WithYes(yes string) Option {
	return func(c *Confirm) {
		c.Yes = yes
	}
}

func WithNo(no string) Option {
	return func(c *Confirm) {
		c.No = no
	}
}

func WithPrompt(prompt string) Option {
	return func(c *Confirm) {
		c.Prompt = prompt
	}
}

func WithSymbol(s string) Option {
	return func(c *Confirm) {
		c.Symbol = s
	}
}

func WithFocusInterval(no string) Option {
	return func(c *Confirm) {
		c.FocusInterval = no
	}
}

func WithUnFocusInterval(no string) Option {
	return func(c *Confirm) {
		c.UnFocusInterval = no
	}
}

func WithSymbolStyle(s *style.Style) Option {
	return func(c *Confirm) {
		c.SymbolStyle = s
	}
}

func WithPromptStyle(s *style.Style) Option {
	return func(c *Confirm) {
		c.PromptStyle = s
	}
}

func WithChoiceStyle(s *style.Style) Option {
	return func(c *Confirm) {
		c.ChoiceStyle = s
	}
}
