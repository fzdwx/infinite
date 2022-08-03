package theme

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/style"
)

type Theme struct {
	PromptStyle                  lipgloss.Style
	MultiSelectedHintSymbolStyle lipgloss.Style
	UnHintSymbolStyle            lipgloss.Style
}

var (
	DefaultTheme = Theme{
		PromptStyle:                  style.New().Foreground(color.Cyan),
		MultiSelectedHintSymbolStyle: style.New().Foreground(color.Special),
		UnHintSymbolStyle:            style.New().Foreground(color.Red),
	}
)
