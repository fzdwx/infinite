package components

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/strx"
	"github.com/fzdwx/infinite/style"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/reflow/ansi"
	"github.com/muesli/termenv"
	"math"
	"strings"
	"sync"
)

var (
	lastID int
	idMtx  sync.Mutex
)

// Return the next ID we should use on the model.
func nextID() int {
	idMtx.Lock()
	defer idMtx.Unlock()
	lastID++
	return lastID
}

const (
	defaultWidth = 40
)

type ProgressMsg struct {
	Id     int
	Amount int64
}
type Progress struct {
	program *tea.Program
	Id      int
	// the progress total
	Total int64
	// Current amount
	Current int64
	// Current / Total
	percent float64

	// Total width of the progress bar, including percentage, if set.
	Width int

	// "Filled" sections of the progress bar.
	Full      rune
	FullColor string

	// "Empty" sections of progress bar.
	Empty      rune
	EmptyColor string

	ShowPercentage  bool
	PercentAgeFunc  func(total, current int64, percent float64) string
	PercentAgeStyle *style.Style

	// Gradient settings
	useRamp    bool
	rampColorA colorful.Color
	rampColorB colorful.Color

	// When true, we scale the gradient to fit the width of the filled section
	// of the progress bar. When false, the width of the gradient will be set
	// to the full width of the progress bar.
	scaleRamp bool

	// Color profile for the progress bar.
	colorProfile termenv.Profile
}

func DefaultPercentAgeFunc(total, current int64, percent float64) string {
	return fmt.Sprintf(" %3.0f%%", percent*100)
}

func NewProgress() *Progress {
	p := &Progress{
		Id:              nextID(),
		Total:           100,
		Current:         0,
		PercentAgeFunc:  DefaultPercentAgeFunc,
		PercentAgeStyle: style.New().Inline(),
		Width:           defaultWidth,
		Full:            '█',
		FullColor:       "#7571F9",
		Empty:           '░',
		EmptyColor:      "#606060",
		ShowPercentage:  true,
		colorProfile:    termenv.ColorProfile(),
	}

	return p
}

// Change current val, add or sub.
func (pro *Progress) Change(amount int64) {
	pro.program.Send(ProgressMsg{
		Id:     pro.Id,
		Amount: amount,
	})
}

// Incr current val
func (pro *Progress) Incr(amount int64) {
	pro.Change(amount)
}

// Decr current val
func (pro *Progress) Decr(amount int64) {
	pro.Change(0 - amount)
}

// IncrOne incr one
func (pro *Progress) IncrOne() {
	pro.Incr(1)
}

// DecrOne decr one
func (pro *Progress) DecrOne() {
	pro.Decr(1)
}

func (pro *Progress) Init() tea.Cmd {
	return nil
}

func (pro *Progress) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ProgressMsg:
		if msg.Id == pro.Id {
			pro.refresh(msg)
		}
	}
	return pro, nil
}

func (pro *Progress) View() string {
	return pro.ViewAs(pro.percent)
}

func (pro *Progress) ViewAs(percent float64) string {
	fluent := strx.NewFluent()
	percentage := pro.viewPercentage()
	pro.barView(fluent, percent, ansi.PrintableRuneWidth(percentage))
	fluent.Write(percentage)
	return fluent.String()
}

func (pro *Progress) barView(b *strx.FluentStringBuilder, percent float64, textWidth int) {
	var (
		tw = max(0, pro.Width-textWidth)            // total width
		fw = int(math.Round(float64(tw) * percent)) // filled width
		p  float64
	)

	fw = max(0, min(tw, fw))

	if pro.useRamp {
		// Gradient fill
		for i := 0; i < fw; i++ {
			if pro.scaleRamp {
				p = float64(i) / float64(fw)
			} else {
				p = float64(i) / float64(tw)
			}
			c := pro.rampColorA.BlendLuv(pro.rampColorB, p).Hex()
			b.Write(termenv.
				String(string(pro.Full)).
				Foreground(pro.color(c)).
				String(),
			)
		}
	} else {
		// Solid fill
		s := termenv.String(string(pro.Full)).Foreground(pro.color(pro.FullColor)).String()
		b.Write(strings.Repeat(s, fw))
	}

	// Empty fill
	e := termenv.String(string(pro.Empty)).Foreground(pro.color(pro.EmptyColor)).String()
	n := max(0, tw-fw)
	b.Write(strings.Repeat(e, n))
}

func (pro *Progress) setRamp(colorA, colorB string, scaled bool) {
	// In the event of an error colors here will default to black. For
	// usability's sake, and because such an error is only cosmetic, we're
	// ignoring the error for sake of usability.
	a, _ := colorful.Hex(colorA)
	b, _ := colorful.Hex(colorB)

	pro.useRamp = true
	pro.scaleRamp = scaled
	pro.rampColorA = a
	pro.rampColorB = b
}

func (pro *Progress) color(c string) termenv.Color {
	return pro.colorProfile.Color(c)
}

func (pro *Progress) refresh(msg ProgressMsg) {
	pro.Current += msg.Amount

	if pro.Current < 0 {
		pro.Current = 0
	}

	if pro.Current > pro.Total {
		pro.Current = pro.Total
	}

	pro.percent = float64(pro.Current) / float64(pro.Total)
}

func (pro *Progress) viewPercentage() string {
	if !pro.ShowPercentage {
		return strx.Empty
	}

	return pro.PercentAgeStyle.Render(pro.PercentAgeFunc(pro.Total, pro.Current, pro.percent))
}

func (pro *Progress) SetProgram(program *tea.Program) {
	pro.program = program
}

// WithFull default '█'
func (pro *Progress) WithFull(full rune) *Progress {
	pro.Full = full
	return pro
}

// WithFullColor default "#7571F9"
func (pro *Progress) WithFullColor(full string) *Progress {
	pro.FullColor = full
	return pro
}

// WithEmpty default '░'
func (pro *Progress) WithEmpty(empty rune) *Progress {
	pro.Empty = empty
	return pro
}

// WithEmptyColor default "#606060"
func (pro *Progress) WithEmptyColor(empty string) *Progress {
	pro.EmptyColor = empty
	return pro
}

// WithTotal default 100
func (pro *Progress) WithTotal(total int64) *Progress {
	pro.Total = total
	return pro
}

// WithDisablePercentage disable output percentage.
func (pro *Progress) WithDisablePercentage() *Progress {
	pro.ShowPercentage = false
	return pro
}

// WithPercentAgeStyle replace percentage style
func (pro *Progress) WithPercentAgeStyle(sty *style.Style) *Progress {
	pro.PercentAgeStyle = sty
	return pro
}

// WithPercentAgeFunc default DefaultPercentAgeFunc
func (pro *Progress) WithPercentAgeFunc(f func(total int64, current int64, percent float64) string) *Progress {
	pro.PercentAgeFunc = f
	return pro
}

// WithDefaultGradient sets a gradient fill with default colors.
func (pro *Progress) WithDefaultGradient() *Progress {
	return pro.WithGradient("#5A56E0", "#EE6FF8")
}

// WithGradient sets a gradient fill blending between two colors.
func (pro *Progress) WithGradient(colorA, colorB string) *Progress {
	pro.setRamp(colorA, colorB, false)
	return pro
}

// WithDefaultScaledGradient sets a gradient with default colors, and scales the
// gradient to fit the filled portion of the ramp.
func (pro *Progress) WithDefaultScaledGradient() *Progress {
	return pro.WithScaledGradient("#5A56E0", "#EE6FF8")
}

// WithScaledGradient scales the gradient to fit the width of the filled portion of
// the progress bar.
func (pro *Progress) WithScaledGradient(colorA, colorB string) *Progress {
	pro.setRamp(colorA, colorB, true)
	return pro
}

// WithSolidFill sets the progress to use a solid fill with the given color.
func (pro *Progress) WithSolidFill(color string) *Progress {
	pro.FullColor = color
	pro.useRamp = false
	return pro
}

// WithWidth sets the initial width of the progress bar. Note that you can also
// set the width via the Width property, which can come in handy if you're
// waiting for a tea.WindowSizeMsg.
func (pro *Progress) WithWidth(w int) *Progress {
	pro.Width = w
	return pro
}

// WithColorProfile sets the color profile to use for the progress bar.
func (pro *Progress) WithColorProfile(p termenv.Profile) *Progress {
	pro.colorProfile = p
	return pro
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
