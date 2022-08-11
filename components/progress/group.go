package progress

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/pkg/strx"
	"sort"
)

// WithDoneView when Group is done, will callback this func,
// will overwrite the progress doneView in the Group.
func (g *Group) WithDoneView(f func() string) *Group {
	g.doneView = f
	return g
}

// Group the progress group
type Group struct {
	*components.PrintHelper
	m        map[int]*progressUpdater
	ids      []int
	startUp  *components.StartUp
	done     int
	doneView func() string
}

type progressUpdater struct {
	progress *components.Progress
	runner   func()
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

	return group
}

func (g *Group) AppendRunner(f func(progress *components.Progress) func()) *Group {
	g.foreach(func(updater *progressUpdater) {
		runner := f(updater.progress)
		if runner != nil {
			updater.runner = runner
			g.done++
		}
	})
	return g
}

func (g *Group) Display() error {
	g.foreach(func(updater *progressUpdater) {
		go func() {
			updater.runner()
			updater.progress.Done()
			g.startUp.Send(done(1))
		}()
	})
	return g.startUp.Start()
}

func (g *Group) Kill() {
	g.startUp.Kill()
}

func (g *Group) Init() tea.Cmd {
	g.foreach(func(updater *progressUpdater) {
		updater.progress.Init()
	})
	return nil
}

func (g *Group) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case done:
		g.done--
		if g.isDone() {
			return g, tea.Quit
		}
	case components.ProgressMsg:
		if updater, ok := g.m[msg.Id]; ok {
			_, cmd := updater.progress.Update(msg)
			return g, cmd
		}
	case tea.WindowSizeMsg:
		g.foreach(func(updater *progressUpdater) {
			updater.progress.Update(msg)
		})
	}

	return g, nil
}

func (g *Group) View() string {
	if g.shouldOutputDoneView() {
		return g.doneView() + strx.NewLine
	}

	sb := strx.NewFluent()

	g.foreach(func(updater *progressUpdater) {
		sb.Write(updater.progress.View()).NewLine()
	})

	return sb.String()
}

func (g *Group) SetProgram(program *tea.Program) {
	for _, updater := range g.m {
		updater.progress.SetProgram(program)
	}
	g.PrintHelper = components.NewPrintHelper(program)
}

func (g *Group) foreach(f func(updater *progressUpdater)) {
	for _, id := range g.ids {
		if updater, ok := g.m[id]; ok {
			f(updater)
		}
	}
}

func (g *Group) shouldOutputDoneView() bool {
	return g.isDone() && g.doneView != nil
}

func (g *Group) isDone() bool {
	return g.done == 0
}
