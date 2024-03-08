package commands

import (
	"fmt"
	"main/markdown"
	"os/exec"

	"github.com/charmbracelet/glamour"
)

func ListCommand(category string) {
	switch category {
	case ALL:
		ListAll()
	case ICONS:
		ListCategory(category)
	case THEMES:
		ListCategory(category)
	case CONFIG:
		ListCategory(category)
	default:
		HelpCommand()
	}
}

func ListCategory(category string) {
	var file string
	if category == ICONS {
		file = markdown.ICONS_MARKDOWN
	}
	if category == THEMES {
		file = markdown.THEMES_MARKDOWN
	}
	if category == CONFIG {
		file = markdown.CONFIG_MARKDOWN
	}

	render_icons, err := glamour.Render(string(file), "dark")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(render_icons)

	cmd := exec.Command("tree", USER_PATH+"/."+category, "-L", "1", "-C")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout))
}

func ListAll() {
	files := []string{markdown.ICONS_MARKDOWN, markdown.THEMES_MARKDOWN, markdown.CONFIG_MARKDOWN}
	paths := []string{ICON_PATH, THEME_PATH, CONFIG_PATH}
	for index, file := range files {
		render_icons, err := glamour.Render(string(file), "dark")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print(render_icons)

		cmd := exec.Command("tree", paths[index], "-L", "1", "-C")
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(stdout))
	}
}
