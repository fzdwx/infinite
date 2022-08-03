package singleselect

import (
	"fmt"
	"github.com/fzdwx/infinite/multiselect"
)

type SingleSelect struct {
	inner *multiselect.MultiSelect
}

func New(choices []string, ops ...Option) *SingleSelect {
	s := &SingleSelect{inner: multiselect.New(choices)}

	s.inner.Apply(getSingleSelectOps()...)

	s.Apply(ops...)
	return s
}

func getSingleSelectOps() []multiselect.Option {
	var msOps []multiselect.Option

	msOps = append(msOps, multiselect.WithRowRender(func(cursorSymbol string, hintSymbol string, choice string) string {
		return fmt.Sprintf("%s %s", cursorSymbol, choice)
	}))

	msOps = append(msOps, multiselect.WithPrompt("Please select your option:"))
	keyMap := multiselect.DefaultKeyMap
	keyMap.Choice.Enabled()
	msOps = append(msOps, multiselect.WithKeyBinding(keyMap))

	// todo
	return msOps
}

// Apply options on MultiSelect
func (ss *SingleSelect) Apply(ops ...Option) *SingleSelect {
	if len(ops) > 0 {
		for _, option := range ops {
			option(ss)
		}
	}
	return ss
}

type Option func(ss *SingleSelect)

// WithPageSize default is 5
func WithPageSize(pageSize int) Option {
	return func(ss *SingleSelect) {
		ss.inner.Apply(multiselect.WithPageSize(pageSize))
	}
}
