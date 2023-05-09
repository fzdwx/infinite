package confirm

import (
	"github.com/fzdwx/infinite/style"
)

type Option func(confirm *Confirm)

// WithPure do not use any beautification features,
// any options you customize will be cleared
func WithPure() Option {
	return func(i *Confirm) {
		i.inner.FocusStyle = resetStyle(i.inner.FocusStyle)
		i.inner.UnFocusStyle = resetStyle(i.inner.UnFocusStyle)
		i.inner.ValueStyle = style.New()
	}
}

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
// Deprecated: use WithFocusStyle or WithUnFocusStyle instead.
func WithPromptStyle(style *style.Style) Option {
	return func(c *Confirm) {
		c.inner.FocusStyle.PromptStyle = style
		c.inner.UnFocusStyle.PromptStyle = style
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
// Deprecated: use WithFocusStyle or WithUnFocusStyle instead.
func WithNotice(notice string) Option {
	return func(c *Confirm) {
		c.inner.FocusStyle.Notice = notice
		c.inner.UnFocusStyle.Notice = notice
	}
}

// WithNoticeStyle replace notice style.
// Deprecated: use WithFocusStyle or WithUnFocusStyle instead.
func WithNoticeStyle(style *style.Style) Option {
	return func(c *Confirm) {
		c.inner.FocusStyle.NoticeStyle = style
		c.inner.UnFocusStyle.NoticeStyle = style
	}
}

// WithFocusSymbol default is FocusStyle#Symbol
// Deprecated: use WithFocusStyle or WithUnFocusStyle instead.
func WithFocusSymbol(s string) Option {
	return func(confirm *Confirm) {
		confirm.inner.FocusStyle.Symbol = s
	}
}

// WithUnFocusSymbol default is UnFocusStyle#Symbol
// Deprecated: use WithFocusStyle or WithUnFocusStyle instead.
func WithUnFocusSymbol(s string) Option {
	return func(confirm *Confirm) {
		confirm.inner.UnFocusStyle.Symbol = s
	}
}

// WithFocusInterval default FocusStyle#Interval
// Deprecated: use WithFocusStyle or WithUnFocusStyle instead.
func WithFocusInterval(s string) Option {
	return func(confirm *Confirm) {
		confirm.inner.FocusStyle.Interval = s
	}
}

// WithUnFocusInterval default is UnFocusStyle#Interval
// Deprecated: use WithFocusStyle or WithUnFocusStyle instead.
func WithUnFocusInterval(s string) Option {
	return func(confirm *Confirm) {
		confirm.inner.UnFocusStyle.Interval = s
	}
}

// WithFocusSymbolStyle default is FocusStyle#SymbolStyle
// Deprecated: use WithFocusStyle or WithUnFocusStyle instead.
func WithFocusSymbolStyle(s *style.Style) Option {
	return func(confirm *Confirm) {
		confirm.inner.FocusStyle.SymbolStyle = s
	}
}

// WithUnFocusSymbolStyle default is UnFocusStyle#SymbolStyle
// Deprecated: use WithFocusStyle or WithUnFocusStyle instead.
func WithUnFocusSymbolStyle(s *style.Style) Option {
	return func(confirm *Confirm) {
		confirm.inner.UnFocusStyle.SymbolStyle = s
	}
}

// WithFocusIntervalStyle default is FocusStyle#IntervalStyle
// Deprecated: use WithFocusStyle or WithUnFocusStyle instead.
func WithFocusIntervalStyle(s *style.Style) Option {
	return func(confirm *Confirm) {
		confirm.inner.FocusStyle.IntervalStyle = s
	}
}

// WithUnFocusIntervalStyle default is UnFocusStyle#IntervalStyle
// Deprecated: use WithFocusStyle or WithUnFocusStyle instead.
func WithUnFocusIntervalStyle(s *style.Style) Option {
	return func(confirm *Confirm) {
		confirm.inner.UnFocusStyle.IntervalStyle = s
	}
}

func WithFocusStyle(s *Style) Option {
	return func(confirm *Confirm) {
		confirm.inner.FocusStyle = s
	}
}

func WithUnFocusStyle(s *Style) Option {
	return func(confirm *Confirm) {
		confirm.inner.UnFocusStyle = s
	}
}
