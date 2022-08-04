package main

import (
	"github.com/fzdwx/infinite/components/input"
	"time"
)

func main() {

	component := input.NewComponent()
	go func() {
		component.Start()
	}()
	time.Sleep(time.Second * 3)

	component.Quit()

	time.Sleep(time.Second * 3)
}
