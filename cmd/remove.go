package cmd

import (
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

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
			println(RED + "Icons removed." + RESET)
			if err := os.RemoveAll(home_path + "/.themes/"); err != nil {
				log.Fatal(err)
			}
			println(RED + "Themes removed." + RESET)
		} else {
			if args[0] == "icons" {
				if err := os.RemoveAll(home_path + "/.icons/"); err != nil {
					log.Fatal(err)
				}
				println(RED + "Icons removed." + RESET)
			}

			if args[0] == "themes" {
				if err := os.RemoveAll(home_path + "/.themes/"); err != nil {
					log.Fatal(err)
				}
				println(RED + "Themes removed." + RESET)
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

			println(YELLOW + "Icons removed: " + RESET)
			for _, result := range form_results_strings {
				if err := os.RemoveAll(home_path + "/.icons/" + result); err != nil {
					println(YELLOW + "Package does not exist: " + result + RESET)
				} else {
					println(RED + "\t" + result + RESET)
				}
			}
		} else {
			println(YELLOW + "Command canceled." + RESET)
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

			println(YELLOW + "Themes removed: " + RESET)
			for _, result := range form_results_strings {
				if err := os.RemoveAll(home_path + "/.themes/" + result); err != nil {
					println(YELLOW + "Package does not exist: " + result + RESET)
				} else {
					println(RED + "\t" + result + RESET)
				}
			}
		} else {
			println(YELLOW + "Command canceled." + RESET)
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
