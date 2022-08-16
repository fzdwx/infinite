package components

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/pkg/strx"
	"github.com/fzdwx/infinite/style"
	"github.com/fzdwx/infinite/theme"
)

var (
	SpinnerDefaultModel               = spinner.New()
	SpinnerDefaultShape               = Line
	SpinnerDefaultShapeStyle          = theme.DefaultTheme.SpinnerShapeStyle
	SpinnerDefaultPrompt              = "Loading..."
	SpinnerDefaultDisableOutPutResult = false
	SpinnerDefaultStatus              = Normal
)

type (
	Spinner struct {
		program *tea.Program
		*PrintHelper
		Model spinner.Model

		Shape               Shape
		ShapeStyle          *style.Style
		Prompt              string
		DisableOutPutResult bool
		Status              Status
	}

	RefreshPromptMsg string
)

// RefreshPrompt refresh prompt.
func (s *Spinner) RefreshPrompt(str string) {
	s.program.Send(RefreshPromptMsg(str))
}

// Quit Spinner
func (s *Spinner) Quit() {
	s.program.Send(QuitCmd())
}

// Quited get quit state.
func (s *Spinner) Quited() bool {
	return s.Status == Quit
}

func (s *Spinner) Init() tea.Cmd {
	s.Model.Spinner = spinner.Spinner{
		Frames: s.Shape.Frames,
		FPS:    s.Shape.FPS,
	}
	s.Model.Style = s.ShapeStyle.Inner()

	return s.Model.Tick
}

func (s *Spinner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case Status:
		switch msg {
		case Quit:
			return s.quit()
		}
	case spinner.TickMsg:
		return s.refreshSpinner(msg)
	case RefreshPromptMsg:
		if !s.Quited() {
			s.Prompt = string(msg)
		}
		return s.refreshSpinner(msg)
	default:
		return s, nil
	}

	return s, nil
}

func (s *Spinner) quit() (tea.Model, tea.Cmd) {
	s.Status = Quit
	return s, tea.Quit
}

func (s *Spinner) View() string {
	viewBuilder := strx.NewFluent().
		Write(s.Model.View()).
		Write(s.Prompt)

	if s.shouldAppendNewLine() {
		viewBuilder.NewLine()
	}

	return viewBuilder.String()
}

func (s *Spinner) SetProgram(program *tea.Program) {
	s.program = program
	s.PrintHelper = NewPrintHelper(program)
}

func (s *Spinner) refreshSpinner(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	s.Model, cmd = s.Model.Update(msg)
	return s, cmd
}

func (s *Spinner) shouldAppendNewLine() bool {
	return s.Quited() && !s.DisableOutPutResult
}
