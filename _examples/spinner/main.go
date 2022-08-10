package main

import (
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/spinner"
	"time"
)

func main() {
	_ = inf.NewSpinner(
		spinner.WithShape(components.Dot),
		//spinner.WithDisableOutputResult(),
		spinner.WithFunc(func(spinner *spinner.Spinner) {
			for i := 0; i < 10; i++ {
				time.Sleep(time.Millisecond * 100)
				spinner.Refreshf("hello world %d", i)
			}

			spinner.Finish("finish")

			spinner.Refresh("is finish?")
		}),
	).Display()

	time.Sleep(time.Millisecond * 100 * 15)
}
