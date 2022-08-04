package main

import (
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/input"
	"github.com/fzdwx/infinite/theme"
	"time"
)

func main() {

	i := inf.NewInput(
		input.WithPrompt("what's your name? "),
		input.WithPromptStyle(theme.DefaultTheme.PromptStyle),
		input.WithPlaceholder(" fzdwx (maybe)"),
	)
	go func() {
		i.Show()
	}()

	go func() {
		time.Sleep(time.Second * 10)
		i.Quit()
	}()

	time.Sleep(time.Second * 11)
}
