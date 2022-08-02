package stringx

import "strings"

//FluentStringBuilder is strings.Builder wrapper,
// but  its api is fluent.
type FluentStringBuilder struct {
	sb strings.Builder
}

// NewFluentSb new fluent string builder
func NewFluentSb() *FluentStringBuilder {
	return &FluentStringBuilder{
		sb: strings.Builder{},
	}
}

// NextLine append "\n"
func (b *FluentStringBuilder) NextLine() *FluentStringBuilder {
	return b.Write("\n")
}

// Write append string
func (b *FluentStringBuilder) Write(s string) *FluentStringBuilder {
	_, _ = b.sb.WriteString(s)
	return b
}

func (b FluentStringBuilder) String() string {
	return b.sb.String()
}
