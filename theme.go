package inf

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/style"
)

type theme struct {
	primaryStyle lipgloss.Style

	multiSelectedStrStyle lipgloss.Style
	unSelectedStrStyle    lipgloss.Style
}

var (
	Theme = theme{
		primaryStyle:          style.New().Foreground(color.Cyan),
		multiSelectedStrStyle: style.New().Foreground(color.Special),
		unSelectedStrStyle:    style.New().Foreground(color.Red),
	}
)
