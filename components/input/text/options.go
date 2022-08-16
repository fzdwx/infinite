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

		if len(maskedSymbol) < 0 {
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

// WithBackgroundStyle set the background style
func WithBackgroundStyle(style *style.Style) Option {
	return func(i *Text) {
		i.inner.BackgroundStyle = style
	}
}

// WithPlaceholderStyle set the placeholder style
func WithPlaceholderStyle(style *style.Style) Option {
	return func(i *Text) {
		i.inner.PlaceholderStyle = style
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
