package spinner

import (
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/style"
)

type Option func(s *Spinner)

// WithPrompt replace default prompt
func WithPrompt(prompt string) Option {
	return func(s *Spinner) {
		s.inner.Prompt = prompt
	}
}

// WithShape default is Line
func WithShape(shape components.Shape) Option {
	return func(s *Spinner) {
		s.inner.Shape = shape
	}
}

// WithShapeStyle default theme.DefaultTheme#SpinnerShapeStyle
func WithShapeStyle(style *style.Style) Option {
	return func(s *Spinner) {
		s.inner.ShapeStyle = style
	}
}

// WithDisableOutputResult disable output result.
func WithDisableOutputResult() Option {
	return func(s *Spinner) {
		s.inner.DisableOutPutResult = true
	}
}
