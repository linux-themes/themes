package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const (
	BLUE      = "\033[1;34m"
	BLUE_THIN = "\033[0;36m"
	CYAN      = "\033[1;36m"
	YELLOW    = "\033[1;33m"
	RED       = "\033[1;31m"
	GREEN     = "\033[1;32m"
	RESET     = "\033[0m"
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

			println(GREEN + "\nIcons" + RESET)
			for _, e := range entries {
				fmt.Println(CYAN + "\t" + e.Name() + RESET)
			}

			if len(entries) == 0 {
				fmt.Println("\tNo icons installed. To add an icon package:")
				fmt.Println()
				fmt.Println("\tthemes install icons gicons")
			}

			fmt.Println()
		}
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

			println(GREEN + "\nThemes" + RESET)
			for _, e := range entries {
				fmt.Println(CYAN + "\t" + e.Name() + RESET)
			}

			if len(entries) == 0 {
				fmt.Println("\tNo themes installed. To add an theme package:")
				fmt.Println()
				fmt.Println("\tthemes install themes gtheme")
			}

			fmt.Println()
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all installed themes and icons",
	Long:  `List all installed themes and icons`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		listIconsCmd.Run(cmd, args)
		listThemesCmd.Run(cmd, args)
		fmt.Println()
	},
}

func init() {
	listCmd.AddCommand(listIconsCmd)
	listCmd.AddCommand(listThemesCmd)

	listCmd.DisableFlagParsing = true
	listCmd.InitDefaultHelpFlag()
	listCmd.Flags().MarkHidden("help")

	rootCmd.AddCommand(listCmd)
}