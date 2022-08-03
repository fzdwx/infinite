package multiselect

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/fzdwx/infinite/stringx"
	"github.com/fzdwx/infinite/theme"
)

type innerMultiSelect struct {
	choices  []string
	selected map[int]struct{}

	// current cursor index in currentChoices
	cursor int
	// the offset of screen
	scrollOffset int
	// usually len(currentChoices)
	availableChoices int
	// currently valid option
	currentChoices []string
	// how many options to display at a time
	pageSize int

	// key binding
	keymap KeyMap
	// key help text
	help help.Model

	cursorSymbol      string
	cursorSymbolStyle lipgloss.Style

	choiceTextStyle lipgloss.Style

	prompt      string
	promptStyle lipgloss.Style

	hintSymbol      string
	hintSymbolStyle lipgloss.Style

	unHintSymbol      string
	unHintSymbolStyle lipgloss.Style

	disableOutPutResult bool
	quited              bool
}

func newInnerSelect(choices []string) *innerMultiSelect {
	return &innerMultiSelect{
		choices:             choices,
		selected:            make(map[int]struct{}),
		cursorSymbol:        ">",
		cursorSymbolStyle:   theme.DefaultTheme.CursorSymbolStyle,
		choiceTextStyle:     theme.DefaultTheme.ChoiceTextStyle,
		prompt:              "Please select your options:",
		promptStyle:         theme.DefaultTheme.PromptStyle,
		hintSymbol:          "✓",
		hintSymbolStyle:     theme.DefaultTheme.MultiSelectedHintSymbolStyle,
		unHintSymbol:        "✗",
		unHintSymbolStyle:   theme.DefaultTheme.UnHintSymbolStyle,
		quited:              false,
		disableOutPutResult: false,
		pageSize:            5,
		keymap:              DefaultKeyMap,
		help:                help.New(),
	}
}

// value get all selected
func (is *innerMultiSelect) value() []int {
	var selected []int
	for s, _ := range is.selected {
		selected = append(selected, s)
	}
	return selected
}

// refreshChoices refresh choices
func (is *innerMultiSelect) refreshChoices() {
	var choices []string
	var available, ignored int

	for _, choice := range is.choices {
		available++

		if is.pageSize > 0 && len(choices) >= is.pageSize {
			break
		}

		if (is.pageSize > 0) && (ignored < is.scrollOffset) {
			ignored++

			continue
		}

		choices = append(choices, choice)
	}

	is.currentChoices = choices
	is.availableChoices = available
}

func (is *innerMultiSelect) Start() error {
	return tea.NewProgram(is).Start()
}

func (is *innerMultiSelect) Init() tea.Cmd {

	is.refreshChoices()

	return nil
}

func (is *innerMultiSelect) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, is.keymap.Confirm):
			return is.quit()
		case key.Matches(msg, is.keymap.Up):
			is.moveUp()
		case key.Matches(msg, is.keymap.Down):
			is.moveDown()
		case key.Matches(msg, is.keymap.Choice):
			is.choice()
		}
	}
	return is, nil
}

func (is *innerMultiSelect) View() string {
	if is.quited {
		return is.viewResult()
	}

	msg := stringx.NewFluentSb()

	// The header
	msg.Write(is.prompt)

	msg.NewLine()

	// Iterate over our choices
	for i, choice := range is.currentChoices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if is.cursor == i {
			cursor = is.cursorSymbol // cursor!
			choice = is.choiceTextStyle.Render(choice)
		}

		// Is this choice selected?
		checked := is.unHintSymbol // not selected
		if _, ok := is.selected[i]; ok {
			checked = is.hintSymbol // selected!
		}

		// Render the row
		msg.Write(fmt.Sprintf("%s [%s] %s", cursor, checked, choice)).
			NewLine()
	}

	// The footer
	view := is.help.View(is.keymap)
	msg.Write(view)

	// Send the UI for rendering
	return msg.String()
}

// viewResult get result
func (is innerMultiSelect) viewResult() string {
	if is.disableOutPutResult || len(is.selected) == 0 {
		return ""
	}

	output := stringx.NewFluentSb().Write(is.prompt).Space()

	for i, _ := range is.selected {
		output.Write(is.choices[i]).Space()
	}

	output.NewLine()

	return output.String()
}

// moveUp The "up" and "k" keys move the cursor up
func (is *innerMultiSelect) moveUp() {
	if is.shouldScrollUp() {
		is.scrollUp()
	}

	is.cursor = mathutil.Max(0, is.cursor-1)
}

// moveDown The "down" and "j" keys move the cursor down
func (is *innerMultiSelect) moveDown() {
	if is.shouldMoveToTop() {
		is.moveToTop()
		return
	}

	if is.shouldScrollDown() {
		is.scrollDown()
	}

	is.cursor = mathutil.Min(len(is.currentChoices)-1, is.cursor+1)
}

// choice
// The "enter" key and the spacebar (a literal space) toggle
// the selected state for the item that the cursor is pointing at.
func (is *innerMultiSelect) choice() {
	_, ok := is.selected[is.cursor]
	if ok {
		delete(is.selected, is.cursor)
	} else {
		is.selected[is.cursor] = struct{}{}
	}
}

// quit These keys should exit the program.
func (is *innerMultiSelect) quit() (tea.Model, tea.Cmd) {
	is.quited = true
	return is, tea.Quit
}

// renderColor set color to text
func (is *innerMultiSelect) renderColor() {
	is.cursorSymbol = is.cursorSymbolStyle.Render(is.cursorSymbol)
	is.prompt = is.promptStyle.Render(is.prompt)
	is.hintSymbol = is.hintSymbolStyle.Render(is.hintSymbol)
	is.unHintSymbol = is.unHintSymbolStyle.Render(is.unHintSymbol)
}

// shouldMoveToTop should move to top?
func (is *innerMultiSelect) shouldMoveToTop() bool {
	return (is.cursor + is.scrollOffset) == (len(is.choices) - 1)
}

// shouldScrollDown should scroll down?
func (is *innerMultiSelect) shouldScrollDown() bool {
	return is.cursor == len(is.currentChoices)-1 && is.canScrollDown()
}

// shouldScrollUp should scroll up?
func (is *innerMultiSelect) shouldScrollUp() bool {
	return is.cursor == 0 && is.canScrollUp()
}

// moveToTop  move cursor to top
func (is *innerMultiSelect) moveToTop() {
	is.cursor = 0
	is.scrollOffset = 0
	is.refreshChoices()
}

func (is *innerMultiSelect) scrollUp() {
	if is.pageSize <= 0 || is.scrollOffset <= 0 {
		return
	}

	is.cursor = mathutil.Min(len(is.currentChoices)-1, is.cursor+1)
	is.scrollOffset--
	is.refreshChoices()
}

func (is *innerMultiSelect) scrollDown() {
	if is.pageSize <= 0 || is.scrollOffset+is.pageSize >= is.availableChoices {
		return
	}

	is.cursor = mathutil.Max(0, is.cursor-1)
	is.scrollOffset++
	is.refreshChoices()
}

func (is *innerMultiSelect) canScrollDown() bool {
	if is.pageSize <= 0 || is.availableChoices <= is.pageSize {
		return false
	}

	if is.scrollOffset+is.pageSize >= len(is.choices) {
		return false
	}

	return true
}

func (is *innerMultiSelect) canScrollUp() bool {
	return is.scrollOffset > 0
}
