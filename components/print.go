package components

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/style"
)

// default prefix
var (
	PrefixFatal = style.New().Fg(color.Red).Render("FATAL ")
	PrefixError = style.New().Fg(color.RedPink).Render("ERROR ")
	PrefixWarn  = style.New().Fg(color.Orange).Render("WARN ")
	PrefixInfo  = style.New().Fg(color.Special).Render("INFO ")
	PrefixDebug = style.New().Fg(color.Gray).Render("DEBUG ")

	PrefixSuccess = style.New().Fg(color.Special).Render("SUCCESS ")
	PrefixFailed  = style.New().Fg(color.Red).Render("FAILED ")
)

// PrintHelper Used for thread-safe output when running the Components
type PrintHelper struct {
	program *tea.Program
}

// NewPrintHelper constructor
func NewPrintHelper(program *tea.Program) *PrintHelper {
	return &PrintHelper{program: program}
}

func (p PrintHelper) Success(format string, a ...any) {
	p.PrintWithPrefix(PrefixSuccess, format, a...)
}

func (p PrintHelper) Failed(format string, a ...any) {
	p.PrintWithPrefix(PrefixFailed, format, a...)
}

func (p PrintHelper) Fatal(format string, a ...any) {
	p.PrintWithPrefix(PrefixFatal, format, a...)
}

func (p PrintHelper) Error(format string, a ...any) {
	p.PrintWithPrefix(PrefixError, format, a...)
}

func (p PrintHelper) Warn(format string, a ...any) {
	p.PrintWithPrefix(PrefixWarn, format, a...)
}

func (p PrintHelper) Info(format string, a ...any) {
	p.PrintWithPrefix(PrefixInfo, format, a...)
}

func (p PrintHelper) Debug(format string, a ...any) {
	p.PrintWithPrefix(PrefixDebug, format, a...)
}

func (p PrintHelper) Print(str string) {
	p.program.Println(str)
}

func (p PrintHelper) PrintWithPrefix(prefix, format string, a ...any) {
	printWithPrefix(p.program, prefix, format, a...)
}

func printWithPrefix(program *tea.Program, prefix, format string, a ...any) {
	program.Printf(prefix + fmt.Sprintf(format, a...))
}
