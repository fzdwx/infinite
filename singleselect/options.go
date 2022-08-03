package singleselect

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/multiselect"
)

// Option the option of Select
type Option func(s *Select)

// WithRowRender default is
//
// `
// fmt.Sprintf("%s [%s] %s", cursorSymbol, hintSymbol, choice)
// `
func WithRowRender(rowRender func(string, string, string) string) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithRowRender(rowRender))
	}
}

// WithPageSize default is 5
func WithPageSize(pageSize int) Option {
	return func(ss *Select) {
		ss.inner.Apply(multiselect.WithPageSize(pageSize))
	}
}

// WithKeyBinding replace key map.
func WithKeyBinding(keymap KeyMap) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithKeyBinding(keymap.MapToMulti()))
	}
}

// WithCursorSymbol default is ">"
func WithCursorSymbol(symbol string) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithCursorSymbol(symbol))
	}
}

// WithCursorSymbolStyle default is theme.DefaultTheme.CursorSymbolStyle.
func WithCursorSymbolStyle(style lipgloss.Style) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithCursorSymbolStyle(style))
	}
}

// WithChoiceTextStyle default is theme.DefaultTheme.ChoiceTextStyle.
func WithChoiceTextStyle(style lipgloss.Style) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithChoiceTextStyle(style))
	}
}

// WithDisableOutputResult disable output result.
func WithDisableOutputResult() Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithDisableOutputResult())
	}
}

// WithPromptStyle default is Theme.PromptStyle.
func WithPromptStyle(style lipgloss.Style) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithPromptStyle(style))
	}
}

// WithPrompt default is "Please select your options:"
func WithPrompt(prompt ...string) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithPrompt(prompt...))
	}
}
