package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/fzdwx/infinite/components/selection"
	"github.com/fzdwx/infinite/strx"
	"github.com/fzdwx/infinite/theme"
	"github.com/mattn/go-runewidth"
)

type Selection struct {
	// result
	Selected map[int]struct{}
	// if true then quit.
	quited bool
	// current cursor index in currentChoices
	cursor int
	// the offset of screen
	scrollOffset int
	// usually len(currentChoices)
	availableChoices int
	// currently valid option
	currentChoices []string

	/* options start */
	Choices []string
	// how many options to display at a time
	PageSize int

	// key binding
	Keymap selection.KeyMap
	// key Help text
	Help help.Model

	Prompt         string
	CursorSymbol   string
	UnCursorSymbol string
	HintSymbol     string
	UnHintSymbol   string

	PromptStyle       lipgloss.Style
	CursorSymbolStyle lipgloss.Style
	HintSymbolStyle   lipgloss.Style
	UnHintSymbolStyle lipgloss.Style
	ChoiceTextStyle   lipgloss.Style

	DisableOutPutResult bool
	// RowRender output options
	// CursorSymbol,HintSymbol,choice
	RowRender func(string, string, string) string
	/* options end */
}

func DefaultRowRender(cursorSymbol string, hintSymbol string, choice string) string {
	return fmt.Sprintf("%s [%s] %s", cursorSymbol, hintSymbol, choice)
}

// NewSelection constructor
func NewSelection(choices []string) *Selection {
	c := &Selection{
		Choices:             choices,
		Selected:            make(map[int]struct{}),
		CursorSymbol:        ">",
		UnCursorSymbol:      " ",
		CursorSymbolStyle:   theme.DefaultTheme.CursorSymbolStyle,
		ChoiceTextStyle:     theme.DefaultTheme.ChoiceTextStyle,
		Prompt:              "Please selection your options:",
		PromptStyle:         theme.DefaultTheme.PromptStyle,
		HintSymbol:          "✓",
		HintSymbolStyle:     theme.DefaultTheme.MultiSelectedHintSymbolStyle,
		UnHintSymbol:        "✗",
		UnHintSymbolStyle:   theme.DefaultTheme.UnHintSymbolStyle,
		quited:              false,
		DisableOutPutResult: false,
		PageSize:            5,
		Keymap:              selection.DefaultMultiKeyMap,
		Help:                help.New(),
		RowRender:           DefaultRowRender,
	}
	return c
}

func (s *Selection) Init() tea.Cmd {

	s.refreshChoices()
	s.UnCursorSymbol = strutil.PadEnd("", runewidth.StringWidth(s.CursorSymbol), " ")

	return nil
}

func (s *Selection) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, s.Keymap.Choice) {
			s.choice()
		}
		if key.Matches(msg, s.Keymap.Up) {
			s.moveUp()
		}
		if key.Matches(msg, s.Keymap.Down) {
			s.moveDown()
		}
		if key.Matches(msg, s.Keymap.Confirm) {
			return s.quit()
		}
	}
	return s, nil
}

func (s *Selection) View() string {
	if s.quited {
		return s.viewResult()
	}

	msg := strx.NewFluent()

	// The header
	msg.Write(s.Prompt)

	// Iterate over our Choices
	for i, choice := range s.currentChoices {

		// Is the CursorSymbol pointing at this choice?
		cursorSymbol := s.UnCursorSymbol // no CursorSymbol
		if s.cursor == i {
			cursorSymbol = s.CursorSymbol // CursorSymbol!
			choice = s.ChoiceTextStyle.Render(choice)
		}

		// Is this choice Selected?
		hintSymbol := s.UnHintSymbol // not Selected
		if _, ok := s.Selected[i+s.scrollOffset]; ok {
			hintSymbol = s.HintSymbol // Selected!
		}

		// Render the row
		msg.NewLine().Write(s.RowRender(cursorSymbol, hintSymbol, choice))
	}

	// The footer
	msg.NewLine().Write(s.Help.View(s.Keymap))

	// Send the UI for rendering
	return msg.String()
}

func (s *Selection) SetProgram(program *tea.Program) {
}

// Value get all Selected
func (s *Selection) Value() []int {
	var selected []int
	for s, _ := range s.Selected {
		selected = append(selected, s)
	}
	return selected
}

// refreshChoices refresh Choices
func (s *Selection) refreshChoices() {
	var choices []string
	var available, ignored int

	for _, choice := range s.Choices {
		available++

		if s.PageSize > 0 && len(choices) >= s.PageSize {
			break
		}

		if (s.PageSize > 0) && (ignored < s.scrollOffset) {
			ignored++

			continue
		}

		choices = append(choices, choice)
	}

	s.currentChoices = choices
	s.availableChoices = available
}

// viewResult get result
func (s Selection) viewResult() string {
	if s.DisableOutPutResult || len(s.Selected) == 0 {
		return ""
	}

	output := strx.NewFluent().Write(s.Prompt).Space()

	for i, _ := range s.Selected {
		output.Write(s.Choices[i]).Space()
	}

	output.NewLine()

	return output.String()
}

// moveUp The "up" and "k" keys move the cursor up
func (s *Selection) moveUp() {
	if s.shouldScrollUp() {
		s.scrollUp()
	}

	s.cursor = mathutil.Max(0, s.cursor-1)
}

// moveDown The "down" and "j" keys move the cursor down
func (s *Selection) moveDown() {
	if s.shouldMoveToTop() {
		s.moveToTop()
		return
	}

	if s.shouldScrollDown() {
		s.scrollDown()
	}

	s.cursor = mathutil.Min(len(s.currentChoices)-1, s.cursor+1)
}

// choice
// The "enter" key and the spacebar (a literal space) toggle
// the Selected state for the item that the cursor is pointing at.
func (s *Selection) choice() {
	_, ok := s.Selected[s.cursor+s.scrollOffset]
	if ok {
		delete(s.Selected, s.cursor+s.scrollOffset)
	} else {
		s.Selected[s.cursor+s.scrollOffset] = struct{}{}
	}
}

// quit These keys should exit the program.
func (s *Selection) quit() (tea.Model, tea.Cmd) {
	s.quited = true
	return s, tea.Quit
}

// RenderColor set color to text
func (s *Selection) RenderColor() {
	s.CursorSymbol = s.CursorSymbolStyle.Render(s.CursorSymbol)
	s.Prompt = s.PromptStyle.Render(s.Prompt)
	s.HintSymbol = s.HintSymbolStyle.Render(s.HintSymbol)
	s.UnHintSymbol = s.UnHintSymbolStyle.Render(s.UnHintSymbol)
}

// shouldMoveToTop should move to top?
func (s *Selection) shouldMoveToTop() bool {
	return (s.cursor + s.scrollOffset) == (len(s.Choices) - 1)
}

// shouldScrollDown should scroll down?
func (s *Selection) shouldScrollDown() bool {
	return s.cursor == len(s.currentChoices)-1 && s.canScrollDown()
}

// shouldScrollUp should scroll up?
func (s *Selection) shouldScrollUp() bool {
	return s.cursor == 0 && s.canScrollUp()
}

// moveToTop  move cursor to top
func (s *Selection) moveToTop() {
	s.cursor = 0
	s.scrollOffset = 0
	s.refreshChoices()
}

func (s *Selection) scrollUp() {
	if s.PageSize <= 0 || s.scrollOffset <= 0 {
		return
	}

	s.cursor = mathutil.Min(len(s.currentChoices)-1, s.cursor+1)
	s.scrollOffset--
	s.refreshChoices()
}

func (s *Selection) scrollDown() {
	if s.PageSize <= 0 || s.scrollOffset+s.PageSize >= s.availableChoices {
		return
	}

	s.cursor = mathutil.Max(0, s.cursor-1)
	s.scrollOffset++
	s.refreshChoices()
}

func (s *Selection) canScrollDown() bool {
	if s.PageSize <= 0 || s.availableChoices <= s.PageSize {
		return false
	}

	if s.scrollOffset+s.PageSize >= len(s.Choices) {
		return false
	}

	return true
}

func (s *Selection) canScrollUp() bool {
	return s.scrollOffset > 0
}
