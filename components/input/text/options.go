package text

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/style"
	"time"
)

type Option func(i *Text)

// WithQuitKey bindQuitKey,default is "enter"
func WithQuitKey(quitKey key.Binding) Option {
	return func(i *Text) {
		i.inner.QuitKey = quitKey
	}
}

// WithPrompt set the prompt
func WithPrompt(prompt string) Option {
	return func(i *Text) {
		i.inner.Prompt = prompt
	}
}

// WithPlaceholder set the placeholder
func WithPlaceholder(placeholder string) Option {
	return func(i *Text) {
		i.inner.Placeholder = placeholder
	}
}

// WithBlinkSpeed set the blink speed
func WithBlinkSpeed(blinkSpeed time.Duration) Option {
	return func(i *Text) {
		i.inner.BlinkSpeed = blinkSpeed
	}
}

// WithEchoMode sets the input behavior of the text input field.
func WithEchoMode(echoMode components.EchoMode) Option {
	return func(i *Text) {
		i.inner.EchoMode = echoMode
	}
}

// WithEchoCharacter setthe echo char shape
func WithEchoCharacter(echoCharacter rune) Option {
	return func(i *Text) {
		i.inner.EchoCharacter = echoCharacter
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

// WithCursorStyle setthe cursor style
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
