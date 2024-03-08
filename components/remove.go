package components

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
)

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

type Option int

const (
	ICONS Option = iota + 1
	THEMES
	CONFIG
)

func (option Option) String() string {
	switch option {
	case ICONS:
		return ".icons"
	case THEMES:
		return ".themes"
	case CONFIG:
		return ".config"
	default:
		return ""
	}
}

type Command struct {
	Type     string
	Option   Option
	Packages []string
}

func MenuSelection() {
	var command Command
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions(ICONS.String(), THEMES.String(), CONFIG.String())...).
				Title("Linux Themes Remover").
				Value(&command.Type),
		),
	).WithAccessible(accessible).WithTheme(huh.ThemeDracula())
	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	packages := GetPackages(command.Type)
	total_packages := len(packages)
	str_builder := "Installed Themes: " + strconv.Itoa(total_packages)
	options := []huh.Option[string]{}
	for index := 0; index < len(packages); index++ {
		option := huh.Option[string]{}
		option.Key = packages[index]
		option.Value = packages[index]
		option.Selected(false)
		options = append(options, option)
	}
	form2 := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Options(options...).
				Title(str_builder).
				Value(&command.Packages),
		),
	).WithAccessible(accessible).WithTheme(huh.ThemeDracula())
	err = form2.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	loading_action_progress := func() {
		// time.Sleep(1 * time.Second)
		Remove(command.Packages)
		fmt.Println(command.Packages)
	}
	_ = spinner.New().
		Title("Removing Themes...").
		TitleStyle(lipgloss.NewStyle().Foreground(green)).
		Style(lipgloss.NewStyle().Foreground(red)).
		Accessible(accessible).
		Action(loading_action_progress).
		Run()

	fmt.Println(lipgloss.NewStyle().Foreground(green).SetString("Themes removed: " + strconv.Itoa(len(command.Packages))))
	for _, packag := range command.Packages {
		fmt.Println(packag)
	}

}

func GetPackages(category string) []string {
	var path string
	if category == ".icons" {
		path = ICON_PATH
	}
	if category == ".themes" {
		path = THEME_PATH
	}
	if category == ".config" {
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

func Remove(packages []string) {
	for _, packag := range packages {
		fmt.Print("Removing: ")
		fmt.Println(packag)
		if err := os.RemoveAll(packag); err != nil {
			fmt.Println(err.Error())
		}
	}
}
