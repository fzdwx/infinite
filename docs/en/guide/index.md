# 简介

[Infinite](https://github.com/fzdwx/infinite) is a component library for developing interactive CLI(tui,terminal)
programs,You can use it in `Golang` Quickly complete a good-looking and practical CLI program in .

![demo](https://user-images.githubusercontent.com/65269574/184916069-076a0f6a-70bd-49e1-b7d7-0d2e7fc5c6bb.gif)

## Features

- Provides a set of out-of-the-box components
    - autocomplete
    - progress-bar group
    - multi/single select
    - spinner
    - confirm(input/selection)
    - [input text](https://fzdwx.github.io/infinite/zh/components/input.html#input-text)
- Cross-platform
  -Customizable, you can replace some options or methods in the component for your own implementation
    - By modifying  `multi select`
      the[some options](https://github.com/fzdwx/infinite/blob/main/components/selection/singleselect/single_select.go#L49)
      implement `single select`
- Composable, you can combine one or more base components together
    - `autocomplete` implements input reception through `input`, and `selection` implements the selection of options.
    - `selection` implements filtering by embedding `input`.

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