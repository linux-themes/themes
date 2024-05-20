package cmd

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// make copy public

func ThemeBase() *huh.Theme {
	var t huh.Theme

	t.FieldSeparator = lipgloss.NewStyle()
	button := lipgloss.NewStyle()

	f := &t.Focused
	f.Base = lipgloss.NewStyle()
	f.Card = lipgloss.NewStyle()
	f.ErrorIndicator = lipgloss.NewStyle().SetString(" *")
	f.ErrorMessage = lipgloss.NewStyle().SetString(" *")
	f.SelectSelector = lipgloss.NewStyle().SetString("> ")
	f.MultiSelectSelector = lipgloss.NewStyle().SetString("> ")
	f.SelectedPrefix = lipgloss.NewStyle().SetString("[•] ")
	f.UnselectedPrefix = lipgloss.NewStyle().SetString("[ ] ")
	f.FocusedButton = button.Copy().Foreground(lipgloss.Color("0")).Background(lipgloss.Color("7"))
	f.BlurredButton = button.Copy().Foreground(lipgloss.Color("7")).Background(lipgloss.Color("0"))
	f.TextInput.Placeholder = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

	t.Help = help.New().Styles
	t.Blurred = f.Copy()
	t.Blurred.Base = t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.MultiSelectSelector = lipgloss.NewStyle().SetString("  ")

	return &t
}

func ThemeCustom() *huh.Theme {
	t := ThemeBase().Copy()

	var (
		normalFg = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}
		indigo   = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
		cream    = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
		fuchsia  = lipgloss.Color("#F780E2")
		green    = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
		red      = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}
	)

	f := &t.Focused
	f.Base = lipgloss.NewStyle().
		PaddingLeft(1)
	f.Title.Foreground(green).Bold(true)
	f.NoteTitle.Foreground(indigo).Bold(true).MarginBottom(1)
	f.Description.Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"})
	f.ErrorIndicator.Foreground(red)
	f.ErrorMessage.Foreground(red)
	f.SelectSelector.Foreground(red)
	f.Option.Foreground(normalFg)
	f.MultiSelectSelector.Foreground(red)
	f.SelectedOption.Foreground(lipgloss.ANSIColor(11))
	f.SelectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#02CF92", Dark: "#02A877"}).SetString("✓ ")
	f.UnselectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"}).SetString("• ")
	f.UnselectedOption.Foreground(normalFg)
	f.FocusedButton.Foreground(cream).Background(fuchsia)
	f.Next = f.FocusedButton.Copy()
	f.BlurredButton.Foreground(normalFg).Background(lipgloss.AdaptiveColor{Light: "252", Dark: "237"})
	f.TextInput.Cursor.Foreground(green)
	f.TextInput.Placeholder.Foreground(lipgloss.AdaptiveColor{Light: "248", Dark: "238"})
	f.TextInput.Prompt.Foreground(fuchsia)

	t.Blurred = f.Copy()

	return &t
}
