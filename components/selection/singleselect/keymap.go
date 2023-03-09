package singleselect

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/fzdwx/infinite/components"
)

type KeyMap struct {
	Up       key.Binding
	Down     key.Binding
	Choice   key.Binding
	Confirm  key.Binding
	Quit     key.Binding
	NextPage key.Binding
	PrevPage key.Binding
}

func DefaultSingleKeyMap() KeyMap {
	keyMap := components.DefaultSingleKeyMap()
	return KeyMap{
		Up:       keyMap.Up,
		Down:     keyMap.Down,
		Choice:   keyMap.Choice,
		Confirm:  keyMap.Confirm,
		Quit:     keyMap.Quit,
		NextPage: keyMap.NextPage,
		PrevPage: keyMap.PrevPage,
	}
}
