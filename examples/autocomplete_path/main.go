package main

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/fzdwx/infinite/components"
	"github.com/sahilm/fuzzy"
	"io/fs"
	"io/ioutil"
	"sort"
)

func main() {
	var f components.Suggester = func(cursorVal, currentWord string) ([]string, bool) {
		//inputFile, err := os.Open(currentWord)
		//if err != nil {
		//	return nil, false
		//}
		//stat, err := inputFile.Stat()
		//if err != nil {
		//	return nil, false
		//}

		var (
			fileInfos []fs.FileInfo
			err       error
		)

		//if stat.IsDir() {
		//fileInfos, err = ioutil.ReadDir(currentWord)
		//} else {
		fileInfos, err = ioutil.ReadDir(".")
		//}

		if err != nil {
			return nil, false
		}

		var ops []string
		for _, info := range fileInfos {
			ops = append(ops, info.Name())
		}

		matches := fuzzy.Find(currentWord, ops)
		if len(matches) == 0 {
			return nil, false
		}

		sort.Stable(matches)

		suggester := slice.Map[fuzzy.Match, string](matches, func(index int, item fuzzy.Match) string {
			return ops[item.Index]
		})
		return suggester, true
	}

	c := components.NewAutocomplete(f)

	components.NewStartUp(c).Start()
}
