package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/pkg/strx"
	"github.com/fzdwx/infinite/style"
	"github.com/fzdwx/infinite/theme"
	"github.com/fzdwx/iter"
	"github.com/fzdwx/iter/mapx"
	"github.com/fzdwx/iter/stream"
	"github.com/mattn/go-runewidth"
	"github.com/sahilm/fuzzy"
	"sort"
)

var (
	SelectionDefaultCursorSymbol        = ">"
	SelectionDefaultUnCursorSymbol      = " "
	SelectionDefaultCursorSymbolStyle   = theme.DefaultTheme.CursorSymbolStyle
	SelectionDefaultChoiceTextStyle     = theme.DefaultTheme.ChoiceTextStyle
	SelectionDefaultPrompt              = "Please Selection your options:"
	SelectionDefaultPromptStyle         = style.New().Bold().Fg(color.White)
	SelectionDefaultHintSymbol          = "✓"
	SelectionDefaultHintSymbolStyle     = theme.DefaultTheme.MultiSelectedHintSymbolStyle
	SelectionDefaultUnHintSymbol        = "✗"
	SelectionDefaultUnHintSymbolStyle   = theme.DefaultTheme.UnHintSymbolStyle
	SelectionDefaultDisableOutPutResult = false
	SelectionDefaultPageSize            = 5
	SelectionDefaultHelp                = help.New()
	SelectionDefaultRowRender           = DefaultRowRender
	SelectionDefaultEnableFilter        = true
	SelectionDefaultFilterInput         = NewInput()
	SelectionDefaultFilterFunc          = DefaultFilterFunc
	SelectionDefaultShowHelp            = true
)

func DefaultMultiKeyMap() SelectionKeyMap {
	return SelectionKeyMap{
		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "move down"),
		),
		Choice: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "choice it"),
		),
		Confirm: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "finish select"),
		),
		SelectAll: key.NewBinding(
			key.WithKeys("right"),
			key.WithHelp("→", "select all"),
		),
		Flip: key.NewBinding(
			key.WithKeys("left"),
			key.WithHelp("←", "flip select"),
		),
		NextPage: key.NewBinding(
			key.WithKeys(tea.KeyPgDown.String()),
			key.WithHelp("pageup", "next page"),
		),
		PrevPage: key.NewBinding(
			key.WithKeys(tea.KeyPgUp.String()),
			key.WithHelp("pagedown", "prev page"),
		),
		Quit: InterruptKey,
	}
}

func DefaultSingleKeyMap() SelectionKeyMap {
	keymap := DefaultMultiKeyMap()
	keymap.SelectAll.SetEnabled(false)
	keymap.Flip.SetEnabled(false)
	return keymap
}

type SelectionItem struct {
	Idx int
	Val string
}

type SelectionKeyMap struct {
	Up        key.Binding
	Down      key.Binding
	Choice    key.Binding
	Confirm   key.Binding
	SelectAll key.Binding // 全选
	Flip      key.Binding // 反选
	// kill program
	Quit     key.Binding
	NextPage key.Binding
	PrevPage key.Binding
}

func keyBindMatch(a key.Binding, b key.Binding) bool {
	a1Map := stream.ToMap[string, string](iter.Stream(a.Keys()), func(s string) string {
		return s
	})

	b1Map := stream.ToMap[string, string](iter.Stream(b.Keys()), func(s string) string {
		return s
	})

	return mapx.EqDefault(a1Map, b1Map)
}

func (k SelectionKeyMap) ShortHelp() []key.Binding {
	if keyBindMatch(k.Choice, k.Confirm) {
		return []key.Binding{k.Up, k.Down, k.Choice, k.Quit}
	}
	return []key.Binding{k.Up, k.Down, k.Choice, k.Confirm, k.Quit}
}

func (k SelectionKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},                // first column
		{k.Choice, k.Confirm, k.Quit}, // second column
	}
}

type Selection struct {
	// result
	Selected map[int]bool
	// Current cursor index in currentChoices
	cursor int
	// currently valid option
	currentChoices []SelectionItem
	program        *tea.Program

	Paginator paginator.Model
	Choices   []SelectionItem

	Validators       []Validator
	validatorsErrMsg []string
	// how many options to display at a time
	PageSize            int
	DisableOutPutResult bool

	// key binding
	Keymap SelectionKeyMap
	// key Help text
	Help     help.Model
	ShowHelp bool

	Prompt         string
	Header         string
	CursorSymbol   string
	UnCursorSymbol string
	HintSymbol     string
	UnHintSymbol   string

	PromptStyle       *style.Style
	CursorSymbolStyle *style.Style
	HintSymbolStyle   *style.Style
	UnHintSymbolStyle *style.Style
	ChoiceTextStyle   *style.Style

	// RowRender output options
	// CursorSymbol,HintSymbol,choice
	RowRender func(CursorSymbol string, HintSymbol string, choice string) string

	EnableFilter bool
	FilterInput  *Input
	FilterFunc   func(input string, items []SelectionItem) []SelectionItem

	FocusSymbol          string
	UnFocusSymbol        string
	FocusInterval        string
	UnFocusInterval      string
	FocusSymbolStyle     *style.Style
	UnFocusSymbolStyle   *style.Style
	FocusIntervalStyle   *style.Style
	UnFocusIntervalStyle *style.Style
	ValueStyle           *style.Style

	status Status
}

func DefaultRowRender(cursorSymbol string, hintSymbol string, choice string) string {
	return fmt.Sprintf("%s [%s] %s", cursorSymbol, hintSymbol, choice)
}

func DefaultFilterFunc(input string, items []SelectionItem) []SelectionItem {
	choiceVals := slice.Map[SelectionItem, string](items, func(index int, item SelectionItem) string {
		return item.Val
	})

	var ranks = fuzzy.Find(input, choiceVals)
	sort.Stable(ranks)

	return slice.Map[fuzzy.Match, SelectionItem](ranks, func(index int, item fuzzy.Match) SelectionItem {
		return items[item.Index]
	})
}

func (s *Selection) Init() tea.Cmd {
	var cmd tea.Cmd

	s.RefreshChoices()

	s.UnCursorSymbol = strutil.PadEnd("", runewidth.StringWidth(s.CursorSymbol), " ")

	if s.shouldFilter() {
		cmd = s.FilterInput.Init()
	}

	return cmd
}

func (s *Selection) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	shouldSkipFiler := false

	switch msg := msg.(type) {
	case tea.KeyMsg:

		// 关于为什么不用 switch, 为了适配单选的key 和 choice 和 confirm 这两个key要相同.

		if key.Matches(msg, s.Keymap.Up) {
			s.moveUp()
			shouldSkipFiler = true
		}

		if key.Matches(msg, s.Keymap.Down) {
			s.moveDown()
			shouldSkipFiler = true
		}

		if key.Matches(msg, s.Keymap.NextPage) {
			s.Paginator.NextPage()
			shouldSkipFiler = true
		}
		str := msg.String()
		_ = str == ""
		if key.Matches(msg, s.Keymap.PrevPage) {
			s.Paginator.PrevPage()
			shouldSkipFiler = true
		}

		if key.Matches(msg, s.Keymap.Choice) {
			s.choice()
			shouldSkipFiler = true
		}
		if key.Matches(msg, s.Keymap.SelectAll) {
			s.selectAll()
			shouldSkipFiler = true
		}
		if key.Matches(msg, s.Keymap.Flip) {
			s.flip()
			shouldSkipFiler = true
		}

		if key.Matches(msg, s.Keymap.Confirm) {
			for _, v := range s.Validators {
				err := v(s.Value())
				if err != nil {
					s.validatorsErrMsg = append(s.validatorsErrMsg, err.Error())
				}
			}
			if len(s.validatorsErrMsg) == 0 {
				return s.finish()
			}

			shouldSkipFiler = true
		}

		if key.Matches(msg, s.Keymap.Quit) {
			s.unselectAll()
			return s, tea.Quit
		}

		if !shouldSkipFiler && s.shouldFilter() {
			_, cmd := s.FilterInput.Update(msg)
			s.moveToTop()
			return s, cmd
		}

	case Status:
		if s.shouldFilter() {
			_, cmd := s.FilterInput.Update(msg)
			return s, cmd
		}
	}
	return s, nil
}

func (s *Selection) View() string {
	if IsFinish(s.status) {
		return s.viewResult()
	}

	msg := s.promptLine()

	if s.shouldShowValidatorsErrMsg() {
		for _, errMsg := range s.validatorsErrMsg {
			msg.NewLine().Style(
				theme.DefaultTheme.UnHintSymbolStyle,
				fmt.Sprintf("%s [%s]", SelectionDefaultUnHintSymbol, errMsg),
			)
			s.clearValidatorsErrMsg()
		}
	}

	if s.shouldFilter() {
		msg.NewLine().Write(s.FilterInput.View())
	}

	if s.Header != "" {
		msg.NewLine().Write(s.Header)
	}

	s.Paginator.PerPage = s.PageSize
	s.Paginator.SetTotalPages(len(s.currentChoices))
	start, end := s.Paginator.GetSliceBounds(len(s.currentChoices))
	// Iterate over our Choices
	for i, choice := range s.currentChoices[start:end] {
		val := choice.Val

		// Is the CursorSymbol pointing at this choice?
		cursorSymbol := s.UnCursorSymbol // no CursorSymbol
		if s.cursor == i {
			cursorSymbol = s.CursorSymbol // CursorSymbol!
			val = s.ChoiceTextStyle.Render(val)
		}

		// Is this choice Selected?
		hintSymbol := s.UnHintSymbol // not Selected
		if _, ok := s.Selected[choice.Idx]; ok {
			hintSymbol = s.HintSymbol // Selected!
		}

		// Render the row
		msg.NewLine().Write(s.RowRender(cursorSymbol, hintSymbol, val))
	}

	msg.NewLine().Write(s.Paginator.View())

	if s.ShowHelp {
		msg.NewLine().Write(s.Help.View(s.Keymap))
	}

	// Send the UI for rendering
	return msg.String()
}

func (s *Selection) SetProgram(program *tea.Program) {
	s.program = program
	if s.shouldFilter() {
		s.FilterInput.SetProgram(program)
	}
}

// Value get all Selected
func (s *Selection) Value() []int {
	var selected []int
	for s := range s.Selected {
		selected = append(selected, s)
	}
	return selected
}

// RenderColor set color to text
func (s *Selection) RenderColor() {
	s.CursorSymbol = s.CursorSymbolStyle.Render(s.CursorSymbol)
	s.Prompt = s.PromptStyle.Render(s.Prompt)
	s.HintSymbol = s.HintSymbolStyle.Render(s.HintSymbol)
	s.UnHintSymbol = s.UnHintSymbolStyle.Render(s.UnHintSymbol)
}

// RefreshChoices refresh Choices
func (s *Selection) RefreshChoices() {
	var filterChoices []SelectionItem

	// filter choice
	if s.shouldFilter() && len(s.FilterInput.Value()) > 0 {
		// do filter
		filterChoices = s.FilterFunc(s.FilterInput.Value(), s.Choices)
	} else {
		filterChoices = s.Choices
	}

	s.currentChoices = filterChoices
}

// viewResult get result
func (s *Selection) viewResult() string {
	if s.DisableOutPutResult || len(s.Selected) == 0 {
		return ""
	}

	output := s.promptLine()

	for i := range s.Selected {
		output.Style(s.ValueStyle, s.Choices[i].Val).Space()
	}

	output.NewLine()

	return output.String()
}

func (s *Selection) promptLine() *strx.FluentStringBuilder {
	builder := strx.NewFluent()

	if IsFinish(s.status) {
		builder.Style(s.UnFocusSymbolStyle, s.UnFocusSymbol).
			Write(s.Prompt).
			Style(s.UnFocusIntervalStyle, s.UnFocusInterval)
	} else {
		builder.Style(s.FocusSymbolStyle, s.FocusSymbol).
			Write(s.Prompt).
			Style(s.FocusIntervalStyle, s.FocusInterval)
	}

	return builder
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
// the Selected state for the SelectionItem that the cursor is pointing at.
func (s *Selection) choice() {
	// get Current choice.
	idx := s.currentChoices[s.cursor].Idx

	_, ok := s.Selected[idx]
	if ok {
		delete(s.Selected, idx)
	} else {
		s.Selected[idx] = true
	}
}

// selectAll add all item to Selected
func (s *Selection) selectAll() {
	for _, choice := range s.Choices {
		s.Selected[choice.Idx] = true
	}
}

// unselectAll change all item to unSelected
func (s *Selection) unselectAll() {
	for idx := range s.Choices {
		delete(s.Selected, idx)
	}
}

// flip all Selected
func (s *Selection) flip() {
	for _, choice := range s.Choices {
		_, ok := s.Selected[choice.Idx]
		if ok {
			delete(s.Selected, choice.Idx)
		} else {
			s.Selected[choice.Idx] = true
		}
	}
}

// finish These keys should exit the Program.
func (s *Selection) finish() (tea.Model, tea.Cmd) {
	s.status = Finish
	return s, tea.Quit
}

// shouldMoveToTop should move to top?
func (s *Selection) shouldMoveToTop() bool {
	return s.cursor == (s.PageSize-1) && s.Paginator.OnLastPage()
}

// shouldScrollDown should scroll down?
func (s *Selection) shouldScrollDown() bool {
	return s.cursor == (s.PageSize-1) && s.canScrollDown()
}

// shouldScrollUp should scroll up?
func (s *Selection) shouldScrollUp() bool {
	return s.cursor == 0 && s.canScrollUp()
}

// moveToTop  move cursor to top
func (s *Selection) moveToTop() {
	s.cursor = 0
	s.Paginator.Page = 0
	s.RefreshChoices()
}

func (s *Selection) scrollUp() {
	s.Paginator.PrevPage()
	s.RefreshChoices()
}

func (s *Selection) scrollDown() {
	if s.PageSize <= 0 {
		return
	}

	s.cursor = mathutil.Max(0, s.cursor-1)
	s.RefreshChoices()
}

func (s *Selection) canScrollDown() bool {
	return s.Paginator.OnLastPage() == false
}

func (s *Selection) canScrollUp() bool {
	return s.Paginator.Page != 0
}

func (s *Selection) shouldFilter() bool {
	return s.EnableFilter && s.FilterFunc != nil && s.FilterInput != nil
}

func (s *Selection) shouldShowValidatorsErrMsg() bool {
	return len(s.validatorsErrMsg) > 0
}

func (s *Selection) clearValidatorsErrMsg() {
	s.validatorsErrMsg = []string{}
}
