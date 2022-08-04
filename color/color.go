package color

import (
	"github.com/charmbracelet/lipgloss"
	"strconv"
)

var (
	Magenta   = New(5)
	Red       = New(9)
	LightBlue = New(12)
	Blank     = New(30)
	Cyan      = New(49)
	Aqua      = New(86)
	HotPink   = New(201)
	Orange    = New(202)
	RedPink   = New(205)

	FullBlue = NewHex("#0000FF")
	DarkGray = NewHex("#3C3C3C")
	Gray     = NewHex("#808080")

	Special   = NewAdaptive("#43BF6D", "#73F59F")
	Highlight = NewAdaptive("#874BFD", "#7D56F4")
	Subtle    = NewAdaptive("#D9DCCF", "#383838")
)

func New(i int) lipgloss.Color {
	return lipgloss.Color(strconv.Itoa(i))
}

func NewHex(hex string) lipgloss.Color {
	return lipgloss.Color(hex)
}

func NewAdaptive(light, dark string) lipgloss.AdaptiveColor {
	return lipgloss.AdaptiveColor{
		Light: light,
		Dark:  dark,
	}
}
