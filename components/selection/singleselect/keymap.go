package singleselect

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/fzdwx/infinite/components"
)

type KeyMap struct {
	Up           key.Binding
	Down         key.Binding
	Choice       key.Binding
	Confirm      key.Binding
	Quit         key.Binding
	NextPage     key.Binding
	PrevPage     key.Binding
	ToggleFilter key.Binding
}

func DefaultSingleKeyMap() KeyMap {
	keyMap := components.DefaultSingleKeyMap()
	return KeyMap{
		Up:           keyMap.Up,
		Down:         keyMap.Down,
		Choice:       keyMap.Choice,
		Confirm:      keyMap.Confirm,
		Quit:         keyMap.Quit,
		ToggleFilter: keyMap.ToggleFilter,
		NextPage:     keyMap.NextPage,
		PrevPage:     keyMap.PrevPage,
	}
}

func (k KeyMap) ToSelectionKeyMap() components.SelectionKeyMap {
	return components.SelectionKeyMap{
		Up:           k.Up,
		Down:         k.Down,
		Choice:       k.Choice,
		Confirm:      k.Confirm,
		Quit:         k.Quit,
		ToggleFilter: k.ToggleFilter,
		NextPage:     k.NextPage,
		PrevPage:     k.PrevPage,
	}
}
