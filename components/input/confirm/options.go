package confirm

import (
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

// WithDisableOutputResult disable output result
func WithDisableOutputResult() Option {
	return func(confirm *Confirm) {
		confirm.inner.OutputResult = false
	}
}

// WithValueStyle render value
func WithValueStyle(s *style.Style) Option {
	return func(c *Confirm) {
		c.inner.ValueStyle = s
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
		c.inner.PromptStyle = style
	}
}

// WithKeyMap change `confirm` default KeyMap.
func WithKeyMap(keyMap KeyMap) Option {
	return func(c *Confirm) {
		c.inner.KeyMap = keyMap
	}
}

// WithDisplayHelp display help view.
func WithDisplayHelp() Option {
	return func(c *Confirm) {
		c.inner.DisplayHelp = true
	}
}

// WithNotice replace notice, default is  " ( y/N ) ".
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

// WithFocusSymbol default is theme.DefaultTheme#FocusSymbol
func WithFocusSymbol(s string) Option {
	return func(confirm *Confirm) {
		confirm.inner.FocusSymbol = s
	}
}

// WithUnFocusSymbol default is theme.DefaultTheme#UnFocusSymbol
func WithUnFocusSymbol(s string) Option {
	return func(confirm *Confirm) {
		confirm.inner.UnFocusSymbol = s
	}
}

// WithFocusInterval default is theme.DefaultTheme#FocusInterval
func WithFocusInterval(s string) Option {
	return func(confirm *Confirm) {
		confirm.inner.FocusInterval = s
	}
}

// WithUnFocusInterval default is theme.DefaultTheme#UnFocusInterval
func WithUnFocusInterval(s string) Option {
	return func(confirm *Confirm) {
		confirm.inner.UnFocusInterval = s
	}
}

// WithFocusSymbolStyle default is theme.DefaultTheme#FocusSymbolStyle
func WithFocusSymbolStyle(s *style.Style) Option {
	return func(confirm *Confirm) {
		confirm.inner.FocusSymbolStyle = s
	}
}

// WithUnFocusSymbolStyle default is theme.DefaultTheme#UnFocusIntervalStyle
func WithUnFocusSymbolStyle(s *style.Style) Option {
	return func(confirm *Confirm) {
		confirm.inner.UnFocusSymbolStyle = s
	}
}

// WithFocusIntervalStyle default is theme.DefaultTheme#FocusIntervalStyle
func WithFocusIntervalStyle(s *style.Style) Option {
	return func(confirm *Confirm) {
		confirm.inner.FocusIntervalStyle = s
	}
}

// WithUnFocusIntervalStyle default is theme.DefaultTheme#UnFocusIntervalStyle
func WithUnFocusIntervalStyle(s *style.Style) Option {
	return func(confirm *Confirm) {
		confirm.inner.UnFocusIntervalStyle = s
	}
}
