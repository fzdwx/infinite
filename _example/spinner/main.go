package main

import (
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/spinner"
	"time"
)

func main() {
	sp := inf.NewSpinner(
		spinner.WithShape(spinner.Dot),
		//spinner.WithDisableOutputResult(),
	).Show()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 100)
			sp.RefreshF("hello world %d", i)
		}
		sp.Finish("qqqqqqqqqqqqqqqqqqq")
	}()

	time.Sleep(time.Second * 5)
}
