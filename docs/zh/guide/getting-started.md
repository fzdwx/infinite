# 快速开始  <a href="https://github.com/fzdwx/infinite/releases"><img style="display: inline" src="https://img.shields.io/github/v/release/fzdwx/infinite.svg" alt="release"></a>

::: tip
`infinite` 依赖于 go 1.18.
:::

## Step.1: 创建一个新的项目

创建并进入新目录:

```shell
mkdir infinite-demo && cd infinite-demo
```

使用 `go mod` 初始化项目:

```shell
go mod init infinite-demo
```

## Step.2: 安装`infinite`

添加`infinite`作为项目的依赖:

```shell
go get github.com/fzdwx/infinite
```

## Step.3: 一个简单的 `confirm` 使用 demo

新建一个`main.go`文件，并将下面代码复制到文件中:

```go
package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/selection/confirm"
)

func main() {

	val, _ := inf.NewConfirmWithSelection(
		confirm.WithDefaultYes(),
	).Display()

	if val {
		fmt.Println("yes, you are.")
	} else {
		fmt.Println("no,you are not.")
	}
}
```

运行这个项目:

```shell
go run .
```