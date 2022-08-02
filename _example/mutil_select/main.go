package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
)

func main() {
	selected, _ := inf.
		NewMultiSelect([]string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
			inf.WithMultiSelectDefaultText("替换！！！"),
			inf.WithMultiSelectStr("x"),
			inf.WithMultiSelectUnStr("√"),
		).
		Show()

	fmt.Println(selected)
}
