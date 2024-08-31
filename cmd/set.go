package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// TODO set themes

// gnome
// XDG_CURRENT_DESKTOP=GNOME
// kde
// cinnamon
// mate

func isValidEnvirnoment(str string) bool {
	return str == "GNOME"
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set selected theme",
	Long:  `Set selected theme`,
	Run: func(cmd *cobra.Command, args []string) {

		desktop_env := os.Getenv("XDG_CURRENT_DESKTOP")
		// fmt.Println(desktop_env)

		if !isValidEnvirnoment(desktop_env) {
			log.Fatal("desktop envirnoment not supported")
		}

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
						huh.NewOption("Set", 1),
					).
					Value(&form_results_cancel),
			),
		).WithTheme(ThemeCustom())

		err = form.Run()
		if err != nil {
			log.Fatal(err)
		}

		if form_results_cancel == 0 {
			fmt.Println("Command Canceled")
		} else {
			// gsettings set org.gnome.desktop.interface icon-theme 'MyIconTheme'
			arg0 := "gsettings"
			arg1 := "set"
			arg2 := "org.gnome.desktop.interface"
			arg3 := "icon-theme"

			cmd := exec.Command(arg0, arg1, arg2, arg3, form_results_icons)
			stdout, err := cmd.Output()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			// Print the output
			fmt.Println(string(stdout))
			fmt.Println("Icons set: " + string(form_results_icons))
		}

	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
