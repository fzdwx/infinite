package main

import (
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/input/text"
	"github.com/fzdwx/infinite/theme"
	"time"
)

func main() {

	i := inf.NewText(
		text.WithPrompt("what's your name? "),
		text.WithPromptStyle(theme.DefaultTheme.PromptStyle),
		text.WithPlaceholder(" fzdwx (maybe)"),
	)
	go func() {
		i.Display()
	}()

	go func() {
		time.Sleep(time.Second * 10)
		i.Quit()
	}()

	time.Sleep(time.Second * 11)
}
