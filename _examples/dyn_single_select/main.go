package main

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/selection/singleselect"
	"github.com/fzdwx/iter"
	"github.com/fzdwx/iter/stream"
	"github.com/google/go-github/v52/github"
)

func main() {
	m := &model{}
	_, err := tea.NewProgram(m).Run()
	if err != nil {
		panic(err)
	}

	fmt.Println(m.Res())
}

type model struct {
	selection   *components.Selection
	filterInput *components.Input
	items       []string
	client      *github.Client
}

func (m *model) Init() tea.Cmd {
	m.client = github.NewClient(nil)
	m.filterInput = components.NewInput()
	m.selection = components.NewSelection(m.items)
	m.selection.RowRender = func(cursorSymbol string, hintSymbol string, choice string) string {
		return fmt.Sprintf("%s %s", cursorSymbol, choice)
	}
	keyMap := singleselect.DefaultSingleKeyMap()
	m.selection.Keymap = components.SelectionKeyMap{
		Up:     keyMap.Up,
		Down:   keyMap.Down,
		Choice: keyMap.Choice,
		Confirm: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "search"),
		),
		Quit:         keyMap.Quit,
		ToggleFilter: keyMap.ToggleFilter,
		NextPage:     keyMap.NextPage,
		PrevPage:     keyMap.PrevPage,
	}
	m.selection.FilterInput = m.filterInput

	return m.selection.Init()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.refreshItem()
			m.selection.Choices = mapper(m.items)
			m.selection.RefreshChoices()
			return m, nil
		case "q", "esc":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	_, cmd = m.selection.Update(msg)
	return m, cmd
}

func mapper(items []string) []components.SelectionItem {
	idx := 0
	return stream.Map[string, components.SelectionItem](iter.Stream(items), func(item string) components.SelectionItem {
		selectionItem := components.SelectionItem{
			Val: item,
			Idx: idx,
		}
		idx++
		return selectionItem
	}).ToArray()
}

func (m *model) View() string {
	return m.selection.View()
}

func (m *model) refreshItem() {
	repositories, _, err := m.client.Search.Repositories(context.Background(), m.filterInput.Value(), nil)
	if err != nil {
		return
	}
	m.items = iter.Stream(repositories.Repositories).MapTo(func(repository *github.Repository) string {
		return repository.GetFullName()
	}).ToArray()
}

func (m *model) Res() string {
	value := m.selection.Value()
	if len(value) == 0 {
		return ""
	}
	return m.items[value[0]]
}
