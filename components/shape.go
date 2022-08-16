package components

import (
	"github.com/fzdwx/infinite/emoji"
	"time"
)

type (
	// Shape the Spinner Shape
	Shape struct {
		Frames []string
		FPS    time.Duration
	}
)

// Some spinners to choose from. You could also make your own.
var (
	Line = Shape{
		Frames: []string{"|", "/", "-", "\\"},
		FPS:    time.Second / 10, //nolint:gomnd
	}
	Dot = Shape{
		Frames: []string{"⣾ ", "⣽ ", "⣻ ", "⢿ ", "⡿ ", "⣟ ", "⣯ ", "⣷ "},
		FPS:    time.Second / 10, //nolint:gomnd
	}
	MiniDot = Shape{
		Frames: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		FPS:    time.Second / 12, //nolint:gomnd
	}
	Jump = Shape{
		Frames: []string{"⢄", "⢂", "⢁", "⡁", "⡈", "⡐", "⡠"},
		FPS:    time.Second / 10, //nolint:gomnd
	}
	Pulse = Shape{
		Frames: []string{"█", "▓", "▒", "░"},
		FPS:    time.Second / 8, //nolint:gomnd
	}
	Points = Shape{
		Frames: []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●"},
		FPS:    time.Second / 7, //nolint:gomnd
	}
	Globe = Shape{
		Frames: []string{"🌍", "🌎", "🌏"},
		FPS:    time.Second / 4, //nolint:gomnd
	}
	Moon = Shape{
		Frames: []string{"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘"},
		FPS:    time.Second / 8, //nolint:gomnd
	}
	Monkey = Shape{
		Frames: []string{"🙈", "🙉", "🙊"},
		FPS:    time.Second / 3, //nolint:gomnd
	}
	Meter = Shape{
		Frames: []string{
			"▱▱▱",
			"▰▱▱",
			"▰▰▱",
			"▰▰▰",
			"▰▰▱",
			"▰▱▱",
			"▱▱▱",
		},
		FPS: time.Second / 7, //nolint:gomnd
	}
	Hamburger = Shape{
		Frames: []string{"☱", "☲", "☴", "☲"},
		FPS:    time.Second / 3, //nolint:gomnd
	}
	Running = Shape{
		Frames: []string{emoji.Walking, emoji.Running},
		FPS:    time.Second / 6, //nolint:gomnd
	}
)
