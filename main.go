package main

import (
	"fmt"
	"main/commands"
	"os"
	"strings"
)

// COMMAND ENUMS
const HELP = "help"
const CREATE = "create"
const BUILD = "build"
const LIST = "list"
const INSTALL = "install"
const SET = "set"
const REMOVE = "remove"

// ENUMS
const ALL = "all"
const ICONS = "icons"
const THEMES = "themes"
const CONFIG = "config"

// PROGRAMS
const TAR = "tar"
const CAT = "cat"
const COPY = "cp"
const MOVE = "mv"
const TOUCH = "touch"
const RM_DIR = "rmdir"
const RM_FILE = "rm"
const MAKE_DIR = "mkdir"

// PATHS
func SetUser() string {
	shell_variables := os.Environ()
	for _, variable := range shell_variables {
		if strings.Contains(variable, "LOGNAME=") {
			strings := strings.Split(variable, "LOGNAME=")
			return strings[1]
		}
	}
	return "User Error"
}

var USER = SetUser()
var HOME_PATH = "/home"
var USER_PATH = HOME_PATH + "/" + USER
var ICON_PATH = USER_PATH + "/.icons"
var CONFIG_PATH = USER_PATH + "/.config"
var THEME_PATH = USER_PATH + "/.themes"

const TEST_PACKAGE = "https://github.com/sudo-adduser-jordan/mint-y-winx/raw/main/mint-y-winx.tar.xz"

func main() {
	executeArguments()
}

func executeArguments() {
	arguments := os.Args
	if len(arguments) == 1 {
		commands.HelpCommand()
		return
	}

	switch arguments[1] {
	case HELP:
		commands.HelpCommand()
	case LIST:
		if len(arguments) == 2 {
			commands.ListCommand("all")
		}
		if len(arguments) == 3 {
			commands.ListCommand(arguments[2])
		}
	case INSTALL:
		commands.InstallCommand(arguments)
	case SET:
		commands.SetCommand()
	case REMOVE:
		commands.RemoveCommand()
	default:
		commands.HelpCommand()
	}

	fmt.Println("\n Program End.")
}
