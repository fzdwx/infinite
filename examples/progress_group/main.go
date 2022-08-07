package main

import (
	"fmt"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/progress"
	"time"
)

func main() {
	group := progress.NewGroup(newP(10), newP(15), newP(20))

	go func() {
		for i := 0; i < 65; i++ {
			sleep()
		}
		group.Kill()
	}()

	group.Display()
}

func newP(total int) *components.Progress {
	p := components.NewProgress().
		WithTotal(int64(total)).
		WithDefaultGradient().
		WithPercentAgeFunc(func(total int64, current int64, percent float64) string {
			return fmt.Sprintf(" %d/%d", current, total)
		})

	go func() {
		sleep()

		for i := 0; i < total+1; i++ {
			p.IncrOne()
			sleep()
		}

		for i := 0; i < total; i++ {
			p.DecrOne()
			sleep()
		}

		for i := 0; i < total+1; i++ {
			p.IncrOne()
			sleep()
		}
	}()
	return p
}

func sleep() {
	time.Sleep(time.Millisecond * 100)
}
