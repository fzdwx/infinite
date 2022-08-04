package main

import (
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/input"
	"time"
)

func main() {

	component := input.NewComponent()
	go func() {
		component.Start()
	}()

	time.Sleep(time.Second * 3)

	component.Status = input.Quit

	inf.NewSpinner().Show()

	time.Sleep(time.Second * 3)
}
