package commands

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
)

func HelpCommand() {
	file_contents, err := os.ReadFile("markdown/help.md")
	if err != nil {
		fmt.Println(err.Error())
	}
	out, err := glamour.Render(string(file_contents), "dark")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(out)
}
