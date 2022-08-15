package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/input/text"
	"github.com/fzdwx/infinite/theme"
)

func main() {

	i := inf.NewText(
		text.WithPrompt("what's your name ? "),
		text.WithPromptStyle(theme.DefaultTheme.PromptStyle),
		text.WithDefaultValue("fzdwx (maybe)"),
		text.WithRequired(),
	)
	//go func() {
	//	i.Display()
	//}()
	//
	//go func() {
	//	time.Sleep(time.Second * 10)
	//	i.Quit()
	//}()
	//
	//time.Sleep(time.Second * 11)

	_, _ = i.Display()

	fmt.Printf("you input: %s\n", i.Value())
}
