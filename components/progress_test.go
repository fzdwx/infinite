package components

import (
	"fmt"
	"testing"
	"time"
)

func TestProgress_ViewAs(t *testing.T) {

	t.Run("111", func(t *testing.T) {
		progress := NewProgress().WithTotal(25)

		progress.Init()

		as := progress.ViewAs(0.06, time.Now().Add(time.Second*1))
		fmt.Println(len(as))
		fmt.Println(as)

		fmt.Println(progress.ViewAs(0.12, time.Now().Add(time.Second*2)))

		fmt.Println(progress.ViewAs(0.15, time.Now().Add(time.Second*3)))

		fmt.Println(progress.ViewAs(0.45, time.Now().Add(time.Second*4)))

		viewAs := progress.ViewAs(0.75, time.Now().Add(time.Second*5))
		fmt.Println(viewAs)
		fmt.Println(len(viewAs))
	})
}
