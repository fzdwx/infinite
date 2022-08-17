# Getting started <a href="https://github.com/fzdwx/infinite/releases"><img style="display: inline" src="https://img.shields.io/github/v/release/fzdwx/infinite.svg" alt="release"></a>

::: tip
`infinite` depends on go 1.18.
:::

## Step.1: Create a new project

Create and enter a new directory:

```shell
mkdir infinite-demo && cd infinite-demo
```

Initialize the project with `go mod`:

```shell
go mod init infinite-demo
```

## Step.2: Install `infinite`

Add `infinite` as a dependency of the project:

```shell
go get github.com/fzdwx/infinite
```

## Step.3: A simple `confirm` using demo

Create a new `main.go` file and copy the following code into the file:

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

Run this project:

```shell
go run .
```