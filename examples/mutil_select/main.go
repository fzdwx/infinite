package main

import (
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/selection/multiselect"
	"github.com/fzdwx/infinite/emoji"
)

func main() {

	//FilterPrompt:        "Filtering: ",
	//FilterPromptStyle:   style.New().Bold().Italic().Bg(color.NewAdaptive("63", "63")).Fg(color.NewAdaptive("#ffffff", "#ffffff")),
	//	FilterPromptStyle: style.New().Bold().Italic().Fg(color.LightBlue),
	_, _ = inf.NewMultiSelect([]string{
		"Buy carrots",
		"Buy celery",
		"Buy kohlrabi",
		"Buy computer",
		"Buy something",
		"Buy car",
		"Buy subway",
	},
		multiselect.WithHintSymbol("x"),
		multiselect.WithUnHintSymbol("√"),
		//multiselect.WithDisableOutputResult(),
		multiselect.WithCursorSymbol(emoji.PointRight),
		//multiselect.WithDisableFilter(),
	).
		Display("替换！！！")

	//_, _ = inf.
	//	NewMultiSelect([]string{"f1", "f2", "f3"}).
	//	Display()
}
