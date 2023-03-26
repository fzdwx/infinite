package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/selection/singleselect"
)

func main() {
	options := []string{
		"1 Buy carrots",
		"2 Buy celery",
		"3 Buy kohlrabi",
		"4 Buy computer",
		"5 Buy something",
		"6 Buy car",
		"7 Buy subway",
	}

	selectKeymap := singleselect.DefaultSingleKeyMap()
	selectKeymap.Confirm = key.NewBinding(
		key.WithKeys("enter"),
	)
	selectKeymap.Choice = key.NewBinding(
		key.WithKeys("enter"),
	)
	selected, err := inf.NewSingleSelect(
		options,
		singleselect.WithDisableFilter(),
		singleselect.WithKeyBinding(selectKeymap),
	).Display("Hello world")

	if err == nil {
		fmt.Printf("you selection %s\n", options[selected])
	}

}
