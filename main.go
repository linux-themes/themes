package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/glamour"
)

const HELP = "help"
const INIT = "init"
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
	arguments := os.Args
	executeArguments(arguments)
}

func executeArguments(arguments []string) {
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
	case INIT:
		init_project()
	case INSTALL:
		install()
	case REMOVE:
		remove()
	default:
		fmt.Println("Execute Arguments Error")
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

func init_project() {
	file_contents, err := os.ReadFile("mardown/init.md")
	if err != nil {
		fmt.Println(err.Error())
	}
	out, err := glamour.Render(string(file_contents), "dark")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(out)
}

func install() {
	cmd_one := exec.Command(MAKE_DIR, TEST_PATH)
	stdout_one, err := cmd_one.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout_one))

	cmd_two := exec.Command(TOUCH, TEST_FILE)
	stdout_two, err := cmd_two.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout_two))
}

func remove() {
	cmd := exec.Command(RM_FILE, TEST_FILE)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout))
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
