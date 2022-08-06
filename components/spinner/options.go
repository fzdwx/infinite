package spinner

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/components"
	"time"
)

type Option func(s *Spinner)

// WithShape default is Line
func WithShape(shape components.Shape) Option {
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

// WithTickStatusDelay default is components.GlobalTickStatusDelay
func WithTickStatusDelay(delay time.Duration) Option {
	return func(s *Spinner) {
		s.inner.TickStatusDelay = delay
	}
}

// WithDisableOutputResult disable output result.
func WithDisableOutputResult() Option {
	return func(s *Spinner) {
		s.inner.DisableOutPutResult = true
	}
}
