package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/glamour"
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
const ICONS = "icons"
const THEMES = "themes"
const ALL = "all"

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
var THEME_PATH = USER_PATH + "/.themes"
var TERMINAL_PATH = USER_PATH + "/.gnome/terminal/themes"
var ULAUNCHER_PATH = USER_PATH + "/.config/ulauncher/user-themes"

const TEST_PACKAGE = "https://github.com/sudo-adduser-jordan/mint-y-winx/raw/main/mint-y-winx.tar.xz"

func main() {
	executeArguments()
}

func executeArguments() {
	arguments := os.Args
	if len(arguments) == 1 {
		help_command()
		return
	}

	switch arguments[1] {
	case HELP:
		help_command()
	case LIST:
		if len(arguments) == 2 {
			list_command("all")
		}
		if len(arguments) == 3 {
			list_command(arguments[2])
		}
	case CREATE:
		create_command()
	case BUILD:
		build_command()
	case INSTALL:
		install_command(arguments)
	case SET:
		set_command()
	case REMOVE:
		remove_command()
	default:
		help_command()
	}

	fmt.Println("Program End.")
}

func help_command() {
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

func list_command(category string) {
	switch category {
	case ALL:
		list_all()
	case ICONS:
		list_category(category)
	case THEMES:
		list_category(category)
	default:
		help_command()
	}
}

func list_category(category string) {
	file, err := os.ReadFile("markdown/." + category + ".md")
	if err != nil {
		fmt.Println(err.Error())
	}
	render_icons, err := glamour.Render(string(file), "dark")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(render_icons)

	cmd := exec.Command("tree", USER_PATH+"."+category, "-L", "1", "-C")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout))
}

func list_all() {
	paths := []string{ICON_PATH, THEME_PATH}
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

func create_command() {
	InDevelopment()
}

func build_command() {
	InDevelopment()
}

func install_command(arguments []string) {
	InDevelopment()

	if len(arguments) == 3 {
		if arguments[2] == "icons" {
			icon_packs := []string{arguments[3]}
			install(icon_packs, ".icons")
		}
		if arguments[2] == "themes" {
			themes_packs := []string{arguments[3]}
			install(themes_packs, ".themes")
		}
		return
	}

	if len(arguments) > 3 {
		urls := arguments[2:]
		packages := []string{}
		for _, url := range urls {
			if !ValidUrl(url) {
				help_command()
			}
			packages = append(packages, url)
		}
		install(packages, "."+arguments[2])
		return
	}

	help_command()
}

func install(links []string, directory string) {
	err := os.Mkdir(directory, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, link := range links {
		file_name := StripFileNameGit(link)
		directory_path := BuildPathHomeUserDirectory(directory)
		file_path := directory_path + "/" + file_name

		fmt.Println("Installing: " + file_path)
		if err := DownloadFile(file_path, link); err != nil {
			fmt.Println(err.Error())
		}
		if err := Extract_Tar(file_path, directory_path); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func set_command() {
	InDevelopment()
}

func remove_command() {
	// all
	// single
	// selected
	InDevelopment()
}
