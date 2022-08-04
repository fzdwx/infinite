package spinner

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/theme"
	"time"
)

type (
	StatusMsg struct {
		Quited bool
	}

	InnerSpinner struct {
		components.Component

		Model spinner.Model

		/* options start */
		Shape           Shape
		ShapeStyle      lipgloss.Style
		TickStatusDelay time.Duration
		/* options end */

		Quited bool
	}
)

func NewInner() *InnerSpinner {
	i := &InnerSpinner{
		Model:           spinner.New(),
		TickStatusDelay: time.Millisecond * 50,
		Shape:           Line,
		ShapeStyle:      theme.DefaultTheme.SpinnerShapeStyle,
	}

	i.Component = components.Component{
		Model: i,
	}

	return i
}

func (i *InnerSpinner) Init() tea.Cmd {
	i.Model.Spinner = spinner.Spinner{
		Frames: i.Shape.Frames,
		FPS:    i.Shape.FPS,
	}
	i.Model.Style = i.ShapeStyle

	return tea.Batch(i.Model.Tick, func() tea.Msg {
		return StatusMsg{Quited: false}
	})
}

func (i *InnerSpinner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case StatusMsg:
		// check should quit
		if msg.Quited {
			return i, tea.Quit
		}

		return i, i.TickStatus(i.Quited)
	case spinner.TickMsg:
		// refresh spinner
		var cmd tea.Cmd
		i.Model, cmd = i.Model.Update(msg)
		return i, cmd
	default:
		return i, nil
	}
}

func (i *InnerSpinner) View() string {
	return i.Model.View()
}

func (i *InnerSpinner) TickStatus(quited bool) tea.Cmd {
	return tea.Tick(i.TickStatusDelay, func(t time.Time) tea.Msg {
		return StatusMsg{Quited: quited}
	})
}
