package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
)

func main() {

	val, _ := inf.NewConfirmWithSelection(
	//confirm.WithDisOutResult(),
	).Display()

	fmt.Println(val)
}
