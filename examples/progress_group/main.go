package main

import (
	"fmt"
	"github.com/duke-git/lancet/v2/random"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/progress"
	"time"
)

func main() {
	group := progress.NewGroupWithCount(10).AppendRunner(func(progress *components.Progress) func() {
		return func() {
			total := random.RandInt(10, 25)

			progress.WithTotal(int64(total)).
				WithDefaultGradient().
				WithPercentAgeFunc(func(total int64, current int64, percent float64) string {
					return fmt.Sprintf(" %d/%d", current, total)
				})

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
