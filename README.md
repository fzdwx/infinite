<div align="center">
<h1>infinite</h1>
<span>:art: <code>infinite</code> 的目标是成为一个易于使用、定制能力强的 tui 组件库</span>
</div>
<br>
<br>

`infinite` 提供一些开箱即用的组件,例如:

- [multi select](https://github.com/fzdwx/infinite#multi-select)
- [single select](https://github.com/fzdwx/infinite#single-select)
- [spinner](https://github.com/fzdwx/infinite#spinner)
- [input text](https://github.com/fzdwx/infinite#input-text)
- [confirm](https://github.com/fzdwx/infinite#confirm)

同时也提供一些更基础的,你可以将它们组装起来使用:

- [input#component](https://github.com/fzdwx/infinite/blob/main/components/input/component.go)
- [selection#component](https://github.com/fzdwx/infinite/blob/main/components/selection/component.go)
- [spinner#component](https://github.com/fzdwx/infinite/blob/main/components/spinner/component.go)

比如说`multi select` 和 `single select` 的基础组件就是 `selection#component`、
`input text` 和 `confirm` 是基于`input#component`.

## install

```bash
go get github.com/fzdwx/infinite
```

## showcase

### multi select

[source code](https://github.com/fzdwx/infinite/blob/main/examples/mutil_select/main.go)

![demo](https://user-images.githubusercontent.com/65269574/183073869-7de79068-0d52-46d2-84aa-4a5130df5634.gif)

### single select

[source code](https://github.com/fzdwx/infinite/blob/main/examples/single_select/main.go)

![demo](https://user-images.githubusercontent.com/65269574/183074455-b09f747f-8f18-4d5e-8286-61d7c9bb963d.gif)

### spinner

[source code](https://github.com/fzdwx/infinite/blob/main/examples/spinner/main.go)

![demo](https://user-images.githubusercontent.com/65269574/183074665-42d7d902-a56c-420c-a740-3aacc7dc922c.gif)

### input text

[source code](https://github.com/fzdwx/infinite/blob/main/examples/input/main.go)

![demo](https://user-images.githubusercontent.com/65269574/183075959-031a068d-6f88-40a0-8b5e-f3d5bba481af.gif)

### confirm

[source code](https://github.com/fzdwx/infinite/blob/main/examples/confirm/main.go)

![demo](https://user-images.githubusercontent.com/65269574/183076452-5fa73013-42de-47df-97b4-7be743d074c1.gif)

## build with

- https://github.com/charmbracelet/bubbletea
- https://github.com/charmbracelet/bubbles
- https://github.com/charmbracelet/lipgloss
- https://github.com/muesli/termenv

## License

MIT