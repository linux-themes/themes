package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var listIconsCmd = &cobra.Command{
	Use:   "icons",
	Short: "List installed icons",
	Long:  `List installed icons`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		home_path, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		icon_path := home_path + "/.icons"

		if _, err := os.Stat(icon_path); os.IsNotExist(err) {
			if err := os.Mkdir(icon_path, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}

		if _, err := os.Stat(icon_path); err == nil {
			entries, err := os.ReadDir(icon_path)
			if err != nil {
				log.Fatal(err)
			}

			print(GREEN + "\nIcons" + RESET)
			for index, value := range entries {
				print(YELLOW, "\n\t", index+1, RESET)
				print("\t", value.Name())
			}

			if len(entries) == 0 {
				fmt.Println("\tNo icons installed.")
			}
		}
		println()
	},
}

var listThemesCmd = &cobra.Command{
	Use:   "themes",
	Short: "List installed themes",
	Long:  `List installed themes`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		home_path, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		themes_path := home_path + "/.themes"

		if _, err := os.Stat(themes_path); os.IsNotExist(err) {
			if err := os.Mkdir(themes_path, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}

		if _, err := os.Stat(themes_path); err == nil {
			entries, err := os.ReadDir(themes_path)
			if err != nil {
				log.Fatal(err)
			}

			print(GREEN + "\nThemes" + RESET)
			for index, value := range entries {
				print(YELLOW, "\n\t", index+1, RESET)
				print("\t", value.Name())
			}

			if len(entries) == 0 {
				fmt.Println("\tNo themes installed.")
			}
		}
		println()
	},
}

var listStoreCmd = &cobra.Command{
	Use:   "store",
	Short: "List store themes",
	Long:  `List store themes`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		themes := Yaml_get_file(".themes")
		icons := Yaml_get_file(".icons")

		index := 1
		print(CYAN + "Store Themes" + RESET)
		for _, value := range themes.List {
			print(YELLOW, "\n\t", index, RESET)
			print("\t", value.Name)
			index++
		}

		index = 1
		print(CYAN + "\nStore Icons" + RESET)
		for _, value := range icons.List {
			print(YELLOW, "\n\t", index, RESET)
			print("\t", value.Name)
			index++
		}
		println()
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all installed themes and icons",
	Long:  `List all installed themes and icons`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		listStoreCmd.Run(cmd, args)
		listIconsCmd.Run(cmd, args)
		listThemesCmd.Run(cmd, args)
		println()
	},
}

func init() {
	listCmd.AddCommand(listStoreCmd)
	listCmd.AddCommand(listIconsCmd)
	listCmd.AddCommand(listThemesCmd)

	listCmd.DisableFlagParsing = true
	listCmd.InitDefaultHelpFlag()
	listCmd.Flags().MarkHidden("help")

	rootCmd.AddCommand(listCmd)
}
