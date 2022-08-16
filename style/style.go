package style

import (
	"github.com/charmbracelet/lipgloss"
)

// New style
func New() *Style {
	style := lipgloss.NewStyle()
	return &Style{inner: style}
}

// Style wrapper lipgloss.Style
type Style struct {
	inner lipgloss.Style
}

// Center set the horizontal position to the center
func (s *Style) Center() *Style {
	s.inner = s.inner.Align(lipgloss.Center)
	return s
}

// Left set the horizontal position to the left
func (s *Style) Left() *Style {
	s.inner = s.inner.Align(lipgloss.Left)
	return s
}

// Right set the horizontal position to the right
func (s *Style) Right() *Style {
	s.inner = s.inner.Align(lipgloss.Right)
	return s
}

// Top set vertical position to top
func (s *Style) Top() *Style {
	s.inner = s.inner.Align(lipgloss.Center)
	return s
}

// Bottom set vertical position to bottom
func (s *Style) Bottom() *Style {
	s.inner = s.inner.Align(lipgloss.Bottom)
	return s
}

// Bold sets a bold formatting rule.
func (s *Style) Bold() *Style {
	s.inner = s.inner.Bold(true)
	return s
}

// Italic sets an italic formatting rule. In some terminal emulators this will
// render with "reverse" coloring if not italic font variant is available.
func (s *Style) Italic() *Style {
	s.inner = s.inner.Italic(true)
	return s
}

// Inline makes rendering output one line and disables the rendering of
// margins, padding and borders. This is useful when you need a style to apply
// only to font rendering and don't want it to change any physical dimensions.
// It works well with Style.MaxWidth.
//
// Because this in intended to be used at the time of render, this method will
// not mutate the style and instead return a copy.
//
// Example:
//
//	var userInput string = "..."
//	var userStyle = text.Style{ /* ... */ }
//	fmt.Println(userStyle.Inline(true).Render(userInput))
func (s *Style) Inline() *Style {
	s.inner = s.inner.Inline(true)
	return s
}

// Underline sets an underline rule. By default, underlines will not be drawn on
// whitespace like margins and padding. To change this behavior set
// renderUnderlinesOnSpaces.
func (s *Style) Underline() *Style {
	s.inner = s.inner.Underline(true)
	return s
}

// Reverse sets a rule for inverting foreground and background colors.
func (s *Style) Reverse() *Style {
	s.inner = s.inner.Reverse(true)
	return s
}

// Strikethrough sets a strikethrough rule. By default, strikes will not be
// drawn on whitespace like margins and padding. To change this behavior set
// renderStrikethroughOnSpaces.
func (s *Style) Strikethrough() *Style {
	s.inner = s.inner.Strikethrough(true)
	return s
}

// Blink sets a rule for blinking foreground text.
func (s *Style) Blink() *Style {
	s.inner = s.inner.Blink(true)
	return s
}

// Faint sets a rule for rendering the foreground color in a dimmer shade.
func (s *Style) Faint() *Style {
	s.inner = s.inner.Faint(true)
	return s
}

// Width sets the width of the block before applying margins. The width, if
// set, also determines where text will wrap.
func (s *Style) Width(i int) *Style {
	s.inner = s.inner.Width(i)
	return s
}

// Height sets the width of the block before applying margins. If the height of
// the text block is less than this value after applying padding (or not), the
// block will be set to this height.
func (s *Style) Height(i int) *Style {
	s.inner = s.inner.Height(i)
	return s
}

// Padding is a shorthand method for setting padding on all sides at once.
//
// With one argument, the value is applied to all sides.
//
// With two arguments, the value is applied to the vertical and horizontal
// sides, in that order.
//
// With three arguments, the value is applied to the top side, the horizontal
// sides, and the bottom side, in that order.
//
// With four arguments, the value is applied clockwise starting from the top
// side, followed by the right side, then the bottom, and finally the left.
//
// With more than four arguments no padding will be added.
func (s *Style) Padding(i ...int) *Style {
	s.inner = s.inner.Padding(i...)
	return s
}

// PaddingLeft adds padding on the left.
func (s *Style) PaddingLeft(i int) *Style {
	s.inner = s.inner.PaddingLeft(i)
	return s
}

// PaddingRight adds padding on the right.
func (s *Style) PaddingRight(i int) *Style {
	s.inner = s.inner.PaddingRight(i)
	return s
}

// PaddingTop adds padding to the top of the block.
func (s *Style) PaddingTop(i int) *Style {
	s.inner = s.inner.PaddingTop(i)
	return s
}

// PaddingBottom adds padding to the bottom of the block.
func (s *Style) PaddingBottom(i int) *Style {
	s.inner = s.inner.PaddingBottom(i)
	return s
}

// Margin is a shorthand method for setting margins on all sides at once.
//
// With one argument, the value is applied to all sides.
//
// With two arguments, the value is applied to the vertical and horizontal
// sides, in that order.
//
// With three arguments, the value is applied to the top side, the horizontal
// sides, and the bottom side, in that order.
//
// With four arguments, the value is applied clockwise starting from the top
// side, followed by the right side, then the bottom, and finally the left.
//
// With more than four arguments no margin will be added.
func (s *Style) Margin(i ...int) *Style {
	s.inner = s.inner.Margin(i...)
	return s
}

// MarginLeft sets the value of the left margin.
func (s *Style) MarginLeft(i int) *Style {
	s.inner = s.inner.MarginLeft(i)
	return s
}

// MarginRight sets the value of the right margin.
func (s *Style) MarginRight(i int) *Style {
	s.inner = s.inner.MarginRight(i)
	return s
}

// MarginTop sets the value of the top margin.
func (s *Style) MarginTop(i int) *Style {
	s.inner = s.inner.MarginTop(i)
	return s
}

// MarginBottom sets the value of the bottom margin.
func (s *Style) MarginBottom(i int) *Style {
	s.inner = s.inner.MarginBottom(i)
	return s
}

// MarginBackground sets the background color of the margin. Note that this is
// also set when inheriting from a style with a background color. In that case
// the background color on that style will set the margin color on this style.
func (s *Style) MarginBackground(color lipgloss.TerminalColor) *Style {
	s.inner = s.inner.MarginBackground(color)
	return s
}

// Border is shorthand for setting a the border style and which sides should
// have a border at once. The variadic argument sides works as follows:
//
// With one value, the value is applied to all sides.
//
// With two values, the values are applied to the vertical and horizontal
// sides, in that order.
//
// With three values, the values are applied to the top side, the horizontal
// sides, and the bottom side, in that order.
//
// With four values, the values are applied clockwise starting from the top
// side, followed by the right side, then the bottom, and finally the left.
//
// With more than four arguments the border will be applied to all sides.
//
// Examples:
//
//	 // Applies borders to the top and bottom only
//	style.New().Border(style.NormalBorder(), true, false)
//
//	 // Applies rounded borders to the right and bottom only
//	style.New().Border(style.RoundedBorder(), false, true, true, false)
func (s *Style) Border(b lipgloss.Border, sides ...bool) *Style {
	s.inner = s.inner.Border(b, sides...)
	return s
}

// BorderStyle defines the Border on a style. A Border contains a series of
// definitions for the sides and corners of a border.
//
// Note that if border visibility has not been set for any sides when setting
// the border style, the border will be enabled for all sides during rendering.
//
// You can define border characters as you'd like, though several default
// styles are included: NormalBorder(), RoundedBorder(), ThickBorder(), and
// DoubleBorder().
//
// Example:
//
//	style.New().BorderStyle(style.ThickBorder())
func (s *Style) BorderStyle(b lipgloss.Border) *Style {
	s.inner = s.inner.BorderStyle(b)
	return s
}

// BorderTop determines whether or not to draw a top border.
func (s *Style) BorderTop(b bool) *Style {
	s.inner = s.inner.BorderTop(b)
	return s
}

// BorderRight determines whether or not to draw a right border.
func (s *Style) BorderRight(b bool) *Style {
	s.inner = s.inner.BorderRight(b)
	return s
}

// BorderBottom determines whether or not to draw a bottom border.
func (s *Style) BorderBottom(b bool) *Style {
	s.inner = s.inner.BorderBottom(b)
	return s
}

// BorderLeft determines whether or not to draw a left border.
func (s *Style) BorderLeft(b bool) *Style {
	s.inner = s.inner.BorderLeft(b)
	return s
}

// BorderForeground is a shorthand function for setting all of the
// foreground colors of the borders at once. The arguments work as follows:
//
// With one argument, the argument is applied to all sides.
//
// With two arguments, the arguments are applied to the vertical and horizontal
// sides, in that order.
//
// With three arguments, the arguments are applied to the top side, the
// horizontal sides, and the bottom side, in that order.
//
// With four arguments, the arguments are applied clockwise starting from the
// top side, followed by the right side, then the bottom, and finally the left.
//
// With more than four arguments nothing will be set.
func (s *Style) BorderForeground(colors ...lipgloss.TerminalColor) *Style {
	s.inner = s.inner.BorderForeground(colors...)
	return s
}

// BorderTopForeground set the foreground color for the top of the border.
func (s *Style) BorderTopForeground(color lipgloss.TerminalColor) *Style {
	s.inner = s.inner.BorderTopForeground(color)
	return s
}

// BorderRightForeground sets the foreground color for the right side of the
// border.
func (s *Style) BorderRightForeground(color lipgloss.TerminalColor) *Style {
	s.inner = s.inner.BorderRightForeground(color)
	return s
}

// BorderBottomForeground sets the foreground color for the bottom of the
// border.
func (s *Style) BorderBottomForeground(color lipgloss.TerminalColor) *Style {
	s.inner = s.inner.BorderBottomForeground(color)
	return s
}

// BorderLeftForeground sets the foreground color for the left side of the
// border.
func (s *Style) BorderLeftForeground(color lipgloss.TerminalColor) *Style {
	s.inner = s.inner.BorderLeftForeground(color)
	return s
}

// BorderBackground is a shorthand function for setting all of the
// background colors of the borders at once. The arguments work as follows:
//
// With one argument, the argument is applied to all sides.
//
// With two arguments, the arguments are applied to the vertical and horizontal
// sides, in that order.
//
// With three arguments, the arguments are applied to the top side, the
// horizontal sides, and the bottom side, in that order.
//
// With four arguments, the arguments are applied clockwise starting from the
// top side, followed by the right side, then the bottom, and finally the left.
//
// With more than four arguments nothing will be set.
func (s *Style) BorderBackground(colors ...lipgloss.TerminalColor) *Style {
	s.inner = s.inner.BorderBackground(colors...)
	return s
}

// BorderTopBackground sets the background color of the top of the border.
func (s *Style) BorderTopBackground(color lipgloss.TerminalColor) *Style {
	s.inner = s.inner.BorderTopBackground(color)
	return s
}

// BorderRightBackground sets the background color of right side the border.
func (s *Style) BorderRightBackground(color lipgloss.TerminalColor) *Style {
	s.inner = s.inner.BorderRightBackground(color)
	return s
}

// BorderBottomBackground sets the background color of the bottom of the
// border.
func (s *Style) BorderBottomBackground(color lipgloss.TerminalColor) *Style {
	s.inner = s.inner.BorderBottomBackground(color)
	return s
}

// BorderLeftBackground set the background color of the left side of the
// border.
func (s *Style) BorderLeftBackground(color lipgloss.TerminalColor) *Style {
	s.inner = s.inner.BorderLeftBackground(color)
	return s
}

// ColorWhitespace determines whether or not the background color should be
// applied to the padding. This is true by default as it's more than likely the
// desired and expected behavior, but it can be disabled for certain graphic
// effects.
func (s *Style) ColorWhitespace(b bool) *Style {
	s.inner = s.inner.ColorWhitespace(b)
	return s
}

// MaxWidth applies a max width to a given style. This is useful in enforcing
// a certain width at render time, particularly with arbitrary strings and
// styles.
//
// Because this in intended to be used at the time of render, this method will
// not mutate the style and instead return a copy.
//
// Example:
//
//	var userInput string = "..."
//	var userStyle = text.Style{ /* ... */ }
//	fmt.Println(userStyle.MaxWidth(16).Render(userInput))
func (s *Style) MaxWidth(i int) *Style {
	s.inner = s.inner.MaxWidth(i)
	return s
}

// MaxHeight applies a max width to a given style. This is useful in enforcing
// a certain width at render time, particularly with arbitrary strings and
// styles.
//
// Because this in intended to be used at the time of render, this method will
// not mutate the style and instead return a copy.
func (s *Style) MaxHeight(i int) *Style {
	s.inner = s.inner.MaxHeight(i)
	return s
}

// Fg sets a foreground color.
//
//	// Sets the foreground to blue
//	s := New().Fg(color.NewHex("#0000ff"))
//
//	// Removes the foreground color
//	s.Foreground(color.NoColor)
func (s *Style) Fg(color lipgloss.TerminalColor) *Style {
	s.inner = s.inner.Foreground(color)
	return s
}

// Bg sets a background color.
func (s *Style) Bg(color lipgloss.TerminalColor) *Style {
	s.inner = s.inner.Background(color)
	return s
}

// UnderlineSpaces determines whether to underline spaces between words. By
// default this is true. Spaces can also be underlined without underlining the
// text itself.
func (s *Style) UnderlineSpaces(b bool) *Style {
	s.inner = s.inner.UnderlineSpaces(b)
	return s
}

// StrikethroughSpaces determines whether to apply strikethroughs to spaces
// between words. By default this is true. Spaces can also be struck without
// underlining the text itself.
func (s *Style) StrikethroughSpaces(b bool) *Style {
	s.inner = s.inner.StrikethroughSpaces(b)
	return s
}

// Render applies the defined style formatting to a given string.
func (s *Style) Render(str string) string {
	return s.inner.Render(str)
}

func (s *Style) Inner() lipgloss.Style {
	return s.inner.Copy()
}
