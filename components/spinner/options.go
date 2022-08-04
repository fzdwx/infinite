package spinner

import (
	"github.com/charmbracelet/lipgloss"
	"time"
)

type Option func(s *Spinner)

// WithShape default is Line
func WithShape(shape Shape) Option {
	return func(s *Spinner) {
		s.inner.Shape = shape
	}
}

// WithShapeStyle default theme.DefaultTheme#SpinnerShapeStyle
func WithShapeStyle(style lipgloss.Style) Option {
	return func(s *Spinner) {
		s.inner.ShapeStyle = style
	}
}

// WithTickStatusDelay default is time.Millisecond * 50
func WithTickStatusDelay(delay time.Duration) Option {
	return func(s *Spinner) {
		s.inner.TickStatusDelay = delay
	}
}
