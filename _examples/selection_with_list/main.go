package main

import (
	"fmt"
	"github.com/fzdwx/infinite/components"
)

func main() {
	c := components.NewSelectionWithList([]item{"a", "b", "c"})
	_, err := components.NewStartUp(c).Run()
	if err != nil {
		panic(err)
	}

	fmt.Println(c.Value())
}

type item string

func (i item) Title() string {
	return string(i)
}

func (i item) Description() string {
	return string(i)
}

func (i item) FilterValue() string {
	return string(i)
}
