package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/selection/multiselect"
	"github.com/fzdwx/infinite/style"
)

func main() {
	input := components.NewInput()
	input.Prompt = "Filtering: "
	input.PromptStyle = style.New().Bold().Italic().Fg(color.LightBlue)

	keymap := components.DefaultMultiKeyMap()
	keymap.Choice = key.NewBinding(
		key.WithKeys(tea.KeySpace.String()),
	)
	_, _ = inf.NewMultiSelect([]string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
	},
		multiselect.WithKeyMap(keymap),
		multiselect.WithHintSymbol("x"),
		multiselect.WithUnHintSymbol("√"),
		multiselect.WithPageSize(10),
		//multiselect.WithDisableOutputResult(),
		//multiselect.WithCursorSymbol(emoji.PointRight),
		//multiselect.WithDisableFilter(),
		//multiselect.WithFilterInput(input),
	).
		Display("select your items!")

	//_, _ = inf.
	//	NewMultiSelect([]string{"f1", "f2", "f3"}).
	//	Display()
}
