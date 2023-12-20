# infinite

![gif](https://user-images.githubusercontent.com/65269574/184916069-076a0f6a-70bd-49e1-b7d7-0d2e7fc5c6bb.gif)

Help you to create interactive command line applications in Go.

## Features

- multi/single select
- progress-bar group
- spinner
- confirm(input/selection)
- input

## Install

```shell
go get github.com/fzdwx/infinite@main
```

## Examples


```go
func main() {
	input := components.NewInput()
	input.Prompt = "Filtering: "
	input.PromptStyle = style.New().Bold().Italic().Fg(color.LightBlue)

	keymap := components.DefaultMultiKeyMap()
	keymap.Choice = key.NewBinding(
		key.WithKeys(tea.KeySpace.String()),
	)
	_, _ = inf.NewMultiSelect([]string{
		"a", "b", "c",
		"d", "e",
		"f",
		"g",
		"h",
	},
		multiselect.WithKeyMap(keymap),
		multiselect.WithHintSymbol("x"),
		multiselect.WithUnHintSymbol("√"),
		multiselect.WithPageSize(3),
		multiselect.WithFilterInput(input),
	).
		Display("select your items!")
}
```

More: https://github.com/fzdwx/infinite/tree/main/_examples

## License

MIT
