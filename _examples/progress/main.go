package main

import (
	"fmt"
	"github.com/fzdwx/infinite/components"
	"time"
)

func main() {
	var total = 10
	progress := components.NewProgress().
		WithTotal(int64(total)).
		WithDefaultGradient().
		WithPercentAgeFunc(func(total int64, current int64, percent float64) string {
			return fmt.Sprintf(" %d/%d", current, total)
		})

	startUp := components.NewStartUp(progress)

	go func() {
		for i := 0; i < 100; i++ {
			sleep()
			startUp.P.Println("hello world")
		}
	}()
	go func() {
		sleep()

		for i := 0; i < total+1; i++ {
			progress.IncrOne()
			sleep()
		}

		for i := 0; i < total; i++ {
			progress.DecrOne()
			sleep()
		}

		for i := 0; i < total+1; i++ {
			progress.IncrOne()
			sleep()
		}

		startUp.Quit()
	}()

	startUp.Start()
}

func sleep() {
	time.Sleep(time.Millisecond * 100)
}
