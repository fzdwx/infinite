<div align="center">
<h1>infinite</h1>
<span>🧬 用于开发交互式 CLI(tui,terminal) 程序的组件库.</span>
<br>
<a href="https://goreportcard.com/report/github.com/fzdwx/infinite"><img src="https://goreportcard.com/badge/github.com/fzdwx/infinite" alt="go report card"></a>
<a href="https://github.com/fzdwx/infinite/releases"><img src="https://img.shields.io/github/v/release/fzdwx/infinite.svg?style=flat-square" alt="release"></a>
</div>
<img src="https://user-images.githubusercontent.com/65269574/184916069-076a0f6a-70bd-49e1-b7d7-0d2e7fc5c6bb.gif" alt="demo">

中文 | [English](https://fzdwx.github.io/infinite/en/)

## Features

- 提供一系列开箱即用的组件
    - autocomplete
    - progress-bar group
    - multi/single select
    - spinner
    - confirm(input/selection)
    - text
- 跨平台
- 可定制,你可以替换组件中的某些选项或方法为你自己的实现
    - 通过修改 `multi select`
      的 [某些选项](https://github.com/fzdwx/infinite/blob/main/components/selection/singleselect/single_select.go#L49)
      实现 `single select`
- 可组合,你可以将一个或多个基础组件联合在一起使用
    - `autocomplete` 通过 `input` 来实现输入接收，通过 `selection` 来实现待选项的选择.
    - `selection` 通过嵌入 `input` 来实现过滤功能.

## Install

```shell
go get github.com/fzdwx/infinite
```

## Getting started

https://fzdwx.github.io/infinite/zh/guide/getting-started

## Examples

https://fzdwx.github.io/infinite/zh/guide/examples

## Documentation

https://fzdwx.github.io/infinite/

## License

MIT