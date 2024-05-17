package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set selected theme",
	Long:  `Set selected theme`,
	Run: func(cmd *cobra.Command, args []string) {

		// home_path, err := os.UserHomeDir()
		// if err != nil {
		// 	log.Fatal(err)
		// }

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
				huh.NewSelect[string]().
					Title("\nIcons").
					Options(icons...),

				huh.NewSelect[string]().
					Title("\nThemes").
					Options(themes...),

				huh.NewSelect[int]().
					Title("\nConfirm").
					Options(
						huh.NewOption("Cancel", 0),
						huh.NewOption("Remove", 1),
					),
			),
		).WithTheme(huh.ThemeCharm())

		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}

		cancel_option := true
		if !cancel_option {
			for _, option := range icons {

				fmt.Println(option.String())

				if option.Value == "true" {
					fmt.Println(RED + "hit" + RESET)
				}
			}

			for _, option := range themes {

				fmt.Println(option.String())

				if option.Value == "true" {
					fmt.Println(RED + "hit" + RESET)
					break
				}
			}
		} else {
			fmt.Println("Command Canceled")
		}

	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
