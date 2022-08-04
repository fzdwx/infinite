# infinite

:art: `infinite` 的目标是成为一个易于使用的Tui组件库

`infinite` 提供一些开箱即用的组件,例如:

- [multi select](https://github.com/fzdwx/infinite#multi-select)
- [single select](https://github.com/fzdwx/infinite#single-select)
- [spinner](https://github.com/fzdwx/infinite#spinner)
- [input](https://github.com/fzdwx/infinite#input)

同时也提供一些更基础的,你可以将它们组装起来使用:

- [input.component](https://github.com/fzdwx/infinite/blob/main/components/input/component.go)
- [selection.component](https://github.com/fzdwx/infinite/blob/main/components/selection/component.go)
- [spinner.component](https://github.com/fzdwx/infinite/blob/main/components/spinner/component.go)

比如说`multi select` 和 `single select` 的基础组件就是 `selection.component`.

## install

```go
go get github.com/fzdwx/infinite@latest
```

## showcase

### multi select

[source code](https://github.com/fzdwx/infinite/blob/main/examples/mutil_select/main.go)

![demo](https://user-images.githubusercontent.com/65269574/182607109-c5969485-4a21-4086-8476-bdb361a7e779.gif)

### single select

[source code](https://github.com/fzdwx/infinite/blob/main/examples/single_select/main.go)

![single_select](https://user-images.githubusercontent.com/65269574/182606494-3462614c-3ffc-49de-884c-5cfa8685aed3.gif)

### spinner

[source code](https://github.com/fzdwx/infinite/blob/main/examples/spinner/main.go)

![spinner](https://user-images.githubusercontent.com/65269574/182842629-6c80daab-5bde-467f-9691-ed2731aeb419.gif)

### input

[source code](https://github.com/fzdwx/infinite/blob/main/examples/spinner/main.go)

![input](https://user-images.githubusercontent.com/65269574/182850907-4b1c8e03-e008-40a2-804e-73cbfcba0c70.gif)

## built with
- https://github.com/charmbracelet/bubbletea
- https://github.com/charmbracelet/bubbles
- https://github.com/charmbracelet/lipgloss
- https://github.com/muesli/termenv
## License

MIT