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
		key.WithHelp("enter", "finish select"),
	)
	selectKeymap.Choice = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "finish select"),
	)
	selectKeymap.NextPage = key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("->", "next page"),
	)
	selectKeymap.PrevPage = key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("<-", "prev page"),
	)
	selected, err := inf.NewSingleSelect(
		options,
		//singleselect.WithDisableFilter(),
		singleselect.WithKeyBinding(selectKeymap),
		singleselect.WithPageSize(5),
	).Display("Hello world")

	if err == nil {
		fmt.Printf("you selection %s\n", options[selected])
		fmt.Printf("selected %d\n", selected)
	}

}
