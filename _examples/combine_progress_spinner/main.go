package main

import (
	"errors"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/pkg/strx"
	"time"
)

func main() {
	total := 10
	spinner := components.NewSpinner()
	spinner.Prompt = strx.Space + spinner.Prompt
	progress := components.NewProgress().WithTotal(int64(total))

	NewComponent(spinner, progress).Display(func(c *Component) {
		sleep()

		for i := 0; i < total+1; i++ {
			progress.IncrOne()
			sleep()
		}

		for i := 0; i < total; i++ {
			progress.DecrOne()
			sleep()
		}

		for i := 0; i < total+1; i++ {
			progress.IncrOne()
			sleep()
		}
	})
}

type Component struct {
	spinner  *components.Spinner
	progress *components.Progress
	*components.StartUp
}

func NewComponent(spinner *components.Spinner, progress *components.Progress) *Component {
	return &Component{spinner: spinner, progress: progress}
}

func (c *Component) Display(runner func(c *Component)) error {
	c.StartUp = components.NewStartUp(c)
	if runner == nil {
		return errors.New("runner is null")
	}

	go func() {
		runner(c)
		c.progress.Done()
		c.Quit()
	}()

	_, err := c.Run()
	return err
}

func (c *Component) Init() tea.Cmd {

	return tea.Batch(c.spinner.Init(), c.progress.Init())
}

func (c *Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return c, tea.Quit
		}
	}
	_, c1 := c.spinner.Update(msg)
	_, c2 := c.progress.Update(msg)

	return c, tea.Batch(c1, c2)
}

func (c *Component) View() string {
	return strx.NewFluent().Write(c.spinner.View()).Space(4).Write(c.progress.View()).String()
}

func (c *Component) SetProgram(program *tea.Program) {
	c.spinner.SetProgram(program)
	c.progress.SetProgram(program)
}

func sleep() {
	time.Sleep(time.Millisecond * 100)
}
