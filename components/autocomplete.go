package components

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/fzdwx/infinite/pkg/strx"
	"strings"
)

// Suggester get suggest options
type Suggester func(valCtx AutocompleteValCtx) ([]string, bool)

// Completer result (newValue,newCursor)
// DefaultCompleter
type Completer func(valCtx AutocompleteValCtx, choiceWord string) (newVal string, newCursor int)

type AutocompleteKeyMap struct {
	Quit           key.Binding
	CloseSelection key.Binding
	Up             key.Binding
	Down           key.Binding
	Complete       key.Binding
	GotoEnd        key.Binding
}

type AutocompleteValCtx struct {
	Cursor       int
	Value        string
	autoComplete *Autocomplete
}

// CursorVal a.Value[:a.Cursor]
func (a AutocompleteValCtx) CursorVal() string {
	if len(a.Value) == 0 {
		return strx.Empty
	}

	// fix https://github.com/fzdwx/infinite/issues/9
	values := strings.Split(a.Value, strx.Empty)
	cursorVal := strings.Join(values[:a.Cursor], strx.Empty)

	return cursorVal
}

// CursorWord current word
func (a AutocompleteValCtx) CursorWord() string {
	ex := strutil.SplitEx(a.CursorVal(), strx.Space, false)
	length := len(ex)
	if length == 0 {
		return strx.Empty
	}

	return ex[length-1]
}

func (a *Autocomplete) WithInput(input *Input) *Autocomplete {
	a.Input = input
	return a
}

// WithCompleter DefaultCompleter
func (a *Autocomplete) WithCompleter(completer Completer) *Autocomplete {
	a.Completer = completer
	return a
}

// WithKeyMap DefaultAutocompleteKeyMap
func (a *Autocomplete) WithKeyMap(keyMap AutocompleteKeyMap) *Autocomplete {
	a.KeyMap = keyMap
	return a
}

// WithSelectionCreator DefaultSelectionCreator
func (a *Autocomplete) WithSelectionCreator(f func(suggester []string, a *Autocomplete) *Selection) *Autocomplete {
	a.SelectionCreator = f
	return a
}

// WithSuggestionViewRender Two implementations are provided by default: NewLineSuggestionRender or TabSuggestionRender,
// of course you can also choose to implement your own `render`
func (a *Autocomplete) WithSuggestionViewRender(f func(suggestionItems []string, a *Autocomplete) string) *Autocomplete {
	a.SuggestionViewRender = f
	return a
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
	selection.FocusSymbol = strx.Empty
	selection.UnFocusSymbol = strx.Empty
	selection.FocusInterval = strx.Empty
	selection.UnFocusInterval = strx.Empty
	selection.Init()
	selection.ShowHelp = false
	selection.ShowPaginator = false
	selection.Keymap = DefaultSingleKeyMap()
	selection.RowRender = func(CursorSymbol string, HintSymbol string, choice string) string {
		return choice
	}
	selection.SetPageSize(10)

	return selection
}

func DefaultCompleter() Completer {
	return func(valCtx AutocompleteValCtx, choiceWord string) (newVal string, newCursor int) {
		cursorVal := valCtx.CursorVal()
		cursorWord := valCtx.CursorWord()

		cursorValSplit := strutil.SplitEx(cursorVal, strx.Space, false)
		cursorValSplitLen := len(cursorValSplit)

		// replace word
		cursorValSplit[cursorValSplitLen-1] = choiceWord
		newCursorVal := strx.NewFluent().Join(cursorValSplit, strx.Space).String()

		// replace val
		newVal = strings.Replace(valCtx.Value, cursorVal, newCursorVal, 1)
		newCursor = valCtx.Cursor + (len(choiceWord) - len(cursorWord))
		return
	}
}

func NewLineSuggestionRender(suggestionItems []string, a *Autocomplete) string {
	return strx.NewFluent().WithSlice(suggestionItems, func(idx int, item string) string {
		if len(item) == 0 {
			return strx.Empty
		}

		return strx.NewFluent().Space(a.Padding + a.Input.Cursor()).Write(item).NewLine().String()
	}).String()
}

func TabSuggestionRender(suggestionItems []string, a *Autocomplete) string {
	return strx.NewFluent().WithSlice(suggestionItems, func(idx int, item string) string {
		if len(item) == 0 {
			return strx.Empty
		}

		return strx.NewFluent().Write(item).Space(4).String()
	}).String()
}

type Autocomplete struct {
	/* custom */
	Input                *Input
	Suggester            Suggester
	Completer            Completer
	KeyMap               AutocompleteKeyMap
	SelectionCreator     func(options []string, a *Autocomplete) *Selection
	SuggestionViewRender func(suggestionItems []string, a *Autocomplete) string

	Padding int
	Program *tea.Program
	*PrintHelper
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
		String()
}

func (a *Autocomplete) SetProgram(program *tea.Program) {
	a.Program = program
	a.Input.SetProgram(program)
	a.PrintHelper = NewPrintHelper(program)
}

func (a *Autocomplete) suggesterView(fluent *strx.FluentStringBuilder) {
	if a.Suggester == nil || a.ShowSelection == false {
		return
	}

	if a.ShouldNewSelection {
		suggester, ok := a.Suggester(a.getValCtx())
		if !ok || len(suggester) == 0 {
			return
		}
		a.Selection = a.SelectionCreator(suggester, a)
	}

	if a.Selection != nil {
		fluent.NewLine().Write(a.SuggestionViewRender(strings.Split(a.Selection.View(), strx.NewLine), a))
	}
}

func (a *Autocomplete) complete() {
	if a.Selection == nil {
		return
	}

	// get complete word
	a.Selection.choice()
	choiceWord := a.Selection.Choices[a.Selection.Value()[0]].Val

	newVal, newCursor := a.Completer(a.getValCtx(), choiceWord)

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

func (a *Autocomplete) getValCtx() AutocompleteValCtx {
	return AutocompleteValCtx{
		Cursor:       a.Input.Cursor(),
		Value:        a.Value(),
		autoComplete: a,
	}
}
