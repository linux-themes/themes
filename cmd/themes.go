package cmd

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var formTheme *huh.Theme

func ThemeCustom() {
	t := copyBaseTheme(huh.ThemeBase())

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

	// t.Blurred = f.Copy() // throws

	formTheme = &t
}

func copyBaseTheme(original *huh.Theme) huh.Theme {
	return huh.Theme{
		Form:           original.Form.Copy(),
		Group:          original.Group.Copy(),
		FieldSeparator: original.FieldSeparator.Copy(),
		Blurred: huh.FieldStyles{
			Base:                original.Blurred.Base.Copy(),
			Title:               original.Blurred.Title.Copy(),
			Description:         original.Blurred.Description.Copy(),
			ErrorIndicator:      original.Blurred.ErrorIndicator.Copy(),
			ErrorMessage:        original.Blurred.ErrorMessage.Copy(),
			SelectSelector:      original.Blurred.SelectSelector.Copy(),
			Option:              original.Blurred.Option.Copy(),
			MultiSelectSelector: original.Blurred.MultiSelectSelector.Copy(),
			SelectedOption:      original.Blurred.SelectedOption.Copy(),
			SelectedPrefix:      original.Blurred.SelectedPrefix.Copy(),
			UnselectedOption:    original.Blurred.UnselectedOption.Copy(),
			UnselectedPrefix:    original.Blurred.UnselectedPrefix.Copy(),
			FocusedButton:       original.Blurred.FocusedButton.Copy(),
			BlurredButton:       original.Blurred.BlurredButton.Copy(),
			TextInput: huh.TextInputStyles{
				Cursor:      original.Blurred.TextInput.Cursor.Copy(),
				Placeholder: original.Blurred.TextInput.Placeholder.Copy(),
				Prompt:      original.Blurred.TextInput.Prompt.Copy(),
				Text:        original.Blurred.TextInput.Text.Copy(),
			},
			Card: original.Blurred.Card.Copy(),
			Next: original.Blurred.Next.Copy(),
		},
		Focused: huh.FieldStyles{
			Base:                original.Focused.Base.Copy(),
			Title:               original.Focused.Title.Copy(),
			Description:         original.Focused.Description.Copy(),
			ErrorIndicator:      original.Focused.ErrorIndicator.Copy(),
			ErrorMessage:        original.Focused.ErrorMessage.Copy(),
			SelectSelector:      original.Focused.SelectSelector.Copy(),
			Option:              original.Focused.Option.Copy(),
			MultiSelectSelector: original.Focused.MultiSelectSelector.Copy(),
			SelectedOption:      original.Focused.SelectedOption.Copy(),
			SelectedPrefix:      original.Focused.SelectedPrefix.Copy(),
			UnselectedOption:    original.Focused.UnselectedOption.Copy(),
			UnselectedPrefix:    original.Focused.UnselectedPrefix.Copy(),
			FocusedButton:       original.Focused.FocusedButton.Copy(),
			BlurredButton:       original.Focused.BlurredButton.Copy(),
			TextInput: huh.TextInputStyles{
				Cursor:      original.Focused.TextInput.Cursor.Copy(),
				Placeholder: original.Focused.TextInput.Placeholder.Copy(),
				Prompt:      original.Focused.TextInput.Prompt.Copy(),
				Text:        original.Focused.TextInput.Text.Copy(),
			},
			Card: original.Focused.Card.Copy(),
			Next: original.Focused.Next.Copy(),
		},
		Help: help.Styles{
			Ellipsis:       original.Help.Ellipsis.Copy(),
			ShortKey:       original.Help.ShortKey.Copy(),
			ShortDesc:      original.Help.ShortDesc.Copy(),
			ShortSeparator: original.Help.ShortSeparator.Copy(),
			FullKey:        original.Help.FullKey.Copy(),
			FullDesc:       original.Help.FullDesc.Copy(),
			FullSeparator:  original.Help.FullSeparator.Copy(),
		},
	}
}

// import (
// 	_ "unsafe" // used for the hacky interalCopy linking.

// 	"github.com/charmbracelet/bubbles/help"
// 	"github.com/charmbracelet/huh"
// 	"github.com/charmbracelet/lipgloss"
// )

// // make copy public

// // maybe a bit hacky but the copy method isn't exported for some reason because its private duh
// //
// //go:linkname Copy github.com/charmbracelet/huh.Theme.copy
// func Copy(huh.Theme) huh.Theme

// func ThemeBase() *huh.Theme {
// 	var t huh.Theme

// 	t.FieldSeparator = lipgloss.NewStyle()
// 	button := lipgloss.NewStyle()

// 	f := &t.Focused
// 	f.Base = lipgloss.NewStyle()
// 	f.Card = lipgloss.NewStyle()
// 	f.ErrorIndicator = lipgloss.NewStyle().SetString(" *")
// 	f.ErrorMessage = lipgloss.NewStyle().SetString(" *")
// 	f.SelectSelector = lipgloss.NewStyle().SetString("> ")
// 	f.MultiSelectSelector = lipgloss.NewStyle().SetString("> ")
// 	f.SelectedPrefix = lipgloss.NewStyle().SetString("[•] ")
// 	f.UnselectedPrefix = lipgloss.NewStyle().SetString("[ ] ")
// 	f.FocusedButton = button.Copy().Foreground(lipgloss.Color("0")).Background(lipgloss.Color("7"))
// 	f.BlurredButton = button.Copy().Foreground(lipgloss.Color("7")).Background(lipgloss.Color("0"))
// 	f.TextInput.Placeholder = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

// 	t.Help = help.New().Styles
// 	t.Blurred = f.Copy()
// 	t.Blurred.Base = t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())
// 	t.Blurred.MultiSelectSelector = lipgloss.NewStyle().SetString("  ")

// 	return &t
// }

// func ThemeCustom() *huh.Theme {
// 	t := ThemeBase().Copy()

// 	var (
// 		normalFg = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}
// 		indigo   = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
// 		cream    = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
// 		fuchsia  = lipgloss.Color("#F780E2")
// 		green    = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
// 		red      = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}
// 	)

// 	f := &t.Focused
// 	f.Base = lipgloss.NewStyle().
// 		PaddingLeft(1)
// 	f.Title.Foreground(green).Bold(true)
// 	f.NoteTitle.Foreground(indigo).Bold(true).MarginBottom(1)
// 	f.Description.Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"})
// 	f.ErrorIndicator.Foreground(red)
// 	f.ErrorMessage.Foreground(red)
// 	f.SelectSelector.Foreground(red)
// 	f.Option.Foreground(normalFg)
// 	f.MultiSelectSelector.Foreground(red)
// 	f.SelectedOption.Foreground(lipgloss.ANSIColor(11))
// 	f.SelectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#02CF92", Dark: "#02A877"}).SetString("✓ ")
// 	f.UnselectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"}).SetString("• ")
// 	f.UnselectedOption.Foreground(normalFg)
// 	f.FocusedButton.Foreground(cream).Background(fuchsia)
// 	f.Next = f.FocusedButton.Copy()
// 	f.BlurredButton.Foreground(normalFg).Background(lipgloss.AdaptiveColor{Light: "252", Dark: "237"})
// 	f.TextInput.Cursor.Foreground(green)
// 	f.TextInput.Placeholder.Foreground(lipgloss.AdaptiveColor{Light: "248", Dark: "238"})
// 	f.TextInput.Prompt.Foreground(fuchsia)

// 	t.Blurred = f.Copy()

// 	return &t
// }
