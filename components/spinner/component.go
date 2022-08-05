package spinner

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/strx"
	"github.com/fzdwx/infinite/theme"
	"time"
)

type (
	StatusMsg struct {
		Quited bool
	}

	Component struct {
		components.Components

		Model spinner.Model

		/* options start */
		Shape               Shape
		ShapeStyle          lipgloss.Style
		Prompt              string
		TickStatusDelay     time.Duration
		DisableOutPutResult bool
		/* options end */

		Quited bool
	}
)

func NewComponent() *Component {
	c := &Component{
		Model:               spinner.New(),
		TickStatusDelay:     time.Millisecond * 50,
		Shape:               Line,
		ShapeStyle:          theme.DefaultTheme.SpinnerShapeStyle,
		Prompt:              "Loading...",
		DisableOutPutResult: false,
	}

	c.Components = components.Components{
		Model: c,
	}

	return c
}

func (c *Component) Init() tea.Cmd {
	c.Model.Spinner = spinner.Spinner{
		Frames: c.Shape.Frames,
		FPS:    c.Shape.FPS,
	}
	c.Model.Style = c.ShapeStyle

	return tea.Batch(c.Model.Tick, func() tea.Msg {
		return StatusMsg{Quited: false}
	})
}

func (c *Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case StatusMsg:
		// check should quit
		if msg.Quited {
			return c, tea.Quit
		}

		return c, c.tickStatus(c.Quited)
	case spinner.TickMsg:
		// refresh spinner
		var cmd tea.Cmd
		c.Model, cmd = c.Model.Update(msg)
		return c, cmd
	default:
		return c, nil
	}
}

func (c *Component) View() string {
	viewBuilder := strx.NewFluent().
		Write(c.Model.View()).
		Write(c.Prompt)

	if c.shouldAppendNewLine() {
		viewBuilder.Write("\n")
	}

	return viewBuilder.String()
}

func (c *Component) shouldAppendNewLine() bool {
	return c.Quited && !c.DisableOutPutResult
}

func (c *Component) tickStatus(quited bool) tea.Cmd {
	return tea.Tick(c.TickStatusDelay, func(t time.Time) tea.Msg {
		return StatusMsg{Quited: quited}
	})
}
