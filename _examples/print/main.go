package main

import (
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/spinner"
	"time"
)

func main() {

	inf.NewSpinner(spinner.WithPrompt(" Loading...")).
		Display(func(spinner *spinner.Spinner) {
			go func() {
				spinner.Success("hello world")
				sleep()
				spinner.Failed("hello world")
				sleep()
				spinner.Info("hello world")
				sleep()
				spinner.Fatal("hello world")
				sleep()
				spinner.Error("hello world")
				sleep()
				spinner.Warn("hello world")
				sleep()
				spinner.Debug("hello world")
			}()
			time.Sleep(time.Second * 3)
		})
}

func sleep() {
	time.Sleep(time.Millisecond * 100)
}
