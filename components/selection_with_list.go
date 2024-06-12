package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
	"io"
	"strings"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type SelectionWithListKeyMap struct {
	Choice key.Binding
}

var (
	ellipsis = "…"
)

type SelectionWithList[T list.DefaultItem] struct {
	// should set before Init
	List   *list.Model
	KeyMap *SelectionWithListKeyMap

	del *selectionWithListDelegate[T]

	program *tea.Program
}

func (s *SelectionWithList[T]) Init() tea.Cmd {
	s.del.List = s.List
	s.del.KeyMap = s.KeyMap
	return nil
}

func (s *SelectionWithList[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return s, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		s.List.SetSize(msg.Width-h, msg.Height-v)
	}

	listModel, cmd := s.List.Update(msg)
	s.List = &listModel
	return s, cmd
}

func (s *SelectionWithList[T]) View() string {
	return docStyle.Render(s.List.View())
}

func (s *SelectionWithList[T]) SetProgram(program *tea.Program) {
	s.program = program
}

func (s *SelectionWithList[T]) SetItems(items []T) {
	var newItems []list.Item
	for _, item := range items {
		newItems = append(newItems, item)
	}
	s.List.SetItems(newItems)
}

func (s *SelectionWithList[T]) Value() []T {
	return s.del.Value()
}

type selectionWithListDelegate[T list.DefaultItem] struct {
	height  int
	spacing int

	Styles          list.DefaultItemStyles
	Selected        map[string]list.DefaultItem
	KeyMap          *SelectionWithListKeyMap
	List            *list.Model
	ShowDescription bool
}

func (s *selectionWithListDelegate[T]) Render(w io.Writer, m list.Model, index int, item list.Item) {
	var (
		title, desc  string
		matchedRunes []int
		st           = &s.Styles
	)

	if i, ok := item.(list.DefaultItem); ok {
		title = i.Title()
		desc = i.Description()
	} else {
		return
	}

	if m.Width() <= 0 {
		// short-circuit
		return
	}

	// Prevent text from exceeding list width
	textwidth := uint(m.Width() - st.NormalTitle.GetPaddingLeft() - st.NormalTitle.GetPaddingRight())
	title = truncate.StringWithTail(title, textwidth, ellipsis)
	if s.ShowDescription {
		var lines []string
		for i, line := range strings.Split(desc, "\n") {
			if i >= s.height-1 {
				break
			}
			lines = append(lines, truncate.StringWithTail(line, textwidth, ellipsis))
		}
		desc = strings.Join(lines, "\n")
	}

	// Conditions
	var (
		isSelected  = index == m.Index()
		emptyFilter = m.FilterState() == list.Filtering && m.FilterValue() == ""
		isFiltered  = m.FilterState() == list.Filtering || m.FilterState() == list.FilterApplied
	)

	if isFiltered && index < len(m.VisibleItems()) {
		// Get indices of matched characters
		matchedRunes = m.MatchesForItem(index)
	}

	if emptyFilter {
		title = st.DimmedTitle.Render(title)
		desc = st.DimmedDesc.Render(desc)
	} else if isSelected && m.FilterState() != list.Filtering {
		if isFiltered {
			// Highlight matches
			unmatched := st.SelectedTitle.Inline(true)
			matched := unmatched.Copy().Inherit(st.FilterMatch)
			title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
		}
		title = st.SelectedTitle.Render(title)
		desc = st.SelectedDesc.Render(desc)
	} else {
		if isFiltered {
			// Highlight matches
			unmatched := st.NormalTitle.Inline(true)
			matched := unmatched.Inherit(st.FilterMatch)
			title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
		}
		title = st.NormalTitle.Render(title)
		desc = st.NormalDesc.Render(desc)
	}

	if s.ShowDescription {
		fmt.Fprintf(w, "%st\n%st", title, desc)
		return
	}
	fmt.Fprintf(w, "%st", title)
}

func (s *selectionWithListDelegate[T]) Height() int {
	return s.height
}

func (s *selectionWithListDelegate[T]) Spacing() int {
	return s.spacing
}

func (s *selectionWithListDelegate[T]) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, s.KeyMap.Choice) {
			s.choice()
			return nil
		}
	}
	return nil
}

func (s *selectionWithListDelegate[T]) choice() {
	item := s.List.SelectedItem()
	if s.Selected[item.FilterValue()] != nil {
		delete(s.Selected, item.FilterValue())
	} else {
		s.Selected[item.FilterValue()] = item.(list.DefaultItem)
	}
}

func (s *selectionWithListDelegate[T]) Value() []T {
	var res []T
	for _, item := range s.Selected {
		res = append(res, item.(T))
	}
	return res
}
