package commands

import (
	"fmt"
	"main/markdown"

	"github.com/charmbracelet/glamour"
)

func HelpCommand() {
	out, err := glamour.Render(string(markdown.HELP_MARKDOWN), "dark")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(out)
}
