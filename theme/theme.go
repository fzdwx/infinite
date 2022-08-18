package theme

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/style"
)

type Theme struct {
	PromptStyle                  *style.Style
	MultiSelectedHintSymbolStyle *style.Style
	ChoiceTextStyle              *style.Style
	CursorSymbolStyle            *style.Style
	UnHintSymbolStyle            *style.Style
	SpinnerShapeStyle            *style.Style
	PlaceholderStyle             *style.Style

	FocusSymbol     string
	UnFocusSymbol   string
	FocusInterval   string
	UnFocusInterval string

	FocusSymbolStyle     *style.Style
	UnFocusSymbolStyle   *style.Style
	FocusIntervalStyle   *style.Style
	UnFocusIntervalStyle *style.Style
}

var (
	DefaultTheme = Theme{
		PromptStyle:                  style.New().Fg(color.Cyan),
		MultiSelectedHintSymbolStyle: style.New().Fg(color.Special),
		ChoiceTextStyle:              style.New().Fg(color.Highlight).Bold(),
		CursorSymbolStyle:            style.New(),
		UnHintSymbolStyle:            style.New().Fg(color.Red),
		SpinnerShapeStyle:            style.New().Fg(color.RedPink),
		PlaceholderStyle:             style.New().Fg(lipgloss.Color("240")),
		FocusSymbol:                  "? ",
		UnFocusSymbol:                "√ ",
		FocusInterval:                " » ",
		UnFocusInterval:              " … ",
		FocusSymbolStyle:             style.New().Fg(color.Cyan),
		UnFocusSymbolStyle:           style.New().Fg(color.Green),
		FocusIntervalStyle:           style.New().Fg(color.Gray),
		UnFocusIntervalStyle:         style.New().Fg(color.Gray).Bold(),
	}

	// fix https://github.com/fzdwx/infinite/issues/5
	_ = DefaultTheme.PromptStyle.Render("123")
	_ = DefaultTheme.MultiSelectedHintSymbolStyle.Render("123")
	_ = DefaultTheme.ChoiceTextStyle.Render("123")
	_ = DefaultTheme.CursorSymbolStyle.Render("123")
	_ = DefaultTheme.UnHintSymbolStyle.Render("123")
	_ = DefaultTheme.SpinnerShapeStyle.Render("123")
	_ = DefaultTheme.PlaceholderStyle.Render("123")
)
