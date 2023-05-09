package confirm

import (
	"github.com/fzdwx/infinite/style"
	"github.com/fzdwx/infinite/theme"
)

type Style struct {
	SymbolStyle *style.Style
	Symbol      string

	IntervalStyle *style.Style
	Interval      string

	NoticeStyle *style.Style
	Notice      string

	PromptStyle *style.Style
}

func resetStyle(s *Style) *Style {
	s.SymbolStyle = style.New()
	s.IntervalStyle = style.New()
	s.NoticeStyle = style.New()
	s.PromptStyle = style.New()
	return s
}

func FocusStyle() *Style {
	return &Style{
		Symbol:        theme.DefaultTheme.FocusSymbol,
		Interval:      theme.DefaultTheme.FocusInterval,
		SymbolStyle:   theme.DefaultTheme.FocusSymbolStyle,
		IntervalStyle: theme.DefaultTheme.FocusIntervalStyle,
		PromptStyle:   theme.DefaultTheme.PromptStyle,
		NoticeStyle:   style.New(),
		Notice:        " ( y/N )",
	}
}

func UnFocusStyle() *Style {
	return &Style{
		Symbol:        theme.DefaultTheme.UnFocusSymbol,
		Interval:      theme.DefaultTheme.UnFocusInterval,
		SymbolStyle:   theme.DefaultTheme.UnFocusSymbolStyle,
		IntervalStyle: theme.DefaultTheme.UnFocusIntervalStyle,
		PromptStyle:   theme.DefaultTheme.PromptStyle,
		NoticeStyle:   style.New(),
		Notice:        " ( y/N )",
	}
}
