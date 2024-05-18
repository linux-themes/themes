package cmd

import (
	"fmt"
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
		fmt.Println(RED + "Icons removed." + RESET)
		// 	log.Fatal(err)
		// }
		// if err := os.RemoveAll(home_path + "/.themes/"); err != nil {
		// fmt.Println("Themes removed.")
		fmt.Println(RED + "Icons removed." + RESET)
		// 	log.Fatal(err)
		// }
	},
}

var removeSelectedCmd = &cobra.Command{
	Use:       "remove",
	Short:     "Remove packages",
	Long:      `Remove packages`,
	ValidArgs: []string{"icons, themes"},
	Args:      cobra.ExactArgs(1),
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

		for _, option := range icons {
			fmt.Println(option.String())
		}

		for _, option := range themes {
			fmt.Println(option.String())
		}

	},
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove theme",
	Long:  `remove theme`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(args)

		if len(args) == 0 {
			removeSelectedCmd.Run(cmd, args)
		}

		if args[0] == "all" {
			removeAllCmd.Run(cmd, args)
		}

		if args[0] == "icons" {
			removeSelectedCmd.Run(cmd, args)
		}

		if args[0] == "themes" {
			removeSelectedCmd.Run(cmd, args)
		}

	},
}

func init() {
	removeCmd.AddCommand(removeAllCmd)
	removeCmd.AddCommand(removeSelectedCmd)

	removeCmd.DisableFlagParsing = true
	removeCmd.InitDefaultHelpFlag()
	removeCmd.Flags().MarkHidden("help")

	rootCmd.AddCommand(removeCmd)
}
