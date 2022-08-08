package progress

import (
	"github.com/fzdwx/infinite/components"
	"io"
)

type Transfer struct {
	src      io.Reader
	dest     io.Writer
	progress *components.Progress
}

func (t Transfer) Write(bytes []byte) (n int, err error) {
	amount := len(bytes)
	t.progress.Incr(int64(amount))

	return amount, nil
}

func (t Transfer) trans() (int64, error) {
	return io.Copy(io.MultiWriter(t.dest, t), t.src)
}

// StartTransfer trans scr to desc, increase the progress by the way
func StartTransfer(src io.Reader, dest io.Writer, progress *components.Progress) (int64, error) {
	return Transfer{
		dest:     dest,
		src:      src,
		progress: progress,
	}.trans()
}
