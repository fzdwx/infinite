package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/duke-git/lancet/v2/slice"
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
		Model:                    textinput.New(),
		Required:                 InputDefaultRequired,
		RequiredMsg:              InputDefaultRequiredMsg,
		RequiredMsgKeepAliveTime: InputDefaultRequiredMsgKeepTime,
		Status:                   InputDefaultStatus,
		Prompt:                   InputDefaultPrompt,
		DefaultValue:             InputDefaultValue,
		BlinkSpeed:               InputDefaultBlinkSpeed,
		EchoMode:                 InputDefaultEchoMode,
		EchoCharacter:            InputDefaultEchoCharacter,
		CharLimit:                InputDefaultCharLimit,
		KeyMap:                   InputDefaultKeyMap,
		PlaceholderStyle:         InputDefaultPlaceholderStyle,
		PromptStyle:              InputDefaultPromptStyle,
		TextStyle:                InputDefaultTextStyle,
		BackgroundStyle:          InputDefaultBackgroundStyle,
		CursorStyle:              InputDefaultCursorStyle,
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
		Total:           ProgressDefaultTotal,
		Current:         ProgressDefaultCurrent,
		PercentAgeFunc:  ProgressDefaultPercentAgeFunc,
		PercentAgeStyle: ProgressDefaultPercentAgeStyle,
		Width:           ProgressDefaultWidth,
		Full:            ProgressDefaultFull,
		FullColor:       ProgressDefaultFullColor,
		Empty:           ProgressDefaultEmpty,
		EmptyColor:      ProgressDefaultEmptyColor,
		ShowPercentage:  ProgressDefaultShowPercentage,
		ShowCost:        ProgressDefaultShowCost,
		prevAmount:      ProgressDefaultPrevAmount,
		CostView:        ProgressDefaultCostView,
		TickCostDelay:   ProgressDefaultTickCostDelay,
		Quit:            ProgressDefaultQuit,
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
		CursorSymbol:        SelectionDefaultCursorSymbol,
		UnCursorSymbol:      SelectionDefaultUnCursorSymbol,
		CursorSymbolStyle:   SelectionDefaultCursorSymbolStyle,
		ChoiceTextStyle:     SelectionDefaultChoiceTextStyle,
		Prompt:              SelectionDefaultPrompt,
		PromptStyle:         SelectionDefaultPromptStyle,
		HintSymbol:          SelectionDefaultHintSymbol,
		HintSymbolStyle:     SelectionDefaultHintSymbolStyle,
		UnHintSymbol:        SelectionDefaultUnHintSymbol,
		UnHintSymbolStyle:   SelectionDefaultUnHintSymbolStyle,
		confirmed:           SelectionDefaultConfirmed,
		DisableOutPutResult: SelectionDefaultDisableOutPutResult,
		PageSize:            SelectionDefaultPageSize,
		Keymap:              SelectionDefaultKeymap,
		Help:                SelectionDefaultHelp,
		RowRender:           SelectionDefaultRowRender,
		EnableFilter:        SelectionDefaultEnableFilter,
		FilterInput:         SelectionDefaultFilterInput,
		FilterFunc:          SelectionDefaultFilterFunc,
		ShowHelp:            SelectionDefaultShowHelp,
	}

	return c
}

// NewSpinner constructor
func NewSpinner() *Spinner {
	c := &Spinner{
		Model:               SpinnerDefaultModel,
		Shape:               SpinnerDefaultShape,
		ShapeStyle:          SpinnerDefaultShapeStyle,
		Prompt:              SpinnerDefaultPrompt,
		DisableOutPutResult: SpinnerDefaultDisableOutPutResult,
		Status:              SpinnerDefaultStatus,
		Quit:                SpinnerDefaultQuitKey,
	}
	return c
}
