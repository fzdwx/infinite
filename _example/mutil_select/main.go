package main

import (
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/selectd/multiselect"
	"github.com/fzdwx/infinite/emoji"
)

func main() {
	_, _ = inf.
		NewMultiSelect([]string{
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
		).
		Show("替换！！！")

	_, _ = inf.
		NewMultiSelect([]string{"f1", "f2", "f3"}).
		Show()
}
