package progress

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/strx"
	"sort"
)

// Group the progress group
type Group struct {
	m       map[int]*components.Progress
	ids     []int
	startUp *components.StartUp
}

// NewGroup progressList the size must > 1
func NewGroup(progressList ...*components.Progress) *Group {
	if len(progressList) <= 0 {
		panic("progressList the size must > 1")
	}

	m := make(map[int]*components.Progress)
	var ids []int
	for _, progress := range progressList {
		m[progress.Id] = progress
		ids = append(ids, progress.Id)
	}

	sort.Ints(ids)

	group := &Group{m: m, ids: ids}
	startUp := components.NewStartUp(group)
	group.startUp = startUp

	return group
}

func (g *Group) Init() tea.Cmd {
	return nil
}

func (g *Group) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case components.ProgressMsg:
		if process, ok := g.m[msg.Id]; ok {
			_, cmd := process.Update(msg)
			return g, cmd
		}
	}

	return g, nil
}

func (g *Group) View() string {
	sb := strx.NewFluent()

	for _, id := range g.ids {
		if process, ok := g.m[id]; ok {
			sb.Write(process.View()).NewLine()
		}
	}

	return sb.String()
}

func (g *Group) SetProgram(program *tea.Program) {
	for _, progress := range g.m {
		progress.SetProgram(program)
	}
}

func (g *Group) Display() error {
	return g.startUp.Start()
}

func (g *Group) Kill() {
	g.startUp.Kill()
}
