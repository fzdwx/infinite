package style

import "github.com/charmbracelet/lipgloss"

var (
	noBorder = lipgloss.Border{}

	normalBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "└",
		BottomRight: "┘",
	}

	roundedBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}

	thickBorder = lipgloss.Border{
		Top:         "━",
		Bottom:      "━",
		Left:        "┃",
		Right:       "┃",
		TopLeft:     "┏",
		TopRight:    "┓",
		BottomLeft:  "┗",
		BottomRight: "┛",
	}

	doubleBorder = lipgloss.Border{
		Top:         "═",
		Bottom:      "═",
		Left:        "║",
		Right:       "║",
		TopLeft:     "╔",
		TopRight:    "╗",
		BottomLeft:  "╚",
		BottomRight: "╝",
	}

	hiddenBorder = lipgloss.Border{
		Top:         " ",
		Bottom:      " ",
		Left:        " ",
		Right:       " ",
		TopLeft:     " ",
		TopRight:    " ",
		BottomLeft:  " ",
		BottomRight: " ",
	}
)

// NormalBorder returns a standard-type border with a normal weight and 90
// degree corners.
func NormalBorder() lipgloss.Border {
	return normalBorder
}

// RoundedBorder returns a border with rounded corners.
func RoundedBorder() lipgloss.Border {
	return roundedBorder
}

// ThickBorder returns a border that's thicker than the one returned by
// NormalBorder.
func ThickBorder() lipgloss.Border {
	return thickBorder
}

// DoubleBorder returns a border comprised of two thin strokes.
func DoubleBorder() lipgloss.Border {
	return doubleBorder
}

// HiddenBorder returns a border that renders as a series of single-cell
// spaces. It's useful for cases when you want to remove a standard border but
// maintain layout positioning. This said, you can still apply a background
// color to a hidden border.
func HiddenBorder() lipgloss.Border {
	return hiddenBorder
}
