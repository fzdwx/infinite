package multiselect

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/fzdwx/infinite/components"
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

// WithKeyMap replace key map.
// see components.DefaultMultiKeyMap
//
// example:
//
// keymap := components.DefaultMultiKeyMap()
//
// keymap.Choice = key.NewBinding(
//
//	key.WithKeys(tea.KeySpace.String()),
//
// )
func WithKeyMap(keymap components.SelectionKeyMap) Option {
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

// WithHeader default is ""
func WithHeader(header string) Option {
	return func(s *Select) {
		s.inner.Header = header
	}
}

// WithDisableOutputResult disable output result.
func WithDisableOutputResult() Option {
	return func(s *Select) {
		s.inner.DisableOutPutResult = true
	}
}

// WithFocusSymbol default is theme.DefaultTheme#FocusSymbol:
func WithFocusSymbol(ss string) Option {
	return func(s *Select) {
		s.inner.FocusSymbol = ss
	}
}

// WithUnFocusSymbol default is theme.DefaultTheme#UnFocusSymbol:
func WithUnFocusSymbol(ss string) Option {
	return func(s *Select) {
		s.inner.UnFocusSymbol = ss
	}
}

// WithFocusInterval default is theme.DefaultTheme#FocusInterval:
func WithFocusInterval(ss string) Option {
	return func(s *Select) {
		s.inner.FocusInterval = ss
	}
}

// WithUnFocusInterval default is theme.DefaultTheme#UnFocusInterval:
func WithUnFocusInterval(ss string) Option {
	return func(s *Select) {
		s.inner.UnFocusInterval = ss
	}
}

// WithFocusSymbolStyle default is theme.DefaultTheme#FocusSymbolStyle:
func WithFocusSymbolStyle(stl *style.Style) Option {
	return func(s *Select) {
		s.inner.FocusSymbolStyle = stl
	}
}

// WithUnFocusSymbolStyle default is theme.DefaultTheme#UnFocusSymbolStyle:
func WithUnFocusSymbolStyle(stl *style.Style) Option {
	return func(s *Select) {
		s.inner.UnFocusSymbolStyle = stl
	}
}

// WithFocusIntervalStyle default is theme.DefaultTheme#FocusIntervalStyle:
func WithFocusIntervalStyle(stl *style.Style) Option {
	return func(s *Select) {
		s.inner.FocusIntervalStyle = stl
	}
}

// WithUnFocusIntervalStyle default is theme.DefaultTheme#UnFocusIntervalStyle:
func WithUnFocusIntervalStyle(stl *style.Style) Option {
	return func(s *Select) {
		s.inner.UnFocusIntervalStyle = stl
	}
}

// WithValueStyle default is theme.DefaultTheme#ChoiceTextStyle.Underline()
func WithValueStyle(stl *style.Style) Option {
	return func(s *Select) {
		s.inner.ValueStyle = stl
	}
}

// WithValidator specifies a validator to use while prompting the user
func WithValidator(v components.Validator) Option {
	return func(s *Select) {
		// add the provided validator to the list
		s.inner.Validators = append(s.inner.Validators, v)
	}
}

// WithDisableShowHelp disable show help.
func WithDisableShowHelp() Option {
	return func(s *Select) {
		s.inner.ShowHelp = false
	}
}

// WithPaginator set paginator.
func WithPaginator(pager paginator.Model) Option {
	return func(s *Select) {
		s.inner.Paginator = pager
	}
}

// WithHiddenPaginator hidden paginator view.
func WithHiddenPaginator() Option {
	return func(s *Select) {
		s.inner.ShowPaginator = false
	}
}
