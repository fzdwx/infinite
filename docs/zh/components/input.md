# input

:::tip
input 是一个用于接收用户输入的基础组件。
:::

## original demo

```go
package main

import (
	"fmt"
	"github.com/fzdwx/infinite/components"
)

func main() {
	input := components.NewInput()

	if err := components.NewStartUp(input).Start(); err != nil {
		panic(err)
	}

	fmt.Println("Get:", input.Value())
}
```