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
	switch str {
	case "GNOME":
		return true
	case "KDE":
		return true
	case "Cinnamon":
		return true
	default:
		return false
	}
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set selected theme",
	Long:  `Set selected theme`,
	Run: func(cmd *cobra.Command, args []string) {

		desktop_env := os.Getenv("XDG_CURRENT_DESKTOP")
		// fmt.Println(desktop_env)

		if !isValidEnvirnoment(desktop_env) {
			log.Fatal(RED + "Desktop envirnoment not supported." + RESET + YELLOW + " Contribute to https://github.com/linux-themes/themes" + RESET)
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
						huh.NewOption("Set", 0),
						huh.NewOption("Cancel", 1).Selected(true),
					).
					Value(&form_results_cancel),
			),
		).WithTheme(ThemeCustom())

		err = form.Run()
		if err != nil {
			log.Fatal(err)
		}

		if form_results_cancel == 1 {
			fmt.Println("Command Canceled")
		} else {
			// gsettings set org.gnome.desktop.interface icon-theme 'mint'
			// dconf write /org/gnome/shell/extensions/user-theme/name "'Marble-purple-dark'"
			runCommand("gsettings", "set", "org.gnome.desktop.interface", "icon-theme", form_results_icons)
			runCommand("dconf", "write", "/org/gnome/shell/extensions/user-theme/name", "'"+form_results_themes+"'")

			fmt.Println(YELLOW + "Icons set: " + RESET + GREEN + string(form_results_icons) + RESET)
			fmt.Println(YELLOW + "Themes set: " + RESET + GREEN + string(form_results_themes) + RESET)
		}
	},
}

func runCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
}

func init() {
	rootCmd.AddCommand(setCmd)
}
