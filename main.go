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
	switch len(arguments) {
	case 4:
		executeArguments(arguments)
	default:
		help()
	}
}

func executeArguments(arguments []string) {
	switch arguments[1] {
	case HELP:
		help()
	case LIST:
		list()
	case INIT:
		init_project(arguments[2])
	case INSTALL:
		install(arguments[1], arguments[2])
	case REMOVE:
		remove(arguments[1], arguments[2])
	default:
		fmt.Println("Execute Arguments Error")
	}
	fmt.Println("Program End.")
}

func help() {
	file_contents, err := os.ReadFile("help.md")
	if err != nil {
		fmt.Println(err.Error())
	}
	render, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	out, err := render.Render(string(file_contents))
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(out)
}

func list() {
	file_contents, err := os.ReadFile("list.md")
	if err != nil {
		fmt.Println(err.Error())
	}
	render, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	out, err := render.Render(string(file_contents))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(out)

	var icons_path string
	shell_variables := os.Environ()
	for _, variable := range shell_variables {
		if strings.Contains(variable, "LOGNAME=") {
			current_user := strings.Split(variable, "LOGNAME=")
			fmt.Println(current_user[1])
			icons_path = "/home/" + current_user[1] + "/.icons"
			// themes_path = "/home/" + current_user + "/.themes"
		}
	}

	cmd := exec.Command("tree", icons_path, "-L", "1", "-C")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout))
}

func init_project(packg_type string) {
	cmd := exec.Command(MAKE_DIR, TEST_PATH)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("STDOUT: " + string(stdout))
	fmt.Println("Project Created: " + packg_type)
}

func install(packg string, packg_id string) {
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

	fmt.Println(packg + " " + packg_id)
}

func remove(packg string, packg_id string) {
	cmd := exec.Command(RM_FILE, TEST_FILE)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout))
	fmt.Println(packg + " " + packg_id)
}

func install_dependenices() {}
