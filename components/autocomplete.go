package components

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/fzdwx/infinite/pkg/strx"
	"strings"
)

type Suggester func(input string) ([]string, bool)

type AutocompleteKeyMap struct {
	Quit           key.Binding
	CloseSelection key.Binding
	Up             key.Binding
	Down           key.Binding
	Complete       key.Binding
	GotoEnd        key.Binding
}

func (a *Autocomplete) WithInput(input *Input) *Autocomplete {
	a.Input = input
	return a
}

func (a *Autocomplete) WithKeyMap(keyMap AutocompleteKeyMap) *Autocomplete {
	a.KeyMap = keyMap
	return a
}

func (a *Autocomplete) WithSelectionCreator(f func(suggester []string, a *Autocomplete) *Selection) *Autocomplete {
	a.SelectionCreator = f
	return a
}

func NewAutocomplete(suggester Suggester) *Autocomplete {
	return &Autocomplete{
		Suggester:          suggester,
		Input:              NewInput(),
		KeyMap:             DefaultAutocompleteKeyMap(),
		ShowSelection:      true,
		ShouldNewSelection: true,
		SelectionCreator:   DefaultSelectionCreator,
	}
}

func DefaultAutocompleteKeyMap() AutocompleteKeyMap {
	return AutocompleteKeyMap{
		Quit:           key.NewBinding(key.WithKeys("ctrl+c")),
		CloseSelection: key.NewBinding(key.WithKeys("esc")),
		Up:             key.NewBinding(key.WithKeys("up")),
		Down:           key.NewBinding(key.WithKeys("down")),
		Complete:       key.NewBinding(key.WithKeys("tab")),
		GotoEnd:        key.NewBinding(key.WithKeys("alt+[F", "end")),
	}
}

func DefaultSelectionCreator(suggester []string, a *Autocomplete) *Selection {
	selection := NewSelection(suggester)
	selection.EnableFilter = false
	selection.Prompt = strx.Empty
	selection.Init()
	selection.ShowHelp = false
	selection.Keymap = DefaultSingleKeyMap
	selection.RowRender = func(CursorSymbol string, HintSymbol string, choice string) string {
		return strx.RepeatSpace(a.Padding+a.Input.Cursor()) + choice
	}

	return selection
}

type Autocomplete struct {
	/* custom */
	Input            *Input
	Suggester        Suggester
	KeyMap           AutocompleteKeyMap
	SelectionCreator func(options []string, a *Autocomplete) *Selection

	Padding            int
	Program            *tea.Program
	Selection          *Selection
	ShowSelection      bool
	ShouldNewSelection bool
}

// Value get user Input
func (a *Autocomplete) Value() string {
	return a.Input.Value()
}

func (a *Autocomplete) Init() tea.Cmd {
	cmd := a.Input.Init()
	a.Padding = len(a.Input.Prompt)
	return cmd
}

func (a *Autocomplete) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, a.KeyMap.Quit):
			return a, tea.Quit
		case key.Matches(msg, a.KeyMap.CloseSelection):
			a.ShowSelection = false
			return a, nil
		case key.Matches(msg, a.KeyMap.GotoEnd):
			a.Input.Model.CursorEnd()
			return a, nil
		case a.shouldMovedSelection(msg):
			a.movedSelection(msg)
			return a, nil
		case a.shouldComplete(msg):
			a.complete()
		}

		// on input text, show selection.
		switch msg.Type {
		case tea.KeyRunes:
			a.ShowSelection = true
		default:
			a.ShowSelection = false
		}
	}

	// update Input line
	_, cmd := a.Input.Update(msg)
	return a, cmd
}

func (a *Autocomplete) View() string {
	return strx.NewFluent().
		Write(a.Input.View()).
		WriteFunc(a.suggesterView).
		NewLine().
		String()
}

func (a *Autocomplete) SetProgram(program *tea.Program) {
	a.Program = program
}

func (a *Autocomplete) suggesterView(fluent *strx.FluentStringBuilder) {
	if a.Suggester == nil || a.ShowSelection == false {
		return
	}

	if a.ShouldNewSelection {
		word := a.getCursorWord()
		suggester, ok := a.Suggester(word)
		if !ok || len(suggester) == 0 {
			return
		}
		a.Selection = a.SelectionCreator(suggester, a)
	}

	if a.Selection != nil {
		fluent.Write(a.Selection.View())
	}
}

// getCursorWord Get the word at the cursor
func (a *Autocomplete) getCursorWord() string {
	ex := strutil.SplitEx(a.cursorVal(), strx.Space, false)
	length := len(ex)
	if length == 0 {
		return strx.Empty
	}

	return ex[length-1]
}

func (a *Autocomplete) complete() {
	if a.Selection == nil {
		return
	}

	// get complete word
	a.Selection.choice()
	wordChoice := a.Selection.Choices[a.Selection.Value()[0]].val

	cursorWord := a.getCursorWord()
	cursorVal := a.cursorVal()

	cursorValSplit := strutil.SplitEx(cursorVal, strx.Space, false)
	cursorValSplitLen := len(cursorValSplit)

	// replace word
	cursorValSplit[cursorValSplitLen-1] = wordChoice
	newCursorVal := strx.NewFluent().WriteStrings(cursorValSplit, strx.Space).String()

	// replace val
	newVal := strings.Replace(a.Value(), cursorVal, newCursorVal, 1)
	newCursor := a.Input.Cursor() + (len(wordChoice) - len(cursorWord))

	a.Input.Model.SetValue(newVal)
	a.Input.Model.SetCursor(newCursor)

	a.ShouldNewSelection = true
	a.Selection = nil
}

func (a *Autocomplete) movedSelection(msg tea.KeyMsg) {
	a.ShowSelection = true
	a.ShouldNewSelection = false

	switch {
	case key.Matches(msg, a.KeyMap.Up):
		a.Selection.moveUp()
	case key.Matches(msg, a.KeyMap.Down):
		a.Selection.moveDown()
	}
}

// shouldMovedSelection match moved selection key?
// e.g up/down
func (a *Autocomplete) shouldMovedSelection(msg tea.KeyMsg) bool {
	return a.Selection != nil && key.Matches(msg, a.KeyMap.Up, a.KeyMap.Down)
}

func (a *Autocomplete) shouldComplete(msg tea.KeyMsg) bool {
	return a.Selection != nil && key.Matches(msg, a.KeyMap.Complete)
}

func (a *Autocomplete) cursorVal() string {
	return a.Value()[:a.Input.Cursor()]
}
