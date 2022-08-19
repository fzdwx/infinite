package main

import "github.com/fzdwx/infinite/components"

func main() {

	choices := []string{"a", "b", "c", "d"}

	c := components.NewSelection(choices)

	if err := components.NewStartUp(c).Start(); err != nil {
		panic(err)
	}
}
