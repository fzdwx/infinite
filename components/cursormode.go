package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
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

func (c CursorMode) Map() textinput.CursorMode {
	switch c {
	case CursorBlink:
		return textinput.CursorBlink
	case CursorStatic:
		return textinput.CursorStatic
	case CursorHide:
		return textinput.CursorHide
	}

	panic(fmt.Sprintf("unknow cursorMode :%d", c))
}

func newCursorMode(other textinput.CursorMode) CursorMode {
	switch other {
	case textinput.CursorBlink:
		return CursorBlink
	case textinput.CursorStatic:
		return CursorStatic
	case textinput.CursorHide:
		return CursorHide
	}

	panic(fmt.Sprintf("unknow cursorMode :%s", other))
}
