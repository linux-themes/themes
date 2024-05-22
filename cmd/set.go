package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// gnome
// kde
// cinnamon
// mate

func getDesktopEnvirnoment() string {
	for _, value := range os.Environ() {
		fmt.Println(value)
	}
	return os.Environ()[0]
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set selected theme",
	Long:  `Set selected theme`,
	Run: func(cmd *cobra.Command, args []string) {

		desktop_env := getDesktopEnvirnoment()
		fmt.Println(desktop_env)

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

		entries, err = os.ReadDir(home_path + "/.themes")
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

		var form_results_icons string
		var form_results_themes string
		var form_results_cancel int
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("\nIcons").
					Options(icons...).
					Value(&form_results_icons),

				huh.NewSelect[string]().
					Title("\nThemes").
					Options(themes...).
					Value(&form_results_themes),

				huh.NewSelect[int]().
					Title("\nConfirm").
					Options(
						huh.NewOption("Cancel", 0).Selected(true),
						huh.NewOption("Remove", 1),
					).
					Value(&form_results_cancel),
			),
		).WithTheme(ThemeCustom())

		err = form.Run()
		if err != nil {
			log.Fatal(err)
		}

		if form_results_cancel == 1 {

		} else {
			fmt.Println("Command Canceled")
		}

	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
