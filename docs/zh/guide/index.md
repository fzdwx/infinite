# 简介

[Infinite](https://github.com/fzdwx/infinite) 是一个用于开发交互式 CLI(tui,terminal) 程序的组件库，你可以用它在`Golang`
中快速的完成一个好看且实用的 CLI 程序。

![demo](https://user-images.githubusercontent.com/65269574/184916069-076a0f6a-70bd-49e1-b7d7-0d2e7fc5c6bb.gif)

## Features

- 提供一系列开箱即用的组件
    - autocomplete
    - progress-bar group
    - multi/single select
    - spinner
    - confirm(input/selection)
    - [text](https://fzdwx.github.io/infinite/zh/components/input.html#input-text)
- 跨平台
- 可定制,你可以替换组件中的某些选项或方法为你自己的实现
    - 通过修改 `multi select`
      的[某些选项](https://github.com/fzdwx/infinite/blob/main/components/selection/singleselect/single_select.go#L49)
      实现 `single select`
- 可组合,你可以将一个或多个基础组件联合在一起使用
    - `autocomplete` 通过 `input` 来实现输入接收，通过 `selection` 来实现待选项的选择.
    - `selection` 通过嵌入 `input` 来实现过滤功能.


## License

MIT

::: details
MIT License

Copyright (c) 2022 [fzdwx](https://github.com/fzdwx) <likelovec@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
:::