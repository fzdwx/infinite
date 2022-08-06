package singleselect

import (
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/selection"
	"github.com/fzdwx/infinite/components/selection/multiselect"
	"github.com/fzdwx/infinite/style"
)

// Option the option of Select
type Option func(s *Select)

// WithDisableFilter disable filter.
func WithDisableFilter() Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithDisableFilter())
	}
}

// WithFilterInput replace filter input.
func WithFilterInput(input *components.Input) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithFilterInput(input))
	}
}

// WithFilterFunc replace filter func.
func WithFilterFunc(f func(input string, items []components.SelectionItem) []components.SelectionItem) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithFilterFunc(f))
	}
}

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
func WithKeyBinding(keymap selection.KeyMap) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithKeyBinding(keymap))
	}
}

// WithCursorSymbol default is ">"
func WithCursorSymbol(symbol string) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithCursorSymbol(symbol))
	}
}

// WithCursorSymbolStyle default is theme.DefaultTheme.CursorSymbolStyle.
func WithCursorSymbolStyle(style *style.Style) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithCursorSymbolStyle(style))
	}
}

// WithChoiceTextStyle default is theme.DefaultTheme.ChoiceTextStyle.
func WithChoiceTextStyle(style *style.Style) Option {
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
func WithPromptStyle(style *style.Style) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithPromptStyle(style))
	}
}

// WithPrompt default is "Please selection your options:"
func WithPrompt(prompt ...string) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithPrompt(prompt...))
	}
}
