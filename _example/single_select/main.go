package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
	"github.com/rotisserie/eris"
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
	selected, err := inf.NewSingleSelect(options).Show("Hello world")
	if err != nil {
		fmt.Println(eris.ToString(err, true))
		return
	}

	fmt.Printf("you selectd %s\n", options[selected])

}
