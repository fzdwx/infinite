package main

import (
	"fmt"
	"github.com/erikgeiser/promptkit/selection"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/multiselect"
	"github.com/fzdwx/infinite/style"
	"os"
)

func main() {
	t1()
	//t2()
}

func t1() {
	_, _ = inf.
		NewMultiSelect([]string{
			"1 Buy carrots",
			"2 Buy celery",
			"3 Buy kohlrabi",
			"4 Buy computer",
			"5 Buy something",
			"6 Buy car",
			"7 Buy subway",
		},
			multiselect.WithHintSymbol("x"),
			multiselect.WithUnHintSymbol("√"),
			//multiselect.WithDisableOutputResult(),
			multiselect.WithCursorSymbol("*"),
		).
		Show("替换！！！")

	fmt.Println(style.New().Foreground(color.Aqua).Render("hello world"))

	_, _ = inf.
		NewMultiSelect([]string{"f1", "f2", "f3"}).
		Show()
	//fmt.Println(selected)
}

func t2() {
	sp := selection.New("What do you pick?",
		[]*selection.Choice{
			selection.NewChoice("Horse"),
			selection.NewChoice("Car"),
			selection.NewChoice("Plane"),
			selection.NewChoice("Bike")})
	sp.PageSize = 3

	choice, err := sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the final choice
	_ = choice
}
