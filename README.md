<div align="center">
<h1>infinite</h1>
<span>🧬 用于开发交互式 CLI(tui,terminal) 程序的组件库.</span>
<br>
<a href="https://goreportcard.com/report/github.com/fzdwx/infinite"><img src="https://goreportcard.com/badge/github.com/fzdwx/infinite" alt="go report card"></a>
<a href="https://github.com/fzdwx/infinite/releases"><img src="https://img.shields.io/github/v/release/fzdwx/infinite.svg?style=flat-square" alt="release"></a>
</div>
<img src="https://user-images.githubusercontent.com/65269574/183641765-e8de7441-3c4e-4008-b2a9-b2ba556ddd72.gif" alt="demo">

中文 | [English](https://fzdwx.github.io/infinite/en/)

## Features

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

## 文档

https://fzdwx.github.io/infinite/

## License

MIT