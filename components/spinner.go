package components

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/emoji"
	"github.com/fzdwx/infinite/pkg/strx"
	"github.com/fzdwx/infinite/style"
	"github.com/fzdwx/infinite/theme"
	"time"
)

type (
	// Shape the Spinner Shape
	Shape struct {
		Frames []string
		FPS    time.Duration
	}
)

// Some spinners to choose from. You could also make your own.
var (
	Line = Shape{
		Frames: []string{"|", "/", "-", "\\"},
		FPS:    time.Second / 10, //nolint:gomnd
	}
	Dot = Shape{
		Frames: []string{"⣾ ", "⣽ ", "⣻ ", "⢿ ", "⡿ ", "⣟ ", "⣯ ", "⣷ "},
		FPS:    time.Second / 10, //nolint:gomnd
	}
	MiniDot = Shape{
		Frames: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		FPS:    time.Second / 12, //nolint:gomnd
	}
	Jump = Shape{
		Frames: []string{"⢄", "⢂", "⢁", "⡁", "⡈", "⡐", "⡠"},
		FPS:    time.Second / 10, //nolint:gomnd
	}
	Pulse = Shape{
		Frames: []string{"█", "▓", "▒", "░"},
		FPS:    time.Second / 8, //nolint:gomnd
	}
	Points = Shape{
		Frames: []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●"},
		FPS:    time.Second / 7, //nolint:gomnd
	}
	Globe = Shape{
		Frames: []string{"🌍", "🌎", "🌏"},
		FPS:    time.Second / 4, //nolint:gomnd
	}
	Moon = Shape{
		Frames: []string{"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘"},
		FPS:    time.Second / 8, //nolint:gomnd
	}
	Monkey = Shape{
		Frames: []string{"🙈", "🙉", "🙊"},
		FPS:    time.Second / 3, //nolint:gomnd
	}
	Meter = Shape{
		Frames: []string{
			"▱▱▱",
			"▰▱▱",
			"▰▰▱",
			"▰▰▰",
			"▰▰▱",
			"▰▱▱",
			"▱▱▱",
		},
		FPS: time.Second / 7, //nolint:gomnd
	}
	Hamburger = Shape{
		Frames: []string{"☱", "☲", "☴", "☲"},
		FPS:    time.Second / 3, //nolint:gomnd
	}
	Running = Shape{
		Frames: []string{emoji.Walking, emoji.Running},
		FPS:    time.Second / 6, //nolint:gomnd
	}
)

type (
	Spinner struct {
		program *tea.Program

		Model spinner.Model

		/* options start */
		Shape               Shape
		ShapeStyle          *style.Style
		Prompt              string
		DisableOutPutResult bool
		/* options end */

		Status Status
	}

	RefreshPromptMsg string
)

// NewSpinner constructor
func NewSpinner() *Spinner {
	c := &Spinner{
		Model:               spinner.New(),
		Shape:               Line,
		ShapeStyle:          theme.DefaultTheme.SpinnerShapeStyle,
		Prompt:              "Loading...",
		DisableOutPutResult: false,
		Status:              Normal,
	}
	return c
}

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
			s.Status = Quit
			return s, tea.Quit
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
}

func (s *Spinner) refreshSpinner(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	s.Model, cmd = s.Model.Update(msg)
	return s, cmd
}

func (s *Spinner) shouldAppendNewLine() bool {
	return s.Quited() && !s.DisableOutPutResult
}
