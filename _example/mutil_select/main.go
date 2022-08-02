package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
)

func main() {
	selected, _ := inf.
		NewMultiSelect([]string{"Buy carrots", "Buy celery", "Buy kohlrabi"}).
		Show()

	fmt.Println(selected)
}
