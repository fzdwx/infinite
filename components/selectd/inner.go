package selectd

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

type InnerSelect struct {
	Choices  []string
	Selected map[int]struct{}

	// current Cursor index in CurrentChoices
	Cursor int
	// the offset of screen
	ScrollOffset int
	// usually len(CurrentChoices)
	AvailableChoices int
	// currently valid option
	CurrentChoices []string
	// how many options to display at a time
	PageSize int

	// key binding
	Keymap KeyMap
	// key Help text
	Help help.Model

	CursorSymbol      string
	CursorSymbolStyle lipgloss.Style

	ChoiceTextStyle lipgloss.Style

	Prompt      string
	PromptStyle lipgloss.Style

	HintSymbol      string
	HintSymbolStyle lipgloss.Style

	UnHintSymbol      string
	UnHintSymbolStyle lipgloss.Style

	DisableOutPutResult bool
	Quited              bool

	// RowRender output options
	// CursorSymbol,HintSymbol,choice
	RowRender func(string, string, string) string
}

func NewInnerSelect(choices []string) *InnerSelect {
	return &InnerSelect{
		Choices:             choices,
		Selected:            make(map[int]struct{}),
		CursorSymbol:        ">",
		CursorSymbolStyle:   theme.DefaultTheme.CursorSymbolStyle,
		ChoiceTextStyle:     theme.DefaultTheme.ChoiceTextStyle,
		Prompt:              "Please selectd your options:",
		PromptStyle:         theme.DefaultTheme.PromptStyle,
		HintSymbol:          "✓",
		HintSymbolStyle:     theme.DefaultTheme.MultiSelectedHintSymbolStyle,
		UnHintSymbol:        "✗",
		UnHintSymbolStyle:   theme.DefaultTheme.UnHintSymbolStyle,
		Quited:              false,
		DisableOutPutResult: false,
		PageSize:            5,
		Keymap:              DefaultMultiKeyMap,
		Help:                help.New(),
		RowRender: func(cursorSymbol string, hintSymbol string, choice string) string {
			return fmt.Sprintf("%s [%s] %s", cursorSymbol, hintSymbol, choice)
		},
	}
}

// Value get all Selected
func (is *InnerSelect) Value() []int {
	var selected []int
	for s, _ := range is.Selected {
		selected = append(selected, s)
	}
	return selected
}

// refreshChoices refresh Choices
func (is *InnerSelect) refreshChoices() {
	var choices []string
	var available, ignored int

	for _, choice := range is.Choices {
		available++

		if is.PageSize > 0 && len(choices) >= is.PageSize {
			break
		}

		if (is.PageSize > 0) && (ignored < is.ScrollOffset) {
			ignored++

			continue
		}

		choices = append(choices, choice)
	}

	is.CurrentChoices = choices
	is.AvailableChoices = available
}

func (is *InnerSelect) Start() error {
	return tea.NewProgram(is).Start()
}

func (is *InnerSelect) Init() tea.Cmd {

	is.refreshChoices()

	return nil
}

func (is *InnerSelect) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, is.Keymap.Choice) {
			is.choice()
		}
		if key.Matches(msg, is.Keymap.Up) {
			is.moveUp()
		}
		if key.Matches(msg, is.Keymap.Down) {
			is.moveDown()
		}
		if key.Matches(msg, is.Keymap.Confirm) {
			return is.quit()
		}
	}
	return is, nil
}

func (is *InnerSelect) View() string {
	if is.Quited {
		return is.viewResult()
	}

	msg := stringx.NewFluentSb()

	// The header
	msg.Write(is.Prompt)

	// Iterate over our Choices
	for i, choice := range is.CurrentChoices {

		// Is the CursorSymbol pointing at this choice?
		cursorSymbol := " " // no CursorSymbol
		if is.Cursor == i {
			cursorSymbol = is.CursorSymbol // CursorSymbol!
			choice = is.ChoiceTextStyle.Render(choice)
		}

		// Is this choice Selected?
		hintSymbol := is.UnHintSymbol // not Selected
		if _, ok := is.Selected[i+is.ScrollOffset]; ok {
			hintSymbol = is.HintSymbol // Selected!
		}

		// Render the row
		msg.NewLine().Write(is.RowRender(cursorSymbol, hintSymbol, choice))
	}

	// The footer
	msg.NewLine().Write(is.Help.View(is.Keymap))

	// Send the UI for rendering
	return msg.String()
}

// viewResult get result
func (is InnerSelect) viewResult() string {
	if is.DisableOutPutResult || len(is.Selected) == 0 {
		return ""
	}

	output := stringx.NewFluentSb().Write(is.Prompt).Space()

	for i, _ := range is.Selected {
		output.Write(is.Choices[i]).Space()
	}

	output.NewLine()

	return output.String()
}

// moveUp The "up" and "k" keys move the Cursor up
func (is *InnerSelect) moveUp() {
	if is.shouldScrollUp() {
		is.scrollUp()
	}

	is.Cursor = mathutil.Max(0, is.Cursor-1)
}

// moveDown The "down" and "j" keys move the Cursor down
func (is *InnerSelect) moveDown() {
	if is.shouldMoveToTop() {
		is.moveToTop()
		return
	}

	if is.shouldScrollDown() {
		is.scrollDown()
	}

	is.Cursor = mathutil.Min(len(is.CurrentChoices)-1, is.Cursor+1)
}

// choice
// The "enter" key and the spacebar (a literal space) toggle
// the Selected state for the item that the Cursor is pointing at.
func (is *InnerSelect) choice() {
	_, ok := is.Selected[is.Cursor+is.ScrollOffset]
	if ok {
		delete(is.Selected, is.Cursor+is.ScrollOffset)
	} else {
		is.Selected[is.Cursor+is.ScrollOffset] = struct{}{}
	}
}

// quit These keys should exit the program.
func (is *InnerSelect) quit() (tea.Model, tea.Cmd) {
	is.Quited = true
	return is, tea.Quit
}

// RenderColor set color to text
func (is *InnerSelect) RenderColor() {
	is.CursorSymbol = is.CursorSymbolStyle.Render(is.CursorSymbol)
	is.Prompt = is.PromptStyle.Render(is.Prompt)
	is.HintSymbol = is.HintSymbolStyle.Render(is.HintSymbol)
	is.UnHintSymbol = is.UnHintSymbolStyle.Render(is.UnHintSymbol)
}

// shouldMoveToTop should move to top?
func (is *InnerSelect) shouldMoveToTop() bool {
	return (is.Cursor + is.ScrollOffset) == (len(is.Choices) - 1)
}

// shouldScrollDown should scroll down?
func (is *InnerSelect) shouldScrollDown() bool {
	return is.Cursor == len(is.CurrentChoices)-1 && is.canScrollDown()
}

// shouldScrollUp should scroll up?
func (is *InnerSelect) shouldScrollUp() bool {
	return is.Cursor == 0 && is.canScrollUp()
}

// moveToTop  move Cursor to top
func (is *InnerSelect) moveToTop() {
	is.Cursor = 0
	is.ScrollOffset = 0
	is.refreshChoices()
}

func (is *InnerSelect) scrollUp() {
	if is.PageSize <= 0 || is.ScrollOffset <= 0 {
		return
	}

	is.Cursor = mathutil.Min(len(is.CurrentChoices)-1, is.Cursor+1)
	is.ScrollOffset--
	is.refreshChoices()
}

func (is *InnerSelect) scrollDown() {
	if is.PageSize <= 0 || is.ScrollOffset+is.PageSize >= is.AvailableChoices {
		return
	}

	is.Cursor = mathutil.Max(0, is.Cursor-1)
	is.ScrollOffset++
	is.refreshChoices()
}

func (is *InnerSelect) canScrollDown() bool {
	if is.PageSize <= 0 || is.AvailableChoices <= is.PageSize {
		return false
	}

	if is.ScrollOffset+is.PageSize >= len(is.Choices) {
		return false
	}

	return true
}

func (is *InnerSelect) canScrollUp() bool {
	return is.ScrollOffset > 0
}
