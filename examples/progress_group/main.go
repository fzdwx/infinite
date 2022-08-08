package main

import (
	"fmt"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/progress"
	"time"
)

func main() {
	cnt := 10

	group := progress.NewGroupWithCount(10).AppendRunner(func(progress *components.Progress) func() {
		total := cnt
		cnt += 1

		progress.WithTotal(int64(total)).
			WithDefaultGradient().
			WithPercentAgeFunc(func(total int64, current int64, percent float64) string {
				return fmt.Sprintf(" %d/%d", current, total)
			})

		return func() {

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
		}
	})

	group.Display()
}

func sleep() {
	time.Sleep(time.Millisecond * 100)
}
