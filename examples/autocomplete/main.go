package main

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/fzdwx/infinite/components"
	"github.com/sahilm/fuzzy"
	"sort"
)

func main() {

	suggesterOptions := []string{
		"hello",
		"world",
		"zzzz",
		"hello 2",
		"java",
		"Java",
		"rust",
		"jvav",
		"golang",
	}

	c := components.NewAutocomplete(func(valCtx components.AutocompleteValCtx) ([]string, bool) {
		matches := fuzzy.Find(valCtx.CursorWord(), suggesterOptions)
		if len(matches) == 0 {
			return nil, false
		}

		sort.Stable(matches)

		suggester := slice.Map[fuzzy.Match, string](matches, func(index int, item fuzzy.Match) string {
			return suggesterOptions[item.Index]
		})

		return suggester, true
	})

	components.NewStartUp(c).Start()
}
