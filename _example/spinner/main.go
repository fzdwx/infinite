package main

import (
	inf "github.com/fzdwx/infinite"
	"time"
)

func main() {
	spinner := inf.NewSpinner().Show()

	time.Sleep(time.Second * 1)

	spinner.Finish()
}
