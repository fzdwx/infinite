<div align="center">
<h1>infinite</h1>
<span>🧬 A component library for developing interactive CLI(tui,terminal) programs.</span>
<br>
<a href="https://goreportcard.com/report/github.com/fzdwx/infinite"><img src="https://goreportcard.com/badge/github.com/fzdwx/infinite" alt="go report card"></a>
<a href="https://github.com/fzdwx/infinite/releases"><img src="https://img.shields.io/github/v/release/fzdwx/infinite.svg?style=flat-square" alt="release"></a>
</div>
<img src="https://user-images.githubusercontent.com/65269574/183641765-e8de7441-3c4e-4008-b2a9-b2ba556ddd72.gif" alt="demo">

[中文](https://github.com/fzdwx/infinite/blob/main/docs/en/README.md) | English

## Features

- Provides a set of out-of-the-box components
    - autocomplete
    - progress-bar group
    - multi/single select
    - spinner
    - confirm
    - input
- Support window/linux (I only have these two operating systems)
- Customizable (You can replace the implementation of some methods in the component)
- Combinable (You can combine multiple components to use)
    - `autocomplete` is composed of `input` and `select`
    - `select` implements the filter function by embedding the `input`
    - ...

## Best Practices

1. Update the status through messages, that is, send messages through `program.Send(msg)`, `Update` listen and update
   the status, and finally feedback the results through `View`.
2. ...

## Install

```bash
go get github.com/fzdwx/infinite
```

## Showcase

### Combined demo

A demo combining `progress` and `spinner`

![demo](https://user-images.githubusercontent.com/65269574/184496950-dbc246e7-5199-4e85-8167-1292b6eeb574.gif)

[code](https://github.com/fzdwx/infinite/blob/main/_examples/combine_progress_spinner/main.go)

---

### Progress group

![demo](https://user-images.githubusercontent.com/65269574/183296585-b0a56827-d9d9-4258-ad32-266ada01b1ed.gif)

<details>
<summary>code</summary>

```go
package main

import (
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/progress"
	"time"
)

func main() {
	cnt := 10

	group := progress.NewGroupWithCount(10).
		AppendRunner(func(progress *components.Progress) func() {
			total := cnt
			cnt += 1
			progress.WithTotal(int64(total)).
				WithDefaultGradient()

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
```

</details>

---

### Multiple select

![demo](https://user-images.githubusercontent.com/65269574/183274216-d2a7af91-0581-4d13-b8c2-00b9aad5ef3a.gif)

<details>
<summary>code</summary>

```go
package main

import (
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/selection/multiselect"
	"github.com/fzdwx/infinite/style"
)

func main() {
	input := components.NewInput()
	input.Prompt = "Filtering: "
	input.PromptStyle = style.New().Bold().Italic().Fg(color.LightBlue)

	_, _ = inf.NewMultiSelect([]string{
		"Buy carrots",
		"Buy celery",
		"Buy kohlrabi",
		"Buy computer",
		"Buy something",
		"Buy car",
		"Buy subway",
	},
		multiselect.WithFilterInput(input),
	).Display("select your items!")
}
```

</details>

---

### Spinner

![demo](https://user-images.githubusercontent.com/65269574/183074665-42d7d902-a56c-420c-a740-3aacc7dc922c.gif)

<details>
<summary>code</summary>

```go
package main

import (
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/spinner"
	"time"
)

func main() {
	_ = inf.NewSpinner(
		spinner.WithShape(components.Dot),
		spinner.WithFunc(func(spinner *spinner.Spinner) {
			for i := 0; i < 10; i++ {
				time.Sleep(time.Millisecond * 100)
				spinner.Refreshf("hello world %d", i)
			}
			spinner.Finish("finish")
			spinner.Refresh("is finish?")
		}),
	).Display()
	time.Sleep(time.Millisecond * 100 * 15)
}
```

</details>

---

### Input text

![demo](https://user-images.githubusercontent.com/65269574/183075959-031a068d-6f88-40a0-8b5e-f3d5bba481af.gif)

<details>
<summary>code</summary>

```go
package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/input/text"
	"github.com/fzdwx/infinite/theme"
)

func main() {

	i := inf.NewText(
		text.WithPrompt("what's your name? "),
		text.WithPromptStyle(theme.DefaultTheme.PromptStyle),
		text.WithPlaceholder(" fzdwx (maybe)"),
	)

	_ = i.Display()

	fmt.Printf("you input: %s\n", i.Value())
}
```

</details>

---

### Confirm

![demo](https://user-images.githubusercontent.com/65269574/183076452-5fa73013-42de-47df-97b4-7be743d074c1.gif)

<details>
<summary>code</summary>

```go
package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/input/confirm"
)

func main() {

	c := inf.NewConfirm(
		confirm.WithDefaultYes(),
		confirm.WithDisplayHelp(),
	)

	c.Display()

	if c.Value() {
		fmt.Println("yes, you are.")
	} else {
		fmt.Println("no,you are not.")
	}
}
```

</details>

[full examples](https://github.com/fzdwx/infinite/tree/main/_examples)

## Build with

- https://github.com/charmbracelet/bubbletea
- https://github.com/charmbracelet/bubbles
- https://github.com/charmbracelet/lipgloss
- https://github.com/muesli/termenv
- https://github.com/sahilm/fuzzy
- ...

[full dependencies](https://github.com/fzdwx/infinite/network/dependencies)

## License

MIT