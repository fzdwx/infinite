package singleselect

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/fzdwx/infinite/components"
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

// WithDefaultFiltering default in filtering
func WithDefaultFiltering() Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithDefaultFiltering())
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

// WithHelpModel replace help model
// default is help.New()
func WithHelpModel(h help.Model) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithHelpModel(h))
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
// see DefaultSingleKeyMap
func WithKeyBinding(keymap KeyMap) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithKeyMap(
			components.SelectionKeyMap{
				Up:           keymap.Up,
				Down:         keymap.Down,
				Choice:       keymap.Choice,
				Confirm:      keymap.Confirm,
				Quit:         keymap.Quit,
				ToggleFilter: keymap.ToggleFilter,
				NextPage:     keymap.NextPage,
				PrevPage:     keymap.PrevPage,
			}))
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

// WithFocusSymbol default is theme.DefaultTheme#FocusSymbol:
func WithFocusSymbol(ss string) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithFocusSymbol(ss))
	}
}

// WithUnFocusSymbol default is theme.DefaultTheme#UnFocusSymbol:
func WithUnFocusSymbol(ss string) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithUnFocusSymbol(ss))
	}
}

// WithFocusInterval default is theme.DefaultTheme#FocusInterval:
func WithFocusInterval(ss string) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithFocusInterval(ss))
	}
}

// WithUnFocusInterval default is theme.DefaultTheme#UnFocusInterval:
func WithUnFocusInterval(ss string) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithUnFocusInterval(ss))
	}
}

// WithFocusSymbolStyle default is theme.DefaultTheme#FocusSymbolStyle:
func WithFocusSymbolStyle(stl *style.Style) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithFocusSymbolStyle(stl))
	}
}

// WithUnFocusSymbolStyle default is theme.DefaultTheme#UnFocusSymbolStyle:
func WithUnFocusSymbolStyle(stl *style.Style) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithUnFocusSymbolStyle(stl))
	}
}

// WithFocusIntervalStyle default is theme.DefaultTheme#FocusIntervalStyle:
func WithFocusIntervalStyle(stl *style.Style) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithFocusIntervalStyle(stl))
	}
}

// WithUnFocusIntervalStyle default is theme.DefaultTheme#UnFocusIntervalStyle:
func WithUnFocusIntervalStyle(stl *style.Style) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithUnFocusIntervalStyle(stl))
	}
}

// WithValueStyle default is theme.DefaultTheme#ChoiceTextStyle.Underline()
func WithValueStyle(stl *style.Style) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithValueStyle(stl))
	}
}

// WithDisableHelp disable show help.
func WithDisableHelp() Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithDisableShowHelp())
	}
}

// WithPaginator set paginator.
func WithPaginator(pager paginator.Model) Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithPaginator(pager))
	}
}

// WithHiddenPaginator hidden paginator view.
func WithHiddenPaginator() Option {
	return func(s *Select) {
		s.inner.Apply(multiselect.WithHiddenPaginator())
	}
}
