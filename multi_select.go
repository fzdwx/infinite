package inf

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/stringx"
	"github.com/fzdwx/infinite/style"
	"github.com/rotisserie/eris"
)

type (
	MultiSelect struct {
		inner *innerMultiSelect
	}

	MultiSelectOption func(ms *MultiSelect)
)

// Show startup MultiSelect
func (ms MultiSelect) Show(text ...string) ([]int, error) {
	ms.apply(WithDefaultText(text...))

	err := ms.inner.Start()
	if err != nil {
		return nil, eris.Wrap(err, "start inner multi select fail")
	}

	return ms.inner.Selected(), nil
}

// WithDefaultText set default text
func WithDefaultText(text ...string) MultiSelectOption {
	return func(ms *MultiSelect) {
		if len(text) >= 1 {
			ms.inner.defaultText = text[0]
		}
	}
}

// apply options on MultiSelect
func (ms *MultiSelect) apply(ops ...MultiSelectOption) *MultiSelect {
	if len(ops) > 0 {
		for _, option := range ops {
			option(ms)
		}
	}
	return ms
}

/* ============================================================== inner */
type (
	innerMultiSelect struct {
		textStyle   lipgloss.Style
		choices     []string
		cursor      int
		selected    map[int]struct{}
		defaultText string
	}
)

func newInnerSelect(choices []string) *innerMultiSelect {
	return &innerMultiSelect{
		choices:     choices,
		selected:    make(map[int]struct{}),
		defaultText: "Please select your options",
		textStyle:   style.PrimaryStyle,
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
	return startUp(is)
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
	msg := stringx.NewFluentSb()

	// The header
	msg.Write(is.textStyle.Render(is.defaultText)).NextLine()

	// Iterate over our choices
	for i, choice := range is.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if is.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := is.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		msg.Write(fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice))
	}

	// The footer
	msg.Write("\nPress q to quit.\n")

	// Send the UI for rendering
	return msg.String()
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
	return is, tea.Quit
}
