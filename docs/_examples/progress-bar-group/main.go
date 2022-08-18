package main

import (
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/progress"
	"time"
)

func main() {
	cnt := 10

	if err := progress.NewGroupWithCount(10).
		AppendRunner(run(cnt)).Display(); err != nil {
		panic(err)
	}
}

func run(cnt int) func(progress *components.Progress) func() {
	return func(progress *components.Progress) func() {
		total := cnt
		cnt += 1

		progress.WithTotal(int64(total)).WithDefaultGradient()

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
	}
}

func sleep() {
	time.Sleep(time.Millisecond * 100)
}
