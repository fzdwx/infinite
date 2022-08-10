<div align="center">
<h1>infinite</h1>
<span>🌊 A component library for developing interactive CLI(tui,terminal) programs.</span>
<br>
</div>
<br>
<img src="https://user-images.githubusercontent.com/65269574/183641765-e8de7441-3c4e-4008-b2a9-b2ba556ddd72.gif" alt="demo">

## Features

- Provides a set of out-of-the-box components
    - autocomplete
    - progress-bar group
    - multi/single select
    - spinner
    - confirm
    - input
- Support window/linux (I only have these two operating systems)
- Customizable (You can replace the implementation of some methods in the component)
- Combinable (You can combine multiple components to use)
    - `autocomplete` is composed of `input` and `select`
    - `select` implements the filter function by embedding the `input`
    - ...

## Install

```bash
go get github.com/fzdwx/infinite
```

## Get Started

todo

## Build with

[dependencies](https://github.com/fzdwx/infinite/network/dependencies)

## License

MIT