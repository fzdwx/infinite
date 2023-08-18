package text

import (
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/style"
	"time"
)

type Option func(i *Text)

// WithKeyMap replace keymap
//
// components.InputDefaultKeyMap
func WithKeyMap(keymap components.InputKeyMap) Option {
	return func(i *Text) {
		i.inner.KeyMap = keymap
	}
}

// WithRequired must input some words.
func WithRequired() Option {
	return func(i *Text) {
		i.inner.Required = true
	}
}

// WithRequiredMsg if there is no input, the `msg` will be prompted
//
// components.InputDefaultRequiredMsg
func WithRequiredMsg(msg string) Option {
	return func(i *Text) {
		i.inner.RequiredMsg = msg
	}
}

// WithRequiredMsgKeepAliveTime set `requiredMsg` keep alive time.
//
// components.InputDefaultRequiredMsgKeepTime
func WithRequiredMsgKeepAliveTime(keepaliveTime time.Duration) Option {
	return func(i *Text) {
		i.inner.RequiredMsgKeepAliveTime = keepaliveTime
	}
}

// WithPrompt set the prompt
func WithPrompt(prompt string) Option {
	return func(i *Text) {
		i.inner.Prompt = prompt
	}
}

// WithDefaultValue set the default value
func WithDefaultValue(s string) Option {
	return func(i *Text) {
		i.inner.DefaultValue = s
	}
}

// WithDefaultValueRequireValue set the default value required
// default is true
func WithDefaultValueRequireValue(b bool) Option {
	return func(i *Text) {
		i.inner.DefaultValueRequired = b
	}
}

// WithBlinkSpeed set the blink speed
func WithBlinkSpeed(blinkSpeed time.Duration) Option {
	return func(i *Text) {
		i.inner.BlinkSpeed = blinkSpeed
	}
}

// WithEchoNone set echoMode use components.EchoNone.
// displays nothing as characters are entered
func WithEchoNone() Option {
	return func(i *Text) {
		i.inner.EchoMode = components.EchoNone
	}
}

// WithEchoPassword set echoMode use components.EchoPassword.
// if maskedSymbol is not empty, then set EchoCharacter use maskedSymbol[0]
func WithEchoPassword(maskedSymbol ...rune) Option {
	return func(i *Text) {
		i.inner.EchoMode = components.EchoPassword

		if len(maskedSymbol) <= 0 {
			return
		}
		i.inner.EchoCharacter = maskedSymbol[0]
	}
}

// WithPromptStyle set the prompt style
func WithPromptStyle(style *style.Style) Option {
	return func(i *Text) {
		i.inner.PromptStyle = style
	}
}

// WithTextStyle set the text style
func WithTextStyle(style *style.Style) Option {
	return func(i *Text) {
		i.inner.TextStyle = style
	}
}

// WithDefaultValueStyle set the default value style
func WithDefaultValueStyle(style *style.Style) Option {
	return func(i *Text) {
		i.inner.DefaultValueStyle = style
	}
}

// WithCursorStyle set the cursor style
func WithCursorStyle(style *style.Style) Option {
	return func(i *Text) {
		i.inner.CursorStyle = style
	}
}

// WithCharLimit is the maximum amount of characters this input element will
// accept. If 0 or less, there's no limit.
func WithCharLimit(charLimit int) Option {
	return func(i *Text) {
		i.inner.CharLimit = charLimit
	}
}

// WithFocusSymbol default is theme.DefaultTheme#FocusSymbol
func WithFocusSymbol(s string) Option {
	return func(i *Text) {
		i.inner.FocusSymbol = s
	}
}

// WithUnFocusSymbol default is theme.DefaultTheme#UnFocusSymbol
func WithUnFocusSymbol(s string) Option {
	return func(i *Text) {
		i.inner.UnFocusSymbol = s
	}
}

// WithFocusInterval default is theme.DefaultTheme#FocusInterval
func WithFocusInterval(s string) Option {
	return func(i *Text) {
		i.inner.FocusInterval = s
	}
}

// WithUnFocusInterval default is theme.DefaultTheme#UnFocusInterval
func WithUnFocusInterval(s string) Option {
	return func(i *Text) {
		i.inner.UnFocusInterval = s
	}
}

// WithFocusSymbolStyle default is theme.DefaultTheme#FocusSymbolStyle
func WithFocusSymbolStyle(s *style.Style) Option {
	return func(i *Text) {
		i.inner.FocusSymbolStyle = s
	}
}

// WithUnFocusSymbolStyle default is theme.DefaultTheme#UnFocusIntervalStyle
func WithUnFocusSymbolStyle(s *style.Style) Option {
	return func(i *Text) {
		i.inner.UnFocusSymbolStyle = s
	}
}

// WithFocusIntervalStyle default is theme.DefaultTheme#FocusIntervalStyle
func WithFocusIntervalStyle(s *style.Style) Option {
	return func(i *Text) {
		i.inner.FocusIntervalStyle = s
	}
}

// WithUnFocusIntervalStyle default is theme.DefaultTheme#UnFocusIntervalStyle
func WithUnFocusIntervalStyle(s *style.Style) Option {
	return func(i *Text) {
		i.inner.UnFocusIntervalStyle = s
	}
}

// WithDisableOutputResult disable output result
func WithDisableOutputResult() Option {
	return func(i *Text) {
		i.inner.OutputResult = false
	}
}

// WithPure do not use any beautification features,
// any options you customize will be cleared
func WithPure() Option {
	return func(i *Text) {
		i.pure = true
	}
}
