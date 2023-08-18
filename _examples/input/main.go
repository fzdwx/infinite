package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/input/text"
	"github.com/fzdwx/infinite/theme"
)

func main() {

	i := inf.NewText(
		text.WithPrompt("what's your name ?"),
		text.WithPromptStyle(theme.DefaultTheme.PromptStyle),
		text.WithDefaultValue("fzdwx (maybe)"),
		text.WithRequired(),
		text.WithDefaultValueRequireValue(false),
		//text.WithFocusSymbol(theme.DefaultTheme.FocusSymbol),
		//text.WithUnFocusSymbol(theme.DefaultTheme.UnFocusSymbol),
		//text.WithFocusInterval(theme.DefaultTheme.FocusInterval),
		//text.WithUnFocusInterval(theme.DefaultTheme.UnFocusInterval),
		//text.WithFocusSymbolStyle(theme.DefaultTheme.FocusSymbolStyle),
		//text.WithUnFocusSymbolStyle(theme.DefaultTheme.UnFocusSymbolStyle),
		//text.WithFocusIntervalStyle(theme.DefaultTheme.FocusIntervalStyle),
		//text.WithUnFocusIntervalStyle(theme.DefaultTheme.UnFocusIntervalStyle),
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
