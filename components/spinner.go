package components

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/strx"
	"github.com/fzdwx/infinite/theme"
)

type (
	SpinnerComponent struct {
		Components

		Model spinner.Model

		/* options start */
		Shape               Shape
		ShapeStyle          lipgloss.Style
		Prompt              string
		DisableOutPutResult bool
		/* options end */

		Status Status
	}

	refreshPromptMsg string
)

func NewSpinner() *SpinnerComponent {
	c := &SpinnerComponent{
		Model:               spinner.New(),
		Shape:               Line,
		ShapeStyle:          theme.DefaultTheme.SpinnerShapeStyle,
		Prompt:              "Loading...",
		DisableOutPutResult: false,
		Status:              Normal,
	}

	c.Components = Components{
		Model: c,
	}

	return c
}

func (s SpinnerComponent) RefreshPrompt(str string) {
	s.P.Send(refreshPromptMsg(str))
}

func (s *SpinnerComponent) Quit() {
	s.P.Send(QuitCmd())
}

func (s *SpinnerComponent) Quited() bool {
	return s.Status == Quit
}

func (s *SpinnerComponent) Init() tea.Cmd {
	s.Model.Spinner = spinner.Spinner{
		Frames: s.Shape.Frames,
		FPS:    s.Shape.FPS,
	}
	s.Model.Style = s.ShapeStyle

	return s.Model.Tick
}

func (s *SpinnerComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case Status:
		switch msg {
		case Quit:
			s.Status = Quit
			return s, tea.Quit
		}
	case spinner.TickMsg:
		return s.refreshSpinner(msg)
	case refreshPromptMsg:
		if !s.Quited() {
			s.Prompt = string(msg)
		}
		return s.refreshSpinner(msg)
	default:
		return s, nil
	}

	return s, nil
}

func (s *SpinnerComponent) View() string {
	viewBuilder := strx.NewFluent().
		Write(s.Model.View()).
		Write(s.Prompt)

	if s.shouldAppendNewLine() {
		viewBuilder.NewLine()
	}

	return viewBuilder.String()
}

func (s *SpinnerComponent) refreshSpinner(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	s.Model, cmd = s.Model.Update(msg)
	return s, cmd
}

func (s *SpinnerComponent) shouldAppendNewLine() bool {
	return s.Quited() && !s.DisableOutPutResult
}
