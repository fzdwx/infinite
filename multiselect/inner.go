package multiselect

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/stringx"
	"github.com/fzdwx/infinite/theme"
)

type innerMultiSelect struct {
	choices  []string
	cursor   int
	selected map[int]struct{}

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
	}
}

// Selected get all selected
func (is innerMultiSelect) Selected() []int {
	var selected []int
	for s, _ := range is.selected {
		selected = append(selected, s)
	}
	return selected
}

func (is *innerMultiSelect) Start() error {
	return tea.NewProgram(is).Start()
}

func (is innerMultiSelect) Init() tea.Cmd {
	return nil
}

func (is *innerMultiSelect) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return is.quit()
		case "up", "k":
			is.moveUp()
		case "down", "j":
			is.moveDown()
		case "enter", " ":
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
	for i, choice := range is.choices {

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
	msg.Write("\nPress q to quit.\n")

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
	if is.cursor > 0 {
		is.cursor--
	}
}

// moveDown The "down" and "j" keys move the cursor down
func (is *innerMultiSelect) moveDown() {
	if is.cursor < len(is.choices)-1 {
		is.cursor++
	}
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
