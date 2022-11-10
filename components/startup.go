package components

import (
	tea "github.com/charmbracelet/bubbletea"
)

type (
	StartUp struct {
		P       *tea.Program
		started bool
	}
)

// NewStartUp new StartUp
func NewStartUp(c Components, ops ...tea.ProgramOption) *StartUp {
	program := tea.NewProgram(c, ops...)

	c.SetProgram(program)

	return &StartUp{
		P:       program,
		started: false,
	}
}

// Start
// Deprecated: please use [StartUp.Run] instead.
func (s *StartUp) Start() error {
	_, err := s.Run()
	return err
}

// Run Returns the final model.
func (s *StartUp) Run() (tea.Model, error) {
	s.started = true
	return s.P.Run()
}

// Kill Components
func (s *StartUp) Kill() {
	if s.started {
		s.started = false
		s.P.Kill()
	}
}

func (s *StartUp) Quit() {
	if s.started {
		s.started = false
		s.P.Quit()
	}
}

// Send message to component
func (s *StartUp) Send(msg tea.Msg) {
	if s.started {
		s.P.Send(msg)
	}
}
