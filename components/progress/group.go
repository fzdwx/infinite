package progress

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/strx"
	"sort"
)

// Group the progress group
type Group struct {
	m       map[int]*progressUpdater
	ids     []int
	startUp *components.StartUp
	done    int
}

type progressUpdater struct {
	progress *components.Progress
	runner   func(progress *components.Progress)
}

type done int

func NewGroupWithCount(processCnt int) *Group {
	if processCnt <= 0 {
		return nil
	}

	var progressList []*components.Progress
	for i := 0; i < processCnt; i++ {
		progressList = append(progressList, components.NewProgress())
	}

	return NewGroup(progressList...)
}

// NewGroup progressList the size must > 1
func NewGroup(progressList ...*components.Progress) *Group {
	if len(progressList) <= 0 {
		return nil
	}

	m := make(map[int]*progressUpdater)
	var ids []int
	for _, progress := range progressList {
		m[progress.Id] = &progressUpdater{
			progress: progress,
		}
		ids = append(ids, progress.Id)
	}

	sort.Ints(ids)

	group := &Group{m: m, ids: ids}
	startUp := components.NewStartUp(group)
	group.startUp = startUp
	group.done = len(m)

	return group
}

func (g *Group) AppendRunner(f func(progress *components.Progress) func(progress *components.Progress)) {
	for _, updater := range g.m {
		updater.runner = f(updater.progress)
	}
}

func (g *Group) Display() error {
	for _, updater := range g.m {
		temp := updater
		go func() {
			temp.runner(temp.progress)
			g.startUp.Send(done(1))
		}()
	}
	return g.startUp.Start()
}

func (g *Group) Kill() {
	g.startUp.Kill()
}

func (g *Group) Init() tea.Cmd {
	return nil
}

func (g *Group) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case done:
		g.done -= 1
		if g.done == 0 {
			return g, tea.Quit
		}
	case components.ProgressMsg:
		if updater, ok := g.m[msg.Id]; ok {
			_, cmd := updater.progress.Update(msg)
			return g, cmd
		}
	}

	return g, nil
}

func (g *Group) View() string {
	sb := strx.NewFluent()

	for _, id := range g.ids {
		if updater, ok := g.m[id]; ok {
			sb.Write(updater.progress.View()).NewLine()
		}
	}

	return sb.String()
}

func (g *Group) SetProgram(program *tea.Program) {
	for _, updater := range g.m {
		updater.progress.SetProgram(program)
	}
}
