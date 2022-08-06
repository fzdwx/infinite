package confirm

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/fzdwx/infinite/style"
)

type Option func(confirm *Confirm)

// WithDefaultYes the `confirm` default is no,
// adding this option will turn into yes.
func WithDefaultYes() Option {
	return func(c *Confirm) {
		c.inner.Value = true
	}
}

// WithPrompt change `confirm` default prompt.
func WithPrompt(prompt string) Option {
	return func(c *Confirm) {
		c.inner.input.Prompt = prompt
	}
}

// WithPromptStyle change `confirm` default promptStyle.
func WithPromptStyle(style *style.Style) Option {
	return func(c *Confirm) {
		c.inner.input.PromptStyle = style
	}
}

// WithKeyMap change `confirm` default KeyMap.
func WithKeyMap(keyMap KeyMap) Option {
	return func(c *Confirm) {
		c.inner.KeyMap = keyMap
	}
}

// WithHelp replace help model.
func WithHelp(help help.Model) Option {
	return func(c *Confirm) {
		c.inner.Help = help
	}
}

// WithDisplayHelp display help view.
func WithDisplayHelp() Option {
	return func(c *Confirm) {
		c.inner.DisplayHelp = true
	}
}

// WithNotice replace notice, default is  " ( y/N ):".
func WithNotice(notice string) Option {
	return func(c *Confirm) {
		c.inner.Notice = notice
	}
}

// WithNoticeStyle replace notice style.
func WithNoticeStyle(style *style.Style) Option {
	return func(c *Confirm) {
		c.inner.NoticeStyle = style
	}
}

// WithSymbol replace symbol, default is "?".
func WithSymbol(symbol string) Option {
	return func(c *Confirm) {
		c.inner.Symbol = symbol
	}
}

// WithSymbolStyle replace symbol style.
func WithSymbolStyle(style *style.Style) Option {
	return func(c *Confirm) {
		c.inner.SymbolStyle = style
	}
}
