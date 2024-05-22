package cmd

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func ThemeCustom() *huh.Theme {
	t := copyBaseTheme(huh.ThemeBase16())

	f := &t.Focused
	f.Base = lipgloss.NewStyle().PaddingLeft(1)
	f.Title.Foreground(green).Bold(true)
	f.NoteTitle.Foreground(indigo).Bold(true).MarginBottom(1)
	f.Description.Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"})
	f.ErrorIndicator.Foreground(red)
	f.ErrorMessage.Foreground(red)
	f.SelectSelector.Foreground(red)
	f.Option.Foreground(normalFg)
	f.MultiSelectSelector.Foreground(red)
	f.SelectedOption.Foreground(YELLOW_L)
	f.SelectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#02CF92", Dark: "#02A877"}).SetString("✓ ")
	f.UnselectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"}).SetString("• ")
	f.UnselectedOption.Foreground(normalFg)
	f.FocusedButton.Foreground(cream).Background(FUCHSIA)
	f.Next = f.FocusedButton.Copy()
	f.BlurredButton.Foreground(normalFg).Background(lipgloss.AdaptiveColor{Light: "252", Dark: "237"})
	f.TextInput.Cursor.Foreground(green)
	f.TextInput.Placeholder.Foreground(lipgloss.AdaptiveColor{Light: "248", Dark: "238"})
	f.TextInput.Prompt.Foreground(FUCHSIA)

	b := &t.Blurred
	b.Description.Italic(true)
	b.TextInput.Placeholder.Italic(true)
	b.SelectedOption.Foreground(lipgloss.NewStyle().GetForeground())
	b.SelectSelector.Foreground(lipgloss.NewStyle().GetForeground())

	return &t
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
