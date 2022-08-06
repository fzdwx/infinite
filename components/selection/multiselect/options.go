package multiselect

import (
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/selection"
	"github.com/fzdwx/infinite/style"
)

// Option the option of Select
type Option func(s *Select)

// WithDisableFilter disable filter.
func WithDisableFilter() Option {
	return func(s *Select) {
		s.inner.EnableFilter = false
	}
}

// WithFilterInput replace filter input.
func WithFilterInput(input *components.Input) Option {
	return func(s *Select) {
		s.inner.FilterInput = input
	}
}

// WithFilterFunc replace filter func.
func WithFilterFunc(f func(input string, items []components.SelectionItem) []components.SelectionItem) Option {
	return func(s *Select) {
		s.inner.FilterFunc = f
	}
}

// WithRowRender default is
//
// `
// fmt.Sprintf("%s [%s] %s", cursorSymbol, hintSymbol, choice)
// `
func WithRowRender(rowRender func(string, string, string) string) Option {
	return func(s *Select) {
		s.inner.RowRender = rowRender
	}
}

// WithPageSize default is 5
func WithPageSize(pageSize int) Option {
	return func(s *Select) {
		s.inner.PageSize = pageSize
	}
}

// WithKeyBinding replace key map.
func WithKeyBinding(keymap selection.KeyMap) Option {
	return func(s *Select) {
		s.inner.Keymap = keymap
	}
}

// WithCursorSymbol default is ">"
func WithCursorSymbol(symbol string) Option {
	return func(s *Select) {
		s.inner.CursorSymbol = symbol
	}
}

// WithCursorSymbolStyle default is theme.DefaultTheme.CursorSymbolStyle.
func WithCursorSymbolStyle(style *style.Style) Option {
	return func(s *Select) {
		s.inner.CursorSymbolStyle = style
	}
}

// WithChoiceTextStyle default is theme.DefaultTheme.ChoiceTextStyle.
func WithChoiceTextStyle(style *style.Style) Option {
	return func(s *Select) {
		s.inner.ChoiceTextStyle = style
	}
}

// WithHintSymbol default is "✓".
func WithHintSymbol(selectedStr string) Option {
	return func(s *Select) {
		s.inner.HintSymbol = selectedStr
	}
}

// WithHintSymbolStyle default is Theme.MultiSelectedHintSymbolStyle.
func WithHintSymbolStyle(style *style.Style) Option {
	return func(s *Select) {
		s.inner.HintSymbolStyle = style
	}
}

// WithUnHintSymbol default is "✗".
func WithUnHintSymbol(unSelectedStr string) Option {
	return func(s *Select) {
		s.inner.UnHintSymbol = unSelectedStr
	}
}

// WithUnHintSymbolStyle default is Theme.UnHintSymbolStyle.
func WithUnHintSymbolStyle(style *style.Style) Option {
	return func(s *Select) {
		s.inner.UnHintSymbolStyle = style
	}
}

// WithPromptStyle default is Theme.PromptStyle.
func WithPromptStyle(style *style.Style) Option {
	return func(s *Select) {
		s.inner.PromptStyle = style
	}
}

// WithPrompt default is "Please selection your options:"
func WithPrompt(prompt ...string) Option {
	return func(s *Select) {
		if len(prompt) >= 1 && len(prompt[0]) > 0 {
			s.inner.Prompt = prompt[0]
		}
	}
}

// WithDisableOutputResult disable output result.
func WithDisableOutputResult() Option {
	return func(s *Select) {
		s.inner.DisableOutPutResult = true
	}
}
