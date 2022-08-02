package color

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"strconv"
)

var (
	Color = termenv.ColorProfile()

	DefaultFg   = New(39)
	DefaultBg   = New(49)
	FgLightCyan = New(96)
)

func New(i int) lipgloss.Color {
	return lipgloss.Color(strconv.Itoa(i))
}

func NewHex(hex string) lipgloss.Color {
	return lipgloss.Color(hex)
}
