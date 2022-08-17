# 示例

::: tip
这里是一些示例,可以帮助你快速了解如何使用各个组件。
:::

## autocomplete path

一个基于 `autocomplete` 的路径补全器。

![demo](https://user-images.githubusercontent.com/65269574/184916654-999cd99d-94bf-4bd8-8d2c-87d547ec20d7.gif)

::: details 代码

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

:::

## progress-bar group

使用 `progress-bar group` 同时运行多个 `progress-bar`。

![demo](https://user-images.githubusercontent.com/65269574/184917598-9ab058a3-30cd-4a4e-ba72-45d138e6b5b5.gif)

::: details 代码

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

:::

## download file

一个使用 `progress bar` 下载文件的demo
![demo](https://user-images.githubusercontent.com/65269574/185172321-d68a1754-7125-45ed-8239-6913c12c94ca.gif)

:::details 代码

```go
package main

import (
	"flag"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/progress"
	"net/http"
	"os"
	"path"
)

var urlF = flag.String("d", "", "download url")

func init() {
	flag.Parse()
}

func main() {
	url := *urlF
	progress.NewGroupWithCount(1).
		AppendRunner(func(pro *components.Progress) func() {
			resp, err := http.Get(url)
			if err != nil {
				pro.Println("get error", err)
				resp.Body.Close()
				return func() {}
			}
			pro.WithTotal(resp.ContentLength)

			return func() {
				defer resp.Body.Close()

				dest, err := os.OpenFile(path.Base(url), os.O_CREATE|os.O_WRONLY, 0o777)
				defer dest.Close()
				if err != nil {
					pro.Println("dest open error", err)
					return
				}

				_, err = progress.StartTransfer(resp.Body, dest, pro)
				if err != nil {
					pro.Println("trans error", err)
				}
			}
		}).Display()
}
```
:::

## multiple select

一个多选的示例。

![demo](https://user-images.githubusercontent.com/65269574/184917889-b24c8777-f142-4b56-bcf0-d1042ef846d2.gif)

:::details 代码

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

:::

## spinner

使用 `spinner` 以及在运行中刷新输入提示文字。

![demo](https://user-images.githubusercontent.com/65269574/184918112-419df5b7-f4f8-44ff-b421-c65841a4e5c7.gif)

:::details 代码

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

:::

## input text

一个文本输入框。

![demo](https://user-images.githubusercontent.com/65269574/184918464-96194014-0063-48bf-85f3-e0410bdaaba6.gif)

:::details 代码

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

:::

## confirm

### with input

使用 `input` 实现的 `confirm`

![demo](https://user-images.githubusercontent.com/65269574/184920302-9c9c2cfd-4ca7-49d8-9192-8487b2832b36.gif)

:::details 代码

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

:::

### with selection

使用 `selection` 实现的 `confirm`

![demo](https://user-images.githubusercontent.com/65269574/184919493-46a36849-d034-4677-92d0-d4bca15f7ac5.gif)

:::details 代码

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

:::