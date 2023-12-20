package components

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/duke-git/lancet/v2/slice"
	"testing"
)

func TestFilter(t *testing.T) {
	strings := []string{"Buy carrots",
		"Buy celery",
		"Buy kohlrabi",
		"Buy computer",
		"Buy something",
		"Buy car",
		"Buy subway",
	}
	selection := NewSelection(strings)
	selection.Init()
	selection.Update(tea.KeyMsg{Runes: []rune{'c'}})
	selection.Update(tea.KeyMsg{Runes: []rune{'a'}})
	selection.Update(tea.KeyMsg{Runes: []rune{'r'}})

	selection.choice()

	items := DefaultFilterFunc("car", slice.Map[string, SelectionItem](strings, func(idx int, item string) SelectionItem {
		return SelectionItem{idx, item}
	}))

	got := items[selection.Value()[0]].Val
	want := strings[5]
	if got != want {
		t.Errorf("choice : got %v, want %v\n", got, want)
	}
}

func TestSelection_MoveDone_Choice(t *testing.T) {
	strings := []string{"Buy carrots",
		"Buy celery",
		"Buy kohlrabi",
		"Buy computer",
		"Buy something",
		"Buy car",
		"Buy subway",
	}

	for i := 0; i < 7; i++ {
		t.Run(fmt.Sprintf("move down %d", i), func(t *testing.T) {
			selection := NewSelection(strings)
			selection.Init()
			testMoveDown(selection, i)
			selection.choice()
			selection.choice()
			selection.choice()

			got := strings[selection.Value()[0]]
			want := strings[i]
			if got != want {
				t.Errorf("choice = %v, want %v\n", got, want)
			}
		})
	}
}

func TestSelection_MoveUp_Choice(t *testing.T) {
	strings := []string{"Buy carrots",
		"Buy celery",
		"Buy kohlrabi",
		"Buy computer",
		"Buy something",
		"Buy car",
		"Buy subway",
	}
	size := len(strings) - 1
	for i := 0; i < 7; i++ {
		t.Run(fmt.Sprintf("move up %d", i), func(t *testing.T) {
			selection := NewSelection(strings)
			selection.Init()
			testMoveUp(selection, i, size)
			selection.choice()

			idx := selection.Value()[0]
			got := strings[idx]
			want := strings[size-i]
			if got != want {
				t.Errorf("choice =idx:%d %v, want %v\n", idx, got, want)
			}
		})
	}
}

func testMoveDown(selection *Selection, tiems int) {
	for i := 0; i < tiems; i++ {
		selection.moveDown()
	}
}

func testMoveUp(selection *Selection, tiems, size int) {
	for i := 0; i < size; i++ {
		selection.moveDown()
	}

	for i := 0; i < tiems; i++ {
		selection.moveUp()
	}
}

func TestKeyHelp(t *testing.T) {
	s := NewSelection([]string{})
	view := s.Help.View(s.Keymap)
	fmt.Println(view)
}
