package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// TODO
// -gnome
// -XDG_CURRENT_DESKTOP=GNOME
// -kde
// -cinnamon
// -mate

func is_valid_envirnoment(str string) bool {
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

func execute_command(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set selected theme",
	Long:  `Set selected theme`,
	Run: func(cmd *cobra.Command, args []string) {

		desktop_env := os.Getenv("XDG_CURRENT_DESKTOP")

		if !is_valid_envirnoment(desktop_env) {
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

		field_icons := huh.NewSelect[string]().
			Title("\nIcons").
			Options(icons...).
			Value(&form_results_icons)

		field_themes :=
			huh.NewSelect[string]().
				Title("\nThemes").
				Options(themes...).
				Value(&form_results_themes)

		field_select :=
			huh.NewSelect[int]().
				Title("\nConfirm").
				Options(
					huh.NewOption("Set", 0),
					huh.NewOption("Cancel", 1).Selected(true),
				).
				Value(&form_results_cancel)

		fields := []huh.Field{field_icons, field_themes, field_select}

		if len(icons) < 1 {
			fields = []huh.Field{field_themes, field_select}
		}

		if len(themes) < 1 {
			fields = []huh.Field{field_icons, field_select}
		}

		form := huh.NewForm(
			huh.NewGroup(fields...),
		).WithTheme(ThemeCustom())

		err = form.Run()
		if err != nil {
			log.Fatal(err)
		}

		if form_results_cancel == 1 {
			println("Command Canceled")
		} else {
			execute_command("gnome-extensions", "enable", "user-theme@gnome-shell-extensions.gcampax.github.com") // needs a variable check
			execute_command("gsettings", "set", "org.gnome.desktop.interface", "icon-theme", form_results_icons)
			execute_command("dconf", "write", "/org/gnome/shell/extensions/user-theme/name", "'"+form_results_themes+"'")

			println(YELLOW + "Icons set: " + RESET + GREEN + string(form_results_icons) + RESET)
			println(YELLOW + "Themes set: " + RESET + GREEN + string(form_results_themes) + RESET)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
