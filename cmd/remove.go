package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

func ThemeCustom() *huh.Theme {
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

	t.Blurred = f.Copy()

	return &t
}

// What I've implemented is a direct duplicate of huh theme.copy()
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

var removeAllCmd = &cobra.Command{
	Use:       "all",
	Short:     "Remove all packages from icons and themes",
	Long:      `Remove all packages from icons and themes`,
	ValidArgs: []string{"icons, themes"},
	Args:      cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		home_path, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		if len(args) == 0 {
			if err := os.RemoveAll(home_path + "/.icons/"); err != nil {
				log.Fatal(err)
			}
			fmt.Println(RED + "Icons removed." + RESET)
			if err := os.RemoveAll(home_path + "/.themes/"); err != nil {
				log.Fatal(err)
			}
			fmt.Println(RED + "Themes removed." + RESET)
		} else {
			if args[0] == "icons" {
				if err := os.RemoveAll(home_path + "/.icons/"); err != nil {
					log.Fatal(err)
				}
				fmt.Println(RED + "Icons removed." + RESET)
			}

			if args[0] == "themes" {
				if err := os.RemoveAll(home_path + "/.themes/"); err != nil {
					log.Fatal(err)
				}
				fmt.Println(RED + "Themes removed." + RESET)
			}
		}
	},
}

var removeIconsCmd = &cobra.Command{
	Use:   "icons",
	Short: "Remove packages from icons",
	Long:  `Remove packages from icons`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		home_path, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		entries, err := os.ReadDir(home_path + "/.icons")
		if err != nil {
			log.Fatal(err)
		}

		icons := []huh.Option[string]{}
		for _, entry := range entries {
			option := huh.Option[string]{}
			option.Key = entry.Name()
			option.Value = entry.Name()
			option.Selected(false)
			icons = append(icons, option)
		}

		var form_results int
		var form_results_strings []string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					Title("\nIcons").
					Options(icons...).
					Value(&form_results_strings),

				huh.NewSelect[int]().
					Title("\nConfirm").
					Options(
						huh.NewOption("Cancel", 0).Selected(true),
						huh.NewOption("Remove", 1),
					).Value(&form_results),
			),
		).WithTheme(ThemeCustom())

		err = form.Run()
		if err != nil {
			log.Fatal(err)
		}

		if form_results == 1 {
			home_path, err := os.UserHomeDir()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(YELLOW + "Icons removed: " + RESET)
			for _, result := range form_results_strings {
				if err := os.RemoveAll(home_path + "/.icons/" + result); err != nil {
					fmt.Println(YELLOW + "Package does not exist: " + result + RESET)
				} else {
					fmt.Println(RED + "\t" + result + RESET)
				}
			}
		} else {
			fmt.Println(YELLOW + "Command canceled." + RESET)
		}
	},
}

var removeThemesCmd = &cobra.Command{
	Use:   "themes",
	Short: "Remove packages from themes",
	Long:  `Remove packages from themes`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		home_path, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		entries, err := os.ReadDir(home_path + "/.themes")
		if err != nil {
			log.Fatal(err)
		}

		themes := []huh.Option[string]{}
		for _, entry := range entries {
			option := huh.Option[string]{}
			option.Key = entry.Name()
			option.Value = entry.Name()
			option.Selected(false)
			themes = append(themes, option)
		}

		var form_results int
		var form_results_strings []string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					Title("\nThemes").
					Options(themes...).
					Value(&form_results_strings),

				huh.NewSelect[int]().
					Title("\nConfirm").
					Options(
						huh.NewOption("Cancel", 0).Selected(true),
						huh.NewOption("Remove", 1),
					).
					Value(&form_results),
			),
		).WithTheme(ThemeCustom())

		err = form.Run()
		if err != nil {
			log.Fatal(err)
		}

		if form_results == 1 {
			home_path, err := os.UserHomeDir()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(YELLOW + "Themes removed: " + RESET)
			for _, result := range form_results_strings {
				if err := os.RemoveAll(home_path + "/.themes/" + result); err != nil {
					fmt.Println(YELLOW + "Package does not exist: " + result + RESET)
				} else {
					fmt.Println(RED + "\t" + result + RESET)
				}
			}
		} else {
			fmt.Println(YELLOW + "Command canceled." + RESET)
		}
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove icons and themes",
	Long:  `remove icons and themes`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if args[0] == "all" {
			removeAllCmd.Run(cmd, args)
		}

		if args[0] == "icons" {
			removeIconsCmd.Run(cmd, args)
		}

		if args[0] == "themes" {
			removeThemesCmd.Run(cmd, args)
		}

	},
}

func init() {
	removeCmd.AddCommand(removeAllCmd)
	removeCmd.AddCommand(removeIconsCmd)
	removeCmd.AddCommand(removeThemesCmd)

	removeCmd.DisableFlagParsing = true
	removeCmd.InitDefaultHelpFlag()
	removeCmd.Flags().MarkHidden("help")

	rootCmd.AddCommand(removeCmd)
}
