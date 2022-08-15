package components

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/fzdwx/infinite/style"
	"github.com/fzdwx/infinite/theme"
)

type (
	/*
		Components, You can use these components directly:
			 	1. Input
			 	2. Selection
			 	3. Spinner
				4. Autocomplete
				5. Progress
		Or use them inline in your custom component,
		for how to embed them, you can refer to the implementation of `Confirm`.
	*/
	Components interface {
		tea.Model

		// SetProgram this method will be called back when the tea.Program starts.
		// please keep passing this method
		SetProgram(program *tea.Program)
	}
)

// NewAutocomplete constructor
func NewAutocomplete(suggester Suggester) *Autocomplete {
	return &Autocomplete{
		Suggester:            suggester,
		Completer:            DefaultCompleter(),
		Input:                NewInput(),
		KeyMap:               DefaultAutocompleteKeyMap(),
		ShowSelection:        true,
		ShouldNewSelection:   true,
		SelectionCreator:     DefaultSelectionCreator,
		SuggestionViewRender: NewLineSuggestionRender,
		//SuggestionViewRender: TabSuggestionRender,
	}
}

// NewInput constructor
func NewInput() *Input {
	c := &Input{
		Model:            textinput.New(),
		Status:           InputDefaultStatus,
		Prompt:           InputDefaultPrompt,
		DefaultValue:     InputDefaultValue,
		BlinkSpeed:       InputDefaultBlinkSpeed,
		EchoMode:         InputDefaultEchoMode,
		EchoCharacter:    InputDefaultEchoCharacter,
		CharLimit:        InputDefaultCharLimit,
		QuitKey:          InputDefaultQuitKey,
		PlaceholderStyle: InputDefaultPlaceholderStyle,
		PromptStyle:      InputDefaultPromptStyle,
		TextStyle:        InputDefaultTextStyle,
		BackgroundStyle:  InputDefaultBackgroundStyle,
		CursorStyle:      InputDefaultCursorStyle,
	}
	return c
}

// NewPrintHelper constructor
func NewPrintHelper(program *tea.Program) *PrintHelper {
	return &PrintHelper{program: program}
}

// NewProgress constructor
func NewProgress() *Progress {
	p := &Progress{
		Id:              nextID(),
		Total:           100,
		Current:         0,
		PercentAgeFunc:  DefaultPercentAgeFunc,
		PercentAgeStyle: style.New().Inline(),
		Width:           defaultWidth,
		Full:            '█',
		FullColor:       "#7571F9",
		Empty:           '░',
		EmptyColor:      "#606060",
		ShowPercentage:  true,
		ShowCost:        true,
		prevAmount:      0,
		CostView:        DefaultCostView,
		TickCostDelay:   defaultTicKCostDelay,
	}

	return p
}

// NewSelection constructor
func NewSelection(choices []string) *Selection {

	items := slice.Map[string, SelectionItem](choices, func(idx int, item string) SelectionItem {
		return SelectionItem{idx, item}
	})

	c := &Selection{
		Choices:             items,
		Selected:            make(map[int]struct{}),
		CursorSymbol:        ">",
		UnCursorSymbol:      " ",
		CursorSymbolStyle:   theme.DefaultTheme.CursorSymbolStyle,
		ChoiceTextStyle:     theme.DefaultTheme.ChoiceTextStyle,
		Prompt:              "Please Selection your options:",
		PromptStyle:         theme.DefaultTheme.PromptStyle,
		HintSymbol:          "✓",
		HintSymbolStyle:     theme.DefaultTheme.MultiSelectedHintSymbolStyle,
		UnHintSymbol:        "✗",
		UnHintSymbolStyle:   theme.DefaultTheme.UnHintSymbolStyle,
		quited:              false,
		DisableOutPutResult: false,
		PageSize:            5,
		Keymap:              DefaultMultiKeyMap,
		Help:                help.New(),
		RowRender:           DefaultRowRender,
		EnableFilter:        true,
		FilterInput:         NewInput(),
		FilterFunc:          DefaultFilterFunc,
		ShowHelp:            true,
	}

	return c
}

// NewSpinner constructor
func NewSpinner() *Spinner {
	c := &Spinner{
		Model:               spinner.New(),
		Shape:               Line,
		ShapeStyle:          theme.DefaultTheme.SpinnerShapeStyle,
		Prompt:              "Loading...",
		DisableOutPutResult: false,
		Status:              Normal,
	}
	return c
}
