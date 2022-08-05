package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/input/confirm"
	"github.com/fzdwx/infinite/emoji"
)

func main() {

	c := inf.NewConfirm(
		confirm.WithDefaultYes(),
		//confirm.WithPrompt("hello world?"),
		confirm.WithDisplayHelp(),
		confirm.WithSymbol(emoji.Question),
		//confirm.WithKeyMap(
		//	confirm.KeyMap{Quit: key.NewBinding(
		//		key.WithKeys("c"))},
		//),
	)

	c.Display()

	if c.Value() {
		fmt.Println("yes, you are.")
	} else {
		fmt.Println("no,you are not.")
	}
}
