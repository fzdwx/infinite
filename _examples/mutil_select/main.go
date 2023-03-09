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
		"Buy carrots",
		"Buy celery",
		"Buy kohlrabi",
		"Buy computer",
		"Buy something",
		"Buy car",
		"Buy subway",
	},
		multiselect.WithKeyMap(keymap),
		multiselect.WithHintSymbol("x"),
		multiselect.WithUnHintSymbol("√"),
		//multiselect.WithDisableOutputResult(),
		//multiselect.WithCursorSymbol(emoji.PointRight),
		//multiselect.WithDisableFilter(),
		//multiselect.WithFilterInput(input),
		multiselect.WithDisableFilter(),
	).
		Display("select your items!")

	//_, _ = inf.
	//	NewMultiSelect([]string{"f1", "f2", "f3"}).
	//	Display()
}
