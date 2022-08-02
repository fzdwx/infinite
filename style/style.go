package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/color"
)

func New() lipgloss.Style {
	return lipgloss.NewStyle()
}

var (
	DefaultStyle = New().Foreground(color.DefaultFg).Background(color.DefaultBg)
	PrimaryStyle = New().Foreground(color.DefaultFg)
)
