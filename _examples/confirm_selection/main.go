package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/selection/confirm"
)

func main() {

	val, _ := inf.NewConfirmWithSelection(
		//confirm.WithDisableOutputResult(),
		confirm.WithDefaultYes(),
	).Display()

	fmt.Println(val)
}
