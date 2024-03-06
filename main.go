package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/glamour"
)

const HELP = "help"
const CREATE = "create"
const LIST = "list"
const INSTALL = "install"
const REMOVE = "remove"

const CAT = "cat"
const FILE = "main.go"
const COPY = "cp"
const MOVE = "mv"
const MAKE_DIR = "mkdir"
const RM_DIR = "rmdir"
const RM_FILE = "rm"
const TOUCH = "touch"

const ICON_PATH = "~/.icons"
const THEME_PATH = "~/.themes"
const STATUS_PATH = "~/."
const VSCODE_PATH = "~/.vscode"
const DESKTOP_PATH = "~/."
const TERMINAL_PATH = "~/."
const ULAUNCHER_PATH = "~/.config/ulauncher/user-themes/"
const TEST_PATH = "./test/"
const TEST_FILE = "./test/test.txt"

var Options struct {
	OptionOne   string
	OptionTwo   string
	OptionThree string
	OptionFour  string
}

func main() {
	executeArguments()
}

func executeArguments() {
	arguments := os.Args
	if len(arguments) == 1 {
		help()
	}

	switch arguments[1] {
	case HELP:
		help()
	case LIST:
		if len(arguments) == 2 {
			list("all")
		}
		if len(arguments) == 3 {
			list(arguments[2])
		}
	case CREATE:
		create_project()
	case INSTALL:
		install()
	case REMOVE:
		remove()
	default:
		help()
	}
	fmt.Println("Program End.")
}

func help() {
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

func list(category string) {
	switch category {
	case "all":
		list_all()
	case "icons":
		list_category(category)
	case "themes":
		list_category(category)
	default:
		help()
	}

}

func create_project() {
	in_development()
}

func install() {
	in_development()
}

func remove() {
	in_development()
}

func build_path(directory string) string {
	shell_variables := os.Environ()
	for _, variable := range shell_variables {
		if strings.Contains(variable, "LOGNAME=") {
			current_user := strings.Split(variable, "LOGNAME=")
			return "/home/" + current_user[1] + "/." + directory
		}
	}
	return "build path error"
}

func list_category(category string) {
	path := build_path(category)
	file, err := os.ReadFile("markdown/" + strings.Split(path, ".")[1] + ".md")
	if err != nil {
		fmt.Println(err.Error())
	}
	render_icons, err := glamour.Render(string(file), "dark")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(render_icons)

	cmd := exec.Command("tree", path, "-L", "1", "-C")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout))
}

func list_all() {
	icon_path := build_path("icons")
	themes_path := build_path("themes")
	paths := []string{icon_path, themes_path}

	for _, path := range paths {
		file, err := os.ReadFile("markdown/" + strings.Split(path, ".")[1] + ".md")
		if err != nil {
			fmt.Println(err.Error())
		}
		render_icons, err := glamour.Render(string(file), "dark")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print(render_icons)

		cmd := exec.Command("tree", path, "-L", "1", "-C")
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(stdout))
	}
}

func in_development() {
	file_contents, err := os.ReadFile("markdown/contribute.md")
	if err != nil {
		fmt.Println(err.Error())
	}
	out, err := glamour.Render(string(file_contents), "dark")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(out)
}
