package confirm

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/input"
)

// inner is confirm impl
type inner struct {
	components.Components

	input *input.Component

	// the KeyMap of Confirm
	KeyMap KeyMap

	value bool
}

func newInner() *inner {
	i := &inner{
		input:  input.NewComponent(),
		KeyMap: DefaultKeyMap,
		value:  false,
	}

	i.input.Prompt = "Are you handsome? "

	i.Components = components.Components{Model: i}
	return i
}

// Init confirm
func (i *inner) Init() tea.Cmd {
	i.input.Init()

	return input.FocusCmd
}

func (i *inner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msgCast := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msgCast, i.KeyMap.Quit):
			msg = input.Quit
		case key.Matches(msgCast, i.KeyMap.Yes):
			msg = input.Quit
			i.value = true
		case key.Matches(msgCast, i.KeyMap.No):
			msg = input.Quit
			i.value = false
		default:
			// discard, maybe output some error msg to user?
			msg = nil
		}
	}

	_, cmd := i.input.Update(msg)
	return i, cmd
}

func (i *inner) View() string {
	return i.input.View()
}
