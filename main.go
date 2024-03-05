package main

import (
	"fmt"
	"os"
	"os/exec"
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

const ICON_PATH = "~/.icons"
const THEME_PATH = "~/.themes"
const STATUS_PATH = "~/."
const VSCODE_PATH = "~/.vscode"
const DESKTOP_PATH = "~/."
const TERMINAL_PATH = "~/."
const ULAUNCHER_PATH = "~/.config/ulauncher/user-themes/"

const HELP_MESSAGE = `
	help
	maunal
	commands
`

func main() {
	arguments := os.Args
	switch len(arguments) {
	case 0:
		fmt.Println("No Arguments Error")
	case 4:
		executeArguments(arguments)
	default:
		help()
	}
}

func executeArguments(arguments []string) {
	fmt.Print("Execute: ")
	switch arguments[1] {
	case HELP:
		help()
	case LIST:
		list()
	case INIT:
		init_project(arguments[1])
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
	fmt.Print(HELP_MESSAGE)
}

func list() {
	cmd := exec.Command(CAT, FILE)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout))
}

func init_project(packg_type string) {
	cmd := exec.Command(CAT, FILE)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout))
	fmt.Println(packg_type)
}

func install(packg string, packg_id string) {
	cmd := exec.Command(CAT, FILE)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout))
	fmt.Println(packg + " " + packg_id)
}

func remove(packg string, packg_id string) {
	cmd := exec.Command(CAT, FILE)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout))
	fmt.Println(packg + " " + packg_id)
}
