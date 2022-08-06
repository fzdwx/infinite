package components

import tea "github.com/charmbracelet/bubbletea"

func FocusCmd() tea.Msg {
	return Focus
}

func BlurCmd() tea.Msg {
	return Blur
}

func QuitCmd() tea.Msg {
	return Quit
}
