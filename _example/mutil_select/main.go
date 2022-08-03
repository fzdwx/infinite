package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/multiselect"
	"github.com/fzdwx/infinite/style"
)

func main() {
	_, _ = inf.
		NewMultiSelect([]string{
			"1 Buy carrots",
			"2 Buy celery",
			"3 Buy kohlrabi",
			"4 Buy computer",
			"5 Buy something",
			"6 Buy car",
			"7 Buy subway",
		},
			multiselect.WithHintSymbol("x"),
			multiselect.WithUnHintSymbol("√"),
			//multiselect.WithDisableOutputResult(),
			multiselect.WithCursorSymbol("*"),
		).
		Show("替换！！！")

	fmt.Println(style.New().Foreground(color.Aqua).Render("hello world"))

	_, _ = inf.
		NewMultiSelect([]string{"f1", "f2", "f3"}).
		Show()
}
