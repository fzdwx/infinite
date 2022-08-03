package multiselect

import "github.com/charmbracelet/lipgloss"

// Option the option of Select
type Option func(s *Select)

// WithRowRender default is
//
// `
// fmt.Sprintf("%s [%s] %s", cursorSymbol, hintSymbol, choice)
// `
func WithRowRender(rowRender func(string, string, string) string) Option {
	return func(s *Select) {
		s.inner.rowRender = rowRender
	}
}

// WithPageSize default is 5
func WithPageSize(pageSize int) Option {
	return func(s *Select) {
		s.inner.pageSize = pageSize
	}
}

// WithKeyBinding replace key map.
func WithKeyBinding(keymap KeyMap) Option {
	return func(s *Select) {
		s.inner.keymap = keymap
	}
}

// WithCursorSymbol default is ">"
func WithCursorSymbol(symbol string) Option {
	return func(s *Select) {
		s.inner.cursorSymbol = symbol
	}
}

// WithCursorSymbolStyle default is theme.DefaultTheme.CursorSymbolStyle.
func WithCursorSymbolStyle(style lipgloss.Style) Option {
	return func(s *Select) {
		s.inner.cursorSymbolStyle = style
	}
}

// WithChoiceTextStyle default is theme.DefaultTheme.ChoiceTextStyle.
func WithChoiceTextStyle(style lipgloss.Style) Option {
	return func(s *Select) {
		s.inner.choiceTextStyle = style
	}
}

// WithHintSymbol default is "✓".
func WithHintSymbol(selectedStr string) Option {
	return func(s *Select) {
		s.inner.hintSymbol = selectedStr
	}
}

// WithHintSymbolStyle default is Theme.MultiSelectedHintSymbolStyle.
func WithHintSymbolStyle(style lipgloss.Style) Option {
	return func(s *Select) {
		s.inner.hintSymbolStyle = style
	}
}

// WithUnHintSymbol default is "✗".
func WithUnHintSymbol(unSelectedStr string) Option {
	return func(s *Select) {
		s.inner.unHintSymbol = unSelectedStr
	}
}

// WithUnHintSymbolStyle default is Theme.UnHintSymbolStyle.
func WithUnHintSymbolStyle(style lipgloss.Style) Option {
	return func(s *Select) {
		s.inner.unHintSymbolStyle = style
	}
}

// WithPromptStyle default is Theme.PromptStyle.
func WithPromptStyle(style lipgloss.Style) Option {
	return func(s *Select) {
		s.inner.promptStyle = style
	}
}

// WithPrompt default is "Please select your options:"
func WithPrompt(prompt ...string) Option {
	return func(s *Select) {
		if len(prompt) >= 1 && len(prompt[0]) > 0 {
			s.inner.prompt = prompt[0]
		}
	}
}

// WithDisableOutputResult disable output result.
func WithDisableOutputResult() Option {
	return func(s *Select) {
		s.inner.disableOutPutResult = true
	}
}
