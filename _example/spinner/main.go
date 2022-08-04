package main

import (
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/spinner"
	"time"
)

func main() {
	sp := inf.NewSpinner(
		spinner.WithShape(spinner.Running),
	).Show()

	time.Sleep(time.Second * 3)

	sp.Finish()
}
