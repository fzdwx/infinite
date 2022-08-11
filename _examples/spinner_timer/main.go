package main

import (
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/spinner"
	"time"
)

func main() {

	inf.NewSpinner(spinner.WithPrompt(" Loading...")).
		Display(func(spinner *spinner.Spinner) {
			start := time.Now()
			go func() {
				for {
					sleep()
					spinner.Refreshf(" stop watch %s", time.Now().Sub(start).Round(time.Millisecond).String())
				}
			}()
			time.Sleep(time.Second * 3)
		})
}

func sleep() {
	time.Sleep(time.Millisecond * 100)
}
