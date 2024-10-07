package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

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

			println(GREEN + "\nIcons" + RESET)
			for _, e := range entries {
				fmt.Println(CYAN + "\t" + e.Name() + RESET)
			}

			if len(entries) == 0 {
				fmt.Println("\tNo icons installed.")
				fmt.Println()
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
				fmt.Println("\tNo themes installed.")
				fmt.Println()
			}

			fmt.Println()
		}
	},
}
var listOfficalCmd = &cobra.Command{
	Use:   "official",
	Short: "List offical themes",
	Long:  `List offical themes`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(GREEN + "Official Icons:" + RESET)

		keys := make([]int, 0)
		for k := range offical_icons {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, k := range keys {
			v := offical_icons[k]
			index := strings.LastIndex(v, "/")
			v = v[index+1:]
			index = strings.Index(v, ".")
			v = v[:index]
			fmt.Printf(CYAN+"\t%d\t%s\n"+RESET, k, v)
		}
		fmt.Println()

		keys = make([]int, 0)
		for k := range offical_themes {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		fmt.Println(GREEN + "Official Themes:" + RESET)
		for _, k := range keys {
			v := offical_icons[k]
			index := strings.LastIndex(v, "/")
			v = v[index+1:]
			index = strings.Index(v, ".")
			v = v[:index]
			fmt.Printf(CYAN+"\t%d\t%s\n"+RESET, k, v)
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all installed themes and icons",
	Long:  `List all installed themes and icons`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// listOfficalCmd.Run(cmd, args)
		listIconsCmd.Run(cmd, args)
		listThemesCmd.Run(cmd, args)
		fmt.Println()
	},
}

func init() {
	listCmd.AddCommand(listOfficalCmd)
	listCmd.AddCommand(listIconsCmd)
	listCmd.AddCommand(listThemesCmd)

	listCmd.DisableFlagParsing = true
	listCmd.InitDefaultHelpFlag()
	listCmd.Flags().MarkHidden("help")

	rootCmd.AddCommand(listCmd)
}
