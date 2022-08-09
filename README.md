<div align="center">
<h1>infinite</h1>
<span>:art: <code>infinite</code> 的目标是成为一个易于使用、定制能力强的 tui 组件库</span>
</div>
<br>
<img src="https://user-images.githubusercontent.com/65269574/183274216-d2a7af91-0581-4d13-b8c2-00b9aad5ef3a.gif">

## TOC

<!-- TOC -->

* [Install](#install)
* [Components](#components)
    * [Selection](#selection)
    * [Input](#input)
    * [Spinner](#spinner)
    * [Progress](#progress)
* [How to use](#how-to-use)
    * [Progress](#progress)
    * [Multi select](#multi-select)
    * [Single select](#single-select)
    * [Spinner](#spinner)
    * [Input text](#input-text)
    * [Confirm](#confirm)
* [Build with](#build-with)
* [License](#license)

<!-- TOC -->

## Install

required: `go 1.19`

```bash
go get github.com/fzdwx/infinite
```

## Components

这里是`infinite`提供的一些组件,你可以把它们变成你的形状.

### Autocomplete

这个一个自动补全的组件

![img.png](img.png)

### Selection

选择器,可更改的选项:

1. 一些提示符的自定义,比如说光标的形状,选中/未选中符号...
2. 键盘映射
3. 每一行输出结果
4. Filter

比如说自定义2和3,我们就可以实现一个`multi select` -> `single select`.

### Input

文本输入框

1. 支持修改光标的形状
2. 输入模式(正常,密码框)
3. 一个`Quit key`,默认是关闭的
4. 一些提示符自定义

### Spinner

微调控制项,什么翻译,就是可以根据定义的一系列的服务,周期性的变换.
除了这个基本功能外,它还可以在刷新提示文字.

### Progress

进度条

## How to use

这里是`infinite`提供的一些开箱即用的组件,它们都是基于上面的基础组件封装而来.

### Progress

这是一个进度条组件的使用案例

<details>
<summary>代码</summary>

```go
package main

import (
	"fmt"
	"github.com/fzdwx/infinite/components"
	"time"
)

func main() {

	var total = 10
	progress := components.NewProgress().
		WithTotal(int64(total)).
		WithDefaultGradient().
		WithPercentAgeFunc(func(total int64, current int64, percent float64) string {
			return fmt.Sprintf(" %d/%d", current, total)
		})

	startUp := components.NewStartUp(progress)
	go func() {
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

		startUp.Kill()
	}()

	startUp.Start()
}

func sleep() {
	time.Sleep(time.Millisecond * 100)
}
```

</details>

![demo](https://user-images.githubusercontent.com/65269574/183285338-fb930b31-91c3-4be2-9068-a36ec51a5f23.gif)

### Progress group

这是一个progress group,它可以管理多个progress.当所有的progress运行完了之后它会自动停止.具体原理可以查看`progressUpdater`
这个结构体和`Group`的`done`参数.

<details>
<summary>代码</summary>

```go
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

```

</details>

![demo](https://user-images.githubusercontent.com/65269574/183296585-b0a56827-d9d9-4258-ad32-266ada01b1ed.gif)

### Multi select

基于`Selection`而来的,基本没有做什么改动.
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
		multiselect.WithHintSymbol("x"),
		multiselect.WithUnHintSymbol("√"),
		multiselect.WithFilterInput(input),
	).Display("select your items!")
}
```

</details>

![demo](https://user-images.githubusercontent.com/65269574/183274216-d2a7af91-0581-4d13-b8c2-00b9aad5ef3a.gif)

### Single select

基于`Multi select`而来的,它减少了一些不要的`option`,以及修改了一些`key mapping`.
具体可以看`singleselect#mapMultiToSingle`这个方法
<details>
<summary>代码</summary>

```go
package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/selection/singleselect"
	"github.com/rotisserie/eris"
)

func main() {
	options := []string{
		"1 Buy carrots",
		"2 Buy celery",
		"3 Buy kohlrabi",
		"4 Buy computer",
		"5 Buy something",
		"6 Buy car",
		"7 Buy subway",
	}
	selected, err := inf.NewSingleSelect(
		options,
		singleselect.WithDisableFilter(),
	).Display("Hello world")
	if err != nil {
		fmt.Println(eris.ToString(err, true))
		return
	}

	fmt.Printf("you selection %s\n", options[selected])
}
```

</details>

![demo](https://user-images.githubusercontent.com/65269574/183074455-b09f747f-8f18-4d5e-8286-61d7c9bb963d.gif)

### Spinner

基于`Spinner`而来的,基本没有做什么改动.
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

![demo](https://user-images.githubusercontent.com/65269574/183074665-42d7d902-a56c-420c-a740-3aacc7dc922c.gif)

### Input text

基于`Input`而来的,基本没有做什么改动.
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

![demo](https://user-images.githubusercontent.com/65269574/183075959-031a068d-6f88-40a0-8b5e-f3d5bba481af.gif)

### Confirm

基于`Input`而来的,根据`Confirm`的使用场景做了一些适配,它只关注`Quit`,`Yes`,`No`这3个事件.
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

![demo](https://user-images.githubusercontent.com/65269574/183076452-5fa73013-42de-47df-97b4-7be743d074c1.gif)

## Build with

- https://github.com/charmbracelet/bubbletea
- https://github.com/charmbracelet/bubbles
- https://github.com/charmbracelet/lipgloss
- https://github.com/muesli/termenv

## License

MIT