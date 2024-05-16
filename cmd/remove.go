package cmd

import (
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var removeAllCmd = &cobra.Command{ // add confirmation
	Use:       "remove",
	Short:     "Remove all packages from icons and themes",
	Long:      `Remove all packages from icons and themes`,
	ValidArgs: []string{"all"},
	Args:      cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// home_path, err := os.UserHomeDir()
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// if err := os.RemoveAll(home_path + "/.icons/"); err != nil {
		// 	log.Fatal(err)
		// }
		// if err := os.RemoveAll(home_path + "/.themes/"); err != nil {
		// 	log.Fatal(err)
		// }
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove packages",
	Long:  `Remove packages`,
	Run: func(cmd *cobra.Command, args []string) {

		icons := []huh.Option[string]{}
		for index := 0; index < 10; index++ {
			option := huh.Option[string]{}
			option.Key = "package " + strconv.Itoa(index)
			option.Selected(false)
			icons = append(icons, option)
		}

		themes := []huh.Option[string]{}
		for index := 0; index < 10; index++ {
			option := huh.Option[string]{}
			option.Key = "package " + strconv.Itoa(index)
			option.Selected(false)
			themes = append(themes, option)
		}

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					Title("\nIcons").
					Options(icons...),

				huh.NewMultiSelect[string]().
					Title("\nThemes").
					Options(themes...),

				huh.NewSelect[int]().
					Title("\nConfirm").
					Options(
						huh.NewOption("Cancel", 0),
						huh.NewOption("Remove", 1).Selected(true),
					),
			),
		).WithTheme(huh.ThemeCharm())

		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}

		// cancel_option := true
		// if !cancel_option {
		// 	for _, option := range icons {
		// 		if option.Value == "true" {
		// 			if err := os.RemoveAll(home_path + "/.icons/" + option.Value); err != nil {
		// 				log.Fatal(err)
		// 			}
		// 			break
		// 		}
		// 	}

		// 	for _, option := range themes {
		// 		if option.Value == "true" {
		// 			if err := os.RemoveAll(home_path + "/.themes/" + option.Value); err != nil {
		// 				log.Fatal(err)
		// 			}
		// 			break
		// 		}
		// 	}
		// } else {
		// 	fmt.Println("Command Canceled")
		// }

	},
}

func init() {
	rootCmd.AddCommand(removeAllCmd)
	rootCmd.AddCommand(removeCmd)
}
