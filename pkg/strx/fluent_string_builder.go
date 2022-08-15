package strx

import (
	"github.com/fzdwx/infinite/style"
	"strings"
)

type (

	// FluentStringBuilder is strings.Builder wrapper,
	// but its api is fluent.
	FluentStringBuilder struct {
		sb strings.Builder
	}

	WriteFunc func(fluent *FluentStringBuilder)
)

// NewFluent new fluent string builder
func NewFluent() *FluentStringBuilder {
	return &FluentStringBuilder{
		sb: strings.Builder{},
	}
}

// NewLine append NewLine
func (b *FluentStringBuilder) NewLine() *FluentStringBuilder {
	return b.Write(NewLine)
}

// Space append Space
func (b *FluentStringBuilder) Space(times ...int) *FluentStringBuilder {
	return b.Write(RepeatSpace(times...))
}

// Write append string
func (b *FluentStringBuilder) Write(s string) *FluentStringBuilder {
	_, _ = b.sb.WriteString(s)
	return b
}

// Brackets wrap ( s )
func (b *FluentStringBuilder) Brackets(s string) *FluentStringBuilder {
	b.Write("(").Write(s).Write(")")
	return b
}

// WrapSpace " " + s + " "
func (b *FluentStringBuilder) WrapSpace(s string) *FluentStringBuilder {
	b.Write(WrapSpace(s))
	return b
}

// WriteFunc call f get string and write into FluentStringBuilder.
func (b *FluentStringBuilder) WriteFunc(f WriteFunc) *FluentStringBuilder {
	f(b)
	return b
}

// WithSlice traverse slice and call mapper
func (b *FluentStringBuilder) WithSlice(slice []string, mapper func(idx int, item string) string) *FluentStringBuilder {
	if len(slice) == 0 {
		return nil
	}

	for i, s := range slice {
		b.Write(mapper(i, s))
	}

	return b
}

func (b *FluentStringBuilder) Join(str []string, seq string) *FluentStringBuilder {
	if len(str) == 0 {
		return b
	}
	return b.Write(strings.Join(str, seq))
}

func (b *FluentStringBuilder) Style(style *style.Style, val string) *FluentStringBuilder {
	b.Write(style.Render(val))
	return b
}

func (b *FluentStringBuilder) Bool(value bool) *FluentStringBuilder {
	if value {
		b.Write("true")
	} else {
		b.Write("false")
	}
	return b
}

// Len returns the number of accumulated bytes; b.Len() == len(b.String()).
func (b *FluentStringBuilder) Len() int {
	return b.sb.Len()
}

func (b *FluentStringBuilder) String() string {
	return b.sb.String()
}
