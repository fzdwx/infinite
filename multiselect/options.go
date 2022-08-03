package multiselect

import "github.com/charmbracelet/lipgloss"

// Option the option of MultiSelect
type Option func(ms *MultiSelect)

// WithRowRender default is
//
// `
// fmt.Sprintf("%s [%s] %s", cursorSymbol, hintSymbol, choice)
// `
func WithRowRender(rowRender func(string, string, string) string) Option {
	return func(ms *MultiSelect) {
		ms.inner.rowRender = rowRender
	}
}

// WithPageSize default is 5
func WithPageSize(pageSize int) Option {
	return func(ms *MultiSelect) {
		ms.inner.pageSize = pageSize
	}
}

// WithKeyBinding replace key map.
func WithKeyBinding(keymap KeyMap) Option {
	return func(ms *MultiSelect) {
		ms.inner.keymap = keymap
	}
}

// WithCursorSymbol default is ">"
func WithCursorSymbol(symbol string) Option {
	return func(ms *MultiSelect) {
		ms.inner.cursorSymbol = symbol
	}
}

// WithCursorSymbolStyle default is theme.DefaultTheme.CursorSymbolStyle.
func WithCursorSymbolStyle(style lipgloss.Style) Option {
	return func(ms *MultiSelect) {
		ms.inner.cursorSymbolStyle = style
	}
}

// WithChoiceTextStyle default is theme.DefaultTheme.ChoiceTextStyle.
func WithChoiceTextStyle(style lipgloss.Style) Option {
	return func(ms *MultiSelect) {
		ms.inner.choiceTextStyle = style
	}
}

// WithHintSymbol default is "✓".
func WithHintSymbol(selectedStr string) Option {
	return func(ms *MultiSelect) {
		ms.inner.hintSymbol = selectedStr
	}
}

// WithHintSymbolStyle default is Theme.MultiSelectedHintSymbolStyle.
func WithHintSymbolStyle(style lipgloss.Style) Option {
	return func(ms *MultiSelect) {
		ms.inner.hintSymbolStyle = style
	}
}

// WithUnHintSymbol default is "✗".
func WithUnHintSymbol(unSelectedStr string) Option {
	return func(ms *MultiSelect) {
		ms.inner.unHintSymbol = unSelectedStr
	}
}

// WithUnHintSymbolStyle default is Theme.UnHintSymbolStyle.
func WithUnHintSymbolStyle(style lipgloss.Style) Option {
	return func(ms *MultiSelect) {
		ms.inner.unHintSymbolStyle = style
	}
}

// WithPromptStyle default is Theme.PromptStyle.
func WithPromptStyle(style lipgloss.Style) Option {
	return func(ms *MultiSelect) {
		ms.inner.promptStyle = style
	}
}

// WithDisableOutputResult disable output result.
func WithDisableOutputResult() Option {
	return func(ms *MultiSelect) {
		ms.inner.disableOutPutResult = true
	}
}

// WithPrompt default is "Please select your options:"
func WithPrompt(prompt ...string) Option {
	return func(ms *MultiSelect) {
		if len(prompt) >= 1 && len(prompt[0]) > 0 {
			ms.inner.prompt = prompt[0]
		}
	}
}
