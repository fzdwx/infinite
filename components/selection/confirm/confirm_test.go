package confirm

import (
	tea "github.com/charmbracelet/bubbletea"
	"testing"
	"time"
)

func TestTrue(t *testing.T) {
	selection := WithSelection()

	go func() {
		time.Sleep(time.Millisecond * 100)
		selection.startUp.Send(tea.KeyMsg{Type: tea.KeyTab})
		selection.startUp.Send(tea.KeyMsg{Type: tea.KeyEnter})

	}()
	val, _ := selection.Display()
	if !val {
		t.Failed()
	}
}

func TestFalse(t *testing.T) {
	selection := WithSelection()

	go func() {
		time.Sleep(time.Millisecond * 100)
		selection.startUp.Send(tea.KeyMsg{Type: tea.KeyEnter})
	}()
	val, _ := selection.Display()

	if val {
		t.Failed()
	}
}
