package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
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

	fmt.Println("\n Program End.")
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
	fmt.Println("\n Installing... ")
	if len(arguments) == 4 {
		if !ValidUrl(arguments[3]) {
			syscall.Exit(1)
		}
		if arguments[2] == "icons" {
			icon_packs := []string{arguments[3]}
			install(icon_packs, ".icons")
		}
		if arguments[2] == "themes" {
			themes_packs := []string{arguments[3]}
			install(themes_packs, ".themes")
		}
		if arguments[2] == "config" {
			InDevelopment()
			return
		}
		return
	}

	if len(arguments) > 3 {
		urls := arguments[3:]
		directory := arguments[2]

		fmt.Print("Packages: ")
		fmt.Println(urls)

		packages := []string{}
		for _, url := range urls {
			if !ValidUrl(url) {
				help_command()
				fmt.Println("Program End.")
				syscall.Exit(0)
			}
			packages = append(packages, url)
		}
		install(packages, "."+directory)
		return
	}

	help_command()
}

func install(links []string, directory string) {
	fmt.Println("\n Creating Directory...")
	err := os.Mkdir(directory, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, link := range links {
		fmt.Println("\n Installing: " + directory + " ...")

		file_name := StripFileNameGit(link)
		directory_path := BuildPathHomeUserDirectory(directory)
		file_path := directory_path + "/" + file_name

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
	// gnome
	// cinnamon
	// i3
	// wayland
}

// exists returns whether the given file or directory exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

type Action int

const (
	Icons Action = iota
	Themes
	Config
	Spin
	Cancel
)

// var highlight = lipgloss.NewStyle().Foreground(lipgloss.Color("#00D7D7"))

func remove_command() {
	spinnerStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("4"))

	theme := huh.ThemeBase16()
	theme.FieldSeparator = lipgloss.NewStyle().SetString("\n")
	theme.Help.FullKey.MarginTop(1)

	var action Action
	f := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[Action]().
				Value(&action).
				Options(
					huh.NewOption("Icons", Icons),
					huh.NewOption("Themes", Themes),
					huh.NewOption("Config", Config),
					huh.NewOption("Spinner", Spin),
					huh.NewOption("Cancel", Cancel),
				).
				Title("Choose Folder"),
		),
	).WithTheme(theme)
	err := f.Run()
	if err != nil {
		log.Fatal(err)
	}
	switch action {
	case Icons:
		fmt.Println("Icons Action")
		packages := GetPackages("icons")
		fmt.Println(packages)
	case Themes:
		fmt.Println("Themes Action")
		packages := GetPackages("themes")
		fmt.Println(packages)
	case Config:
		fmt.Println("Config")
		packages := GetPackages("config")
		fmt.Println(packages)
	case Spin:
		fmt.Println("Spinner Example")
		_ = spinner.New().Title("Spinner example...").Style(spinnerStyle).Run()
	case Cancel:
		fmt.Println("Cancelling...")
		os.Exit(1)
	}
	// fmt.Printf("Selected themes are: \n%s \n%s \n ", highlight.Render("theme_one"), highlight.Render("theme_two"))

	var nextAction string
	f = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("Remove")...).
				Title("What's next?").Value(&nextAction),
		),
	).WithTheme(theme)

	err = f.Run()
	if err != nil {
		log.Fatal(err)
	}
	if nextAction == "Remove" {
		_ = spinner.New().Title("Removing Packages: list...").Style(spinnerStyle).Run()
		fmt.Println("Packages Removed.")
	}

}

func GetPackages(category string) []string {

	var path string
	if category == "icons" {
		path = ICON_PATH
	}
	if category == "themes" {
		path = THEME_PATH
	}
	if category == "config" {
		path = CONFIG_PATH

	}

	packages := []string{}
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			packages = append(packages, path+"/"+file.Name())
		}
	}
	return packages
}

// func getThemeOptions() []string {}
// func getThemeOptions() []string {}
// func getThemeOptions() []string {}
// func getThemeOptions() []string {}
// func getThemeOptions() []string {}
// func getThemeOptions() []string {}

// func remove_command(options []string) {
// 	fmt.Println("\n Removing... ")
// 	if len(arguments) == 4 {
// 		if _, err := Exists(arguments[3]); err != nil {
// 			fmt.Println(err.Error())
// 			panic(err)
// 		}
// 		if arguments[2] == "icons" {
// 			icon_packs := []string{arguments[3]}
// 			remove(icon_packs, ".icons")
// 		}
// 		if arguments[2] == "themes" {
// 			themes_packs := []string{arguments[3]}
// 			remove(themes_packs, ".themes")
// 		}
// 		if arguments[2] == "config" {
// 			InDevelopment()
// 			return
// 		}
// 		return
// 	}

// 	if len(arguments) > 3 {
// 		file_path := arguments[3:]
// 		directory := arguments[2]

// 		fmt.Print("Packages: ")
// 		fmt.Println(file_path)

// 		packages := []string{}
// 		for _, url := range file_path {
// 			if !ValidUrl(url) {
// 				help_command()
// 				fmt.Println("Program End.")
// 				syscall.Exit(0)
// 			}
// 			packages = append(packages, url)
// 		}
// 		install(packages, "."+directory)
// 		return
// 	}

// 	help_command()
// }

// func remove(packages []string, directory string) {
// 	for _, packag := range packages {
// 		is_valid, err := Exists(packag)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		if is_valid {
// 			fmt.Println("directory to be removed: " + packag)
// 			// if err := os.Remove(""); err != nil {
// 			// 	fmt.Println(err.Error())
// 			// }
// 		}
// 	}
// }
