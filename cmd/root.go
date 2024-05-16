package cmd

import (
	"os"

	"github.com/spf13/cobra"

	cc "github.com/ivanpirog/coloredcobra"
)

var rootCmd = &cobra.Command{
	Use:   "themes",
	Short: "Install and set themes for linux desktop environments",
	Long:  `Install and set themes for linux desktop environments`,
}

func Execute() {

	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		Headings:        cc.HiCyan + cc.Bold,
		Commands:        cc.HiYellow + cc.Bold,
		Example:         cc.Italic,
		ExecName:        cc.Bold,
		Flags:           cc.Bold,
		NoExtraNewlines: true,
	})

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	rootCmd.InitDefaultHelpFlag()
	rootCmd.Flags().MarkHidden("help")
}
