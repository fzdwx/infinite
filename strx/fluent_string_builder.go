package strx

import "strings"

//FluentStringBuilder is strings.Builder wrapper,
// but  its api is fluent.
type (
	FluentStringBuilder struct {
		sb strings.Builder
	}

	WriteFunc func() string
)

// NewFluent new fluent string builder
func NewFluent() *FluentStringBuilder {
	return &FluentStringBuilder{
		sb: strings.Builder{},
	}
}

// NewLine append "\n"
func (b *FluentStringBuilder) NewLine() *FluentStringBuilder {
	return b.Write("\n")
}

// Space append " "
func (b *FluentStringBuilder) Space() *FluentStringBuilder {
	return b.Write(" ")
}

// Write append string
func (b *FluentStringBuilder) Write(s string) *FluentStringBuilder {
	_, _ = b.sb.WriteString(s)
	return b
}

// WriteFunc call f get string and write into FluentStringBuilder.
func (b *FluentStringBuilder) WriteFunc(f WriteFunc) *FluentStringBuilder {
	return b.Write(f())
}

func (b FluentStringBuilder) String() string {
	return b.sb.String()
}
