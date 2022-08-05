package selection

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/strx"
	"github.com/fzdwx/infinite/theme"
	"github.com/mattn/go-runewidth"
)

type Component struct {
	components.Components

	// result
	Selected map[int]struct{}
	// if true then quit.
	Quited bool
	// current Cursor index in CurrentChoices
	Cursor int
	// the offset of screen
	ScrollOffset int
	// usually len(CurrentChoices)
	AvailableChoices int
	// currently valid option
	CurrentChoices []string

	/* options start */
	Choices []string
	// how many options to display at a time
	PageSize int

	// key binding
	Keymap KeyMap
	// key Help text
	Help help.Model

	CursorSymbol      string
	UnCursorSymbol    string
	CursorSymbolStyle lipgloss.Style

	ChoiceTextStyle lipgloss.Style

	Prompt      string
	PromptStyle lipgloss.Style

	HintSymbol      string
	HintSymbolStyle lipgloss.Style

	UnHintSymbol      string
	UnHintSymbolStyle lipgloss.Style

	DisableOutPutResult bool
	// RowRender output options
	// CursorSymbol,HintSymbol,choice
	RowRender func(string, string, string) string
	/* options end */
}

func NewComponent(choices []string) *Component {
	c := &Component{
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
		Quited:              false,
		DisableOutPutResult: false,
		PageSize:            5,
		Keymap:              DefaultMultiKeyMap,
		Help:                help.New(),
		RowRender: func(cursorSymbol string, hintSymbol string, choice string) string {
			return fmt.Sprintf("%s [%s] %s", cursorSymbol, hintSymbol, choice)
		},
	}

	c.Components = components.Components{
		Model: c,
	}

	return c
}

func (c *Component) Init() tea.Cmd {

	c.refreshChoices()
	c.UnCursorSymbol = strutil.PadEnd("", runewidth.StringWidth(c.CursorSymbol), " ")

	return nil
}

func (c *Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, c.Keymap.Choice) {
			c.choice()
		}
		if key.Matches(msg, c.Keymap.Up) {
			c.moveUp()
		}
		if key.Matches(msg, c.Keymap.Down) {
			c.moveDown()
		}
		if key.Matches(msg, c.Keymap.Confirm) {
			return c.quit()
		}
	}
	return c, nil
}

func (c *Component) View() string {
	if c.Quited {
		return c.viewResult()
	}

	msg := strx.NewFluentSb()

	// The header
	msg.Write(c.Prompt)

	// Iterate over our Choices
	for i, choice := range c.CurrentChoices {

		// Is the CursorSymbol pointing at this choice?
		cursorSymbol := c.UnCursorSymbol // no CursorSymbol
		if c.Cursor == i {
			cursorSymbol = c.CursorSymbol // CursorSymbol!
			choice = c.ChoiceTextStyle.Render(choice)
		}

		// Is this choice Selected?
		hintSymbol := c.UnHintSymbol // not Selected
		if _, ok := c.Selected[i+c.ScrollOffset]; ok {
			hintSymbol = c.HintSymbol // Selected!
		}

		// Render the row
		msg.NewLine().Write(c.RowRender(cursorSymbol, hintSymbol, choice))
	}

	// The footer
	msg.NewLine().Write(c.Help.View(c.Keymap))

	// Send the UI for rendering
	return msg.String()
}

// Value get all Selected
func (c *Component) Value() []int {
	var selected []int
	for s, _ := range c.Selected {
		selected = append(selected, s)
	}
	return selected
}

// refreshChoices refresh Choices
func (c *Component) refreshChoices() {
	var choices []string
	var available, ignored int

	for _, choice := range c.Choices {
		available++

		if c.PageSize > 0 && len(choices) >= c.PageSize {
			break
		}

		if (c.PageSize > 0) && (ignored < c.ScrollOffset) {
			ignored++

			continue
		}

		choices = append(choices, choice)
	}

	c.CurrentChoices = choices
	c.AvailableChoices = available
}

// viewResult get result
func (c Component) viewResult() string {
	if c.DisableOutPutResult || len(c.Selected) == 0 {
		return ""
	}

	output := strx.NewFluentSb().Write(c.Prompt).Space()

	for i, _ := range c.Selected {
		output.Write(c.Choices[i]).Space()
	}

	output.NewLine()

	return output.String()
}

// moveUp The "up" and "k" keys move the Cursor up
func (c *Component) moveUp() {
	if c.shouldScrollUp() {
		c.scrollUp()
	}

	c.Cursor = mathutil.Max(0, c.Cursor-1)
}

// moveDown The "down" and "j" keys move the Cursor down
func (c *Component) moveDown() {
	if c.shouldMoveToTop() {
		c.moveToTop()
		return
	}

	if c.shouldScrollDown() {
		c.scrollDown()
	}

	c.Cursor = mathutil.Min(len(c.CurrentChoices)-1, c.Cursor+1)
}

// choice
// The "enter" key and the spacebar (a literal space) toggle
// the Selected state for the item that the Cursor is pointing at.
func (c *Component) choice() {
	_, ok := c.Selected[c.Cursor+c.ScrollOffset]
	if ok {
		delete(c.Selected, c.Cursor+c.ScrollOffset)
	} else {
		c.Selected[c.Cursor+c.ScrollOffset] = struct{}{}
	}
}

// quit These keys should exit the program.
func (c *Component) quit() (tea.Model, tea.Cmd) {
	c.Quited = true
	return c, tea.Quit
}

// RenderColor set color to text
func (c *Component) RenderColor() {
	c.CursorSymbol = c.CursorSymbolStyle.Render(c.CursorSymbol)
	c.Prompt = c.PromptStyle.Render(c.Prompt)
	c.HintSymbol = c.HintSymbolStyle.Render(c.HintSymbol)
	c.UnHintSymbol = c.UnHintSymbolStyle.Render(c.UnHintSymbol)
}

// shouldMoveToTop should move to top?
func (c *Component) shouldMoveToTop() bool {
	return (c.Cursor + c.ScrollOffset) == (len(c.Choices) - 1)
}

// shouldScrollDown should scroll down?
func (c *Component) shouldScrollDown() bool {
	return c.Cursor == len(c.CurrentChoices)-1 && c.canScrollDown()
}

// shouldScrollUp should scroll up?
func (c *Component) shouldScrollUp() bool {
	return c.Cursor == 0 && c.canScrollUp()
}

// moveToTop  move Cursor to top
func (c *Component) moveToTop() {
	c.Cursor = 0
	c.ScrollOffset = 0
	c.refreshChoices()
}

func (c *Component) scrollUp() {
	if c.PageSize <= 0 || c.ScrollOffset <= 0 {
		return
	}

	c.Cursor = mathutil.Min(len(c.CurrentChoices)-1, c.Cursor+1)
	c.ScrollOffset--
	c.refreshChoices()
}

func (c *Component) scrollDown() {
	if c.PageSize <= 0 || c.ScrollOffset+c.PageSize >= c.AvailableChoices {
		return
	}

	c.Cursor = mathutil.Max(0, c.Cursor-1)
	c.ScrollOffset++
	c.refreshChoices()
}

func (c *Component) canScrollDown() bool {
	if c.PageSize <= 0 || c.AvailableChoices <= c.PageSize {
		return false
	}

	if c.ScrollOffset+c.PageSize >= len(c.Choices) {
		return false
	}

	return true
}

func (c *Component) canScrollUp() bool {
	return c.ScrollOffset > 0
}
