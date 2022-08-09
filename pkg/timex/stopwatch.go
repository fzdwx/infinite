package timex

import (
	"github.com/fzdwx/infinite/pkg/strx"
	"time"
)

type StopWatchWithFunc struct {
	cost time.Duration
}

func (s StopWatchWithFunc) WriteTo(fluent *strx.FluentStringBuilder) {
	fluent.NewLine().Write(s.cost.String())
}

func StopWatch(f func()) StopWatchWithFunc {
	start := time.Now()
	f()
	return StopWatchWithFunc{cost: time.Now().Sub(start)}
}
