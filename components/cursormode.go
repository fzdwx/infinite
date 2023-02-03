package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/cursor"
)

// CursorMode describes the behavior of the cursor.
type CursorMode int

const (
	CursorBlink CursorMode = iota
	CursorStatic
	CursorHide
)

// String returns the cursor mode in a human-readable format. This method is
// provisional and for informational purposes only.
func (c CursorMode) String() string {
	return [...]string{
		"blink",
		"static",
		"hidden",
	}[c]
}

func (c CursorMode) Map() cursor.Mode {
	switch c {
	case CursorBlink:
		return cursor.CursorBlink
	case CursorStatic:
		return cursor.CursorStatic
	case CursorHide:
		return cursor.CursorHide
	}

	panic(fmt.Sprintf("unknow cursorMode :%d", c))
}

func newCursorMode(other cursor.Mode) CursorMode {
	switch other {
	case cursor.CursorBlink:
		return CursorBlink
	case cursor.CursorStatic:
		return CursorStatic
	case cursor.CursorHide:
		return CursorHide
	}

	panic(fmt.Sprintf("unknow cursorMode :%s", other))
}
