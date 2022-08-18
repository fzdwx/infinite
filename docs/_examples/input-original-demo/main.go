package main

import (
	"fmt"
	"github.com/fzdwx/infinite/components"
)

func main() {
	input := components.NewInput()

	if err := components.NewStartUp(input).Start(); err != nil {
		panic(err)
	}

	fmt.Println("input:", input.Value())
}
