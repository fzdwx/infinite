package spinner

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type (
	status struct {
		quited bool
	}

	innerSpinner struct {
		spinner spinner.Model

		tickStatusDelay time.Duration
		quited          bool
	}
)

func newInner() *innerSpinner {
	return &innerSpinner{
		spinner:         spinner.New(),
		tickStatusDelay: time.Millisecond * 50,
	}
}

func (i *innerSpinner) Start() error {
	return tea.NewProgram(i).Start()
}

func (i *innerSpinner) Init() tea.Cmd {
	return tea.Sequentially(i.spinner.Tick, func() tea.Msg {
		return status{quited: false}
	})
}

func (i *innerSpinner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case status:
		// check should quit
		if msg.quited {
			return i, tea.Quit
		}

		return i, i.tickStatus(i.quited)
	case spinner.TickMsg:
		// refresh spinner
		var cmd tea.Cmd
		i.spinner, cmd = i.spinner.Update(msg)
		return i, cmd
	default:
		return i, nil
	}
}

func (i *innerSpinner) View() string {
	return i.spinner.View()
}

func (i *innerSpinner) tickStatus(quited bool) tea.Cmd {
	return tea.Tick(i.tickStatusDelay, func(t time.Time) tea.Msg {
		return status{quited: quited}
	})
}
