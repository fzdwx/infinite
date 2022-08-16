<div align="center">
<h1>infinite</h1>
<span>🧬 用于开发交互式 CLI(tui,terminal) 程序的组件库.</span>
<br>
<a href="https://goreportcard.com/report/github.com/fzdwx/infinite"><img src="https://goreportcard.com/badge/github.com/fzdwx/infinite" alt="go report card"></a>
<a href="https://github.com/fzdwx/infinite/releases"><img src="https://img.shields.io/github/v/release/fzdwx/infinite.svg?style=flat-square" alt="release"></a>
</div>
<img src="https://user-images.githubusercontent.com/65269574/184916069-076a0f6a-70bd-49e1-b7d7-0d2e7fc5c6bb.gif" alt="demo">

中文 | [English](https://fzdwx.github.io/infinite/en/)

## 特性

- 提供一系列开箱即用的组件
    - autocomplete
    - progress bar / progress-bar group
    - multi/single select
    - spinner
    - confirm
    - input
- 支持 window/linux (我现在只有这两种操作系统)
- 可定制,你可以替换组件中的某些选项或方法为你自己的实现
- 可组合,你可以将一个或多个基础组件联合在一起使用
    - `autocomplete` 由`input` 和 `selection` 组成
    - `selection` 通过嵌入`input` 来实现过滤功能.
    - ...

## 最佳实践

1. 通过消息来更新状态,也就是通过`program.Send(msg)`来发送消息,`Update`监听并进行状态更新,最后通过`View`来反馈结果.
2. ...

## 安装

```bash
go get github.com/fzdwx/infinite
```

## 详细文档

https://fzdwx.github.io/infinite/

## 使用案例

### Combined demo

![demo](https://user-images.githubusercontent.com/65269574/184917224-5048b5c4-e063-4379-a00c-e28388e13c41.gif)

<details>
<summary>代码</summary>

```go
package main

import (
	"errors"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/pkg/strx"
	"time"
)

func main() {
	total := 10
	spinner := components.NewSpinner()
	spinner.Prompt = strx.Space + spinner.Prompt
	progress := components.NewProgress().WithTotal(int64(total))

	NewComponent(spinner, progress).Display(func(c *Component) {
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
	})
}

type Component struct {
	spinner  *components.Spinner
	progress *components.Progress
	*components.StartUp
}

func NewComponent(spinner *components.Spinner, progress *components.Progress) *Component {
	return &Component{spinner: spinner, progress: progress}
}

func (c *Component) Display(runner func(c *Component)) error {
	c.StartUp = components.NewStartUp(c)
	if runner == nil {
		return errors.New("runner is null")
	}

	go func() {
		runner(c)
		c.progress.Done()
		c.Quit()
	}()

	return c.Start()
}

func (c *Component) Init() tea.Cmd {

	return tea.Batch(c.spinner.Init(), c.progress.Init())
}

func (c *Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return c, tea.Quit
		}
	}
	_, c1 := c.spinner.Update(msg)
	_, c2 := c.progress.Update(msg)

	return c, tea.Batch(c1, c2)
}

func (c *Component) View() string {
	return strx.NewFluent().Write(c.spinner.View()).Space(4).Write(c.progress.View()).String()
}

func (c *Component) SetProgram(program *tea.Program) {
	c.spinner.SetProgram(program)
	c.progress.SetProgram(program)
}

func sleep() {
	time.Sleep(time.Millisecond * 100)
}
```

</details>

### Autocomplete

![demo](https://user-images.githubusercontent.com/65269574/184916654-999cd99d-94bf-4bd8-8d2c-87d547ec20d7.gif)
<details>
<summary>代码</summary>

```go
package main

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/fzdwx/infinite/components"
	"github.com/sahilm/fuzzy"
	"path/filepath"
	"sort"
)

func main() {
	var f components.Suggester = func(valCtx components.AutocompleteValCtx) ([]string, bool) {
		cursorWord := valCtx.CursorWord()
		files, err := filepath.Glob(cursorWord + "*")
		if err != nil {
			return nil, false
		}

		matches := fuzzy.Find(cursorWord, files)
		if len(matches) == 0 {
			return nil, false
		}

		sort.Stable(matches)

		suggester := slice.Map[fuzzy.Match, string](matches, func(index int, item fuzzy.Match) string {
			return files[item.Index]
		})
		return suggester, true
	}

	c := components.NewAutocomplete(f)

	components.NewStartUp(c).Start()
}

```

</details>

### Progress group

![demo](https://user-images.githubusercontent.com/65269574/184917598-9ab058a3-30cd-4a4e-ba72-45d138e6b5b5.gif)

<details>
<summary>代码</summary>

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

### Multiple select

![demo](https://user-images.githubusercontent.com/65269574/184917889-b24c8777-f142-4b56-bcf0-d1042ef846d2.gif)

<details>
<summary>代码</summary>

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

### Spinner

![demo](https://user-images.githubusercontent.com/65269574/184918112-419df5b7-f4f8-44ff-b421-c65841a4e5c7.gif)

<details>
<summary>代码</summary>

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
		//spinner.WithDisableOutputResult(),
	).Display(func(spinner *spinner.Spinner) {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 100)
			spinner.Refreshf("hello world %d", i)
		}

		spinner.Finish("finish")

		spinner.Refresh("is finish?")
	})

	time.Sleep(time.Millisecond * 100 * 15)
}
```

</details>

### Input text

![demo](https://user-images.githubusercontent.com/65269574/184918464-96194014-0063-48bf-85f3-e0410bdaaba6.gif)

<details>
<summary>代码</summary>

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

### Confirm with Input

![demo](https://user-images.githubusercontent.com/65269574/184919493-46a36849-d034-4677-92d0-d4bca15f7ac5.gif)

<details>
<summary>代码</summary>

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

### Confirm With Selection

![demo](https://user-images.githubusercontent.com/65269574/184919493-46a36849-d034-4677-92d0-d4bca15f7ac5.gif)

<details>
<summary>代码</summary>

```go
package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
)

func main() {

	val, _ := inf.NewConfirmWithSelection(
		//confirm.WithDisOutResult(),
	).Display()

	fmt.Println(val)
}
```

</details>

[所有示例](https://github.com/fzdwx/infinite/tree/main/_examples)

## 依赖

- https://github.com/charmbracelet/bubbletea
- https://github.com/charmbracelet/bubbles
- https://github.com/charmbracelet/lipgloss
- https://github.com/muesli/termenv
- https://github.com/sahilm/fuzzy
- ...

[所有依赖](https://github.com/fzdwx/infinite/network/dependencies)

## 开源协议

MIT